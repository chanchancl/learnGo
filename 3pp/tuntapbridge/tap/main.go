package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/songgao/water"
)

const (
	deviceName = "tap0"
	deviceIP   = "10.0.0.100"
)

func main() {

	iface, err := water.New(water.Config{
		DeviceType: water.TAP,
		PlatformSpecificParams: water.PlatformSpecificParams{
			InterfaceName: deviceName,
		},
	})
	if err != nil {
		fmt.Println("Error to create tap")
		fmt.Println("This is only available on linux")
		return
	}

	// Setup tap
	cmd := exec.Command("ip", "link", "set", "dev", deviceName, "up")
	if err = cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}

	cmd = exec.Command("ip", "addr", "add", deviceIP+"/24", "dev", deviceName)
	if err = cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}

	log.Printf("Start listen on %v", deviceIP)

	buffer := make([]byte, 1500)
	for {
		n, err := iface.Read(buffer)
		if err != nil {
			log.Printf("iface read failed : %v", err)
			continue
		}

		ethernetPacket := gopacket.NewPacket(buffer[:n], layers.LayerTypeEthernet, gopacket.Default)

		// Tap Device could receive level 2 packet, which is ethernet packet
		if ethernetPacket == nil {
			fmt.Println("Invalid ethernet packet")
			continue
		}

		ipv6 := ethernetPacket.Layer(layers.LayerTypeIPv6)
		if ipv6 != nil {
			// fmt.Println(ipv6.LayerType())
			// fmt.Println("Skip ipv6 packet")
			continue
		}

		// printPacketInHex("Received", buffer[:n])
		// fmt.Println(ethernetPacket.String())

		if arpLayer := ethernetPacket.Layer(layers.LayerTypeARP); arpLayer != nil {
			printPacketInHex("ARP REQUEST", buffer[0:n])
			// fmt.Println(gopacket.LayerString(arpLayer))
			arpRequest := arpLayer.(*layers.ARP)
			fmt.Printf("%v ask who has %v\n", arpRequest.SourceProtAddress, arpRequest.DstProtAddress)

			handleARPRequest(iface, ethernetPacket, arpLayer)
		}

		if icmpLayer := ethernetPacket.Layer(layers.LayerTypeICMPv4); icmpLayer != nil {
			printPacketInHex("ICMP REQUEST", buffer[0:n])
			fmt.Println(ethernetPacket.String())

			handleICMPRequest(iface, ethernetPacket, icmpLayer)
		}

		// fmt.Println(packet.Dump())
		fmt.Println()
	}
}

func handleARPRequest(iface *water.Interface, packet gopacket.Packet, arpLayer gopacket.Layer) {
	arpRequest := arpLayer.(*layers.ARP)

	fakeMACAddr, err := net.ParseMAC("00:00:00:00:00:01")
	if err != nil {
		fmt.Printf("Parse MAC failed : %v", err)
		return
	}
	// sourceIpAddr := net.ParseIP("10.0.0.1")

	// 你要找的 DstIP 的 MAC 地址是 sourceMacAddr
	arpReply := &layers.ARP{
		AddrType:          arpRequest.AddrType,
		Protocol:          arpRequest.Protocol,
		HwAddressSize:     arpRequest.HwAddressSize,
		ProtAddressSize:   arpRequest.ProtAddressSize,
		DstHwAddress:      arpRequest.SourceHwAddress,
		DstProtAddress:    arpRequest.SourceProtAddress,
		SourceProtAddress: arpRequest.DstProtAddress,

		Operation:       layers.ARPReply,
		SourceHwAddress: fakeMACAddr,
	}

	ethernetLayer := &layers.Ethernet{
		SrcMAC:       fakeMACAddr,
		DstMAC:       arpRequest.SourceHwAddress,
		EthernetType: layers.EthernetTypeARP,
	}

	frame := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(frame, gopacket.SerializeOptions{}, ethernetLayer, arpReply)
	_, err = iface.Write(frame.Bytes())
	if err != nil {
		fmt.Printf("send arp reply failed : %v", err.Error())
	}
}

func handleICMPRequest(iface *water.Interface, packet gopacket.Packet, icmpLayer gopacket.Layer) {
	icmpPacket, _ := icmpLayer.(*layers.ICMPv4)

	if icmpPacket.TypeCode.Type() == layers.ICMPv4TypeEchoRequest {
		// printPacketInHex("icmp reqeust payload", icmpPacket.Payload)

		icmpReplyPacket := &layers.ICMPv4{
			TypeCode: layers.ICMPv4TypeEchoReply,
			Id:       icmpPacket.Id,
			Seq:      icmpPacket.Seq,
		}

		ipPacket := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
		ipPacket.DstIP, ipPacket.SrcIP = ipPacket.SrcIP, ipPacket.DstIP

		ethernetPacket := packet.Layer(layers.LayerTypeEthernet).(*layers.Ethernet)
		ethernetPacket.DstMAC, ethernetPacket.SrcMAC = ethernetPacket.SrcMAC, ethernetPacket.DstMAC

		frame := gopacket.NewSerializeBuffer()
		err := gopacket.SerializeLayers(frame, gopacket.SerializeOptions{
			FixLengths:       true,
			ComputeChecksums: true,
		}, ethernetPacket, ipPacket, icmpReplyPacket, gopacket.Payload(icmpPacket.Payload))
		if err != nil {
			log.Printf("serialize layers failed: %v", err)
			return
		}

		_, err = iface.Write(frame.Bytes())
		if err != nil {
			log.Printf("iface write failed: %v", err)
		}
	}
}

func printPacketInHex(name string, bytes []byte) {
	fmt.Printf("%s %s: ", time.Now().Format("2006-01-02 15:04:05"), name)
	for _, b := range bytes {
		fmt.Printf("%02x ", b)
	}

	fmt.Println()
}

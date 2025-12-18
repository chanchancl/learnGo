package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/songgao/water"
)

const (
	deviceName = "tun0"
	deviceIP   = "10.0.1.100"
)

func main() {

	iface, err := water.New(water.Config{
		DeviceType: water.TUN,
		PlatformSpecificParams: water.PlatformSpecificParams{
			InterfaceName: deviceName,
		},
	})
	if err != nil {
		fmt.Println("Error to create tun")
		return
	}

	// Setup tun
	cmd := exec.Command("ip", "link", "set", "dev", deviceName, "up")
	if err = cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}

	cmd = exec.Command("ip", "addr", "add", deviceIP+"/24", "dev", deviceName)
	if err = cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Listen on %v\n", deviceIP)

	buffer := make([]byte, 1500)
	for {
		n, err := iface.Read(buffer)
		if err != nil {
			log.Printf("iface read failed : %v", err)
			continue
		}

		// TUN device only receive ip level packet

		ethernetPacket := gopacket.NewPacket(buffer[:n], layers.EthernetTypeIPv4, gopacket.Default)
		if ethernetPacket.ErrorLayer() != nil {
			// This maybe ipv6 packet
			// ipv6 := gopacket.NewPacket(buffer[:n], layers.EthernetTypeIPv6, gopacket.Default)
			// fmt.Println(ipv6.String())
			continue
		}

		// printPacketInHex("Received", buffer[:n])
		fmt.Println(ethernetPacket.String())

		if icmpLayer := ethernetPacket.Layer(layers.LayerTypeICMPv4); icmpLayer != nil {
			printPacketInHex("ICMP REQUEST", buffer[0:n])
			fmt.Println(ethernetPacket.String())

			handleICMPRequest(iface, ethernetPacket, icmpLayer)
		}

		// fmt.Println(packet.Dump())
		fmt.Println()
	}
}

func handleICMPRequest(iface *water.Interface, packet gopacket.Packet, icmpLayer gopacket.Layer) {
	icmpPacket, _ := icmpLayer.(*layers.ICMPv4)

	if icmpPacket.TypeCode.Type() == layers.ICMPv4TypeEchoRequest {
		printPacketInHex("icmp reqeust payload", icmpPacket.Payload)

		icmpReplyPacket := &layers.ICMPv4{
			TypeCode: layers.ICMPv4TypeEchoReply,
			Id:       icmpPacket.Id,
			Seq:      icmpPacket.Seq,
		}

		ipPacket := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
		ipPacket.DstIP, ipPacket.SrcIP = ipPacket.SrcIP, ipPacket.DstIP

		frame := gopacket.NewSerializeBuffer()
		err := gopacket.SerializeLayers(frame, gopacket.SerializeOptions{
			FixLengths:       true,
			ComputeChecksums: true,
		}, ipPacket, icmpReplyPacket, gopacket.Payload(icmpPacket.Payload))
		if err != nil {
			log.Printf("serialize layers failed: %v", err)
			return
		}

		printPacketInHex("ICMP REPLY", frame.Bytes())

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

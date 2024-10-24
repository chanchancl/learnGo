#!/bin/bash

# ns1 veth1 <-> veth1_br br0
#                         |
# ns2 veth2 <-> veth2_br br0

# create 2 network namespace ns1 and ns2
# link ns1 and bridge br0 by veth1 and veth1_br
# link ns2 and bridge br0 by veth2 and veth2_br

# ns1 ping to ns2 through br0

# if we add addr for br0, such like 10.0.0.254
# then we can directly ping ns1 or ns2 through this bridge

create() {
    # create network namespace ns1 ns2
    ip netns add ns1
    ip netns add ns2

    # create bridge br0
    ip link add br0 type bridge
    ip link set br0 up

    # veth1 link veth2
    ip link add veth1 type veth peer name veth1_br
    ip link add veth2 type veth peer name veth2_br

    # set veth1 to netnamespace ns1, veth2 to ns2
    ip link set veth1 netns ns1
    ip link set veth2 netns ns2

    # ns1 veth1 <-> veth1_br br0
    #                         |
    # ns2 veth2 <-> veth2_br br0
    brctl addif br0 veth1_br
    brctl addif br0 veth2_br

    # add ip addr for veth1 and veth2
    ip netns exec ns1 ip addr add 10.0.0.100/24 dev veth1
    ip netns exec ns2 ip addr add 10.0.0.101/24 dev veth2

    # set up for veth1, veth2, veth1_br and veth2_br
    ip netns exec ns1 ip link set veth1 up
    ip netns exec ns2 ip link set veth2 up

    # set up for veth1_br veth2_br
    ip link set veth1_br up
    ip link set veth2_br up
}

show() {
    ip addr add 10.0.0.254/24 dev br0

    echo "Default link dev:"
    ip link list
    echo ""

    echo "ns1 link dev :"
    ip netns exec ns1 ip link list
    ip netns exec ns1 ifconfig
    echo ""

    echo "ns2 link dev :"
    ip netns exec ns2 ip link list
    ip netns exec ns2 ifconfig
    echo ""
}

pping() {
    ip netns exec ns1 ping -c2 10.0.0.101
}

clean() {
    ip netns exec ns1 ip link del veth1
    ip netns exec ns2 ip link del veth2
    ip link del br0
    ip netns del ns1
    ip netns del ns2
}

if [ "$1" = "" ]; then
    create
    show
    pping
    clean
elif [ "$1" = "create" ]; then
    create
    show
elif [ "$1" = "show" ]; then
    show
elif [ "$1" = "clean" ]; then
    clean
    show
elif [ "$1" = "ping" ]; then
    pping
fi

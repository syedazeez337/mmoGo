package main

import (
	"fmt"

	"github.com/syedazeez337/mmoGo/server/pkg/packets"
	"google.golang.org/protobuf/proto"
)

func main() {
	packet := &packets.Packet{
		SenderId: 69,
		Msg: packets.NewChat("Hello world!"),
	}

	fmt.Println(packet)

	data, err := proto.Marshal(packet)
	if err != nil {
		fmt.Println("Error mashalling packet:", err)
		return
	}

	fmt.Println(data)

	packet2 := &packets.Packet{}
	proto.Unmarshal(data, packet2)
	fmt.Println(packet2)
}

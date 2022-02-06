package main

import (
	“encoding/binary”
	“fmt”
)

func main () {
	data := encode("Goim")
	decode(data)
}

// PacketLen HeaderLen Version Operation Sequence Body
// 4bytes       2bytes         2bytes  4bytes      4bytes       PacketLen - HeaderLen

//PacketLen 包长度，在数据流传输过程中，先写入整个包的长度，方便整个包的数据读取。
//HeaderLen 头长度，在处理数据时，会先解析头部，可以知道具体业务操作。
//Version 协议版本号，主要用于上行和下行数据包按版本号进行解析。
//Operation 业务操作码，可以按操作码进行分发数据包到具体业务当中。
//Sequence 序列号，数据包的唯一标记，可以做具体业务处理，或者数据包去重。
//Body 实际业务数据，在业务层中会进行数据解码和编码。

func encode (body string) {
	headerLen := 16
	packetLen := len(body) + headerLen
	ret := make([]byte, packetLen)

	binary.BigEndian.PutUint32(ret[:4], uint32(packetLen))
	binary.BigEndian.PutUint16(ret[4:6], uint16(headerLen))

	version := 1
	binary.BigEndian.PutUint16(ret[6:8], uint16(version))
	operation := 2
	binary.BigEndian.PutUint32(ret[8:12], uint32(operation))
	sequence := 3
	binary.BigEndian.PutUint32(ret[12:16], uint32(sequence))

	byteBody := []byte(body)
	copy(ret[16:], byteBody)

	return ret
}

func decode (data []byte) {
	if len(data) <= 16 {
		fmt.Println("decode error: data len < 16.")
		return
	}

	packetLen := binary.BigEndian.Uint32(data[:4])
	fmt.Printf("packetLen:%v\n", packetLen)

	headerLen := binary.BigEndian.Uint16(data[4:6])
	fmt.Printf("headerLen:%v\n", headerLen)

	version := binary.BigEndian.Uint16(data[6:8])
	fmt.Printf("version:%v\n", version)

	operation := binary.BigEndian.Uint32(data[8:12])
	fmt.Printf("operation:%v\n", operation)

	sequence := binary.BigEndian.Uint32(data[12:16])
	fmt.Printf("sequence:%v\n", sequence)

	body := string(data[16:])
	fmt.Printf("body:%v\n", body)
}

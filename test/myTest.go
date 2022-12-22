package test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/gofrs/uuid"
	"log"
)

func Test1() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", "127.0.0.1", 8081)
	fmt.Println("hello zinx learn..")

	v1, err := uuid.NewV1()
	fmt.Println(v1, err)
}

func Test2() {
	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteString("hello")
	buffer.WriteString("world")
	fmt.Println("buffer : ", buffer.String())

	var buf bytes.Buffer
	buf.WriteString("你好")
	buf.WriteString("世界")
	fmt.Println("buf : ", buf.String())
}

/**
参数列表：
1）r  可以读出字节流的数据源
2）order  特殊字节序，包中提供大端字节序和小端字节序
3）data  需要解码成的数据
返回值：error  返回错误
功能说明：Read从r中读出字节数据并反序列化成结构数据。data必须是固定长的数据值或固定长数据的slice。从r中读出的数据可以使用特殊的 字节序来解码，并顺序写入value的字段。当填充结构体时，使用(_)名的字段讲被跳过。
*/
func Test3() {

	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		log.Fatalln("binary.Read failed:", err)
	}
	fmt.Println(pi)

	//buf := new(bytes.Buffer)
	//pi := int32(5)  // todo 不能使用int类型 源码不支持
	//
	//err := binary.Write(buf, binary.LittleEndian, pi)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(buf.Bytes())
}

func Test4() {

	//序列化
	var dataA uint64 = 6010
	var buffer bytes.Buffer
	err1 := binary.Write(&buffer, binary.BigEndian, &dataA)
	if err1 != nil {
		log.Panic(err1)
	}
	byteA := buffer.Bytes()
	fmt.Println("序列化之前：", byteA)

	//反序列化
	var dataB uint64
	var byteB []byte = byteA
	err2 := binary.Read(bytes.NewReader(byteB), binary.BigEndian, &dataB)
	if err2 != nil {
		log.Panic(err2)
	}
	fmt.Println("反序列化后：", dataB)
}

func Test5() {

	i1 := []byte{1, 2, 3}
	i2 := []byte{11, 21, 31}
	//i3 := []byte{12,22,32}
	i := append(i1, i2...)
	fmt.Println(i)
}

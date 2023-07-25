package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client)

	}

}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//for true {
	//	b, err := reader.ReadByte()
	//	if err != nil {
	//		break
	//	}
	//	_, err = conn.Write([]byte{b})
	//	if err != nil {
	//		break
	//	}
	//}
	err := auth(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed :%v", conn.RemoteAddr(), err)
		return
	}
	log.Println("auth success")
}

func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	ver, err := reader.ReadByte()
	if err != nil {
		return err
	}
	socks5Ver := byte(1)
	if ver != socks5Ver {
		return fmt.Errorf("bad ver")

	}

	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("bad methodsize")

	}
	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("bad method")
	}

	log.Println("ver", ver, "method", method)
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return err
	}
	return nil
}

func connect(reader *bufio.Reader, conn net.Conn) (err error) {
	buf := make([]byte, 4)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return err
	}
	ver, cmd, atyp := buf[0], buf[1], buf[3]

}

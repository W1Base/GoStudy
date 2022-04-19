package coo

import (
	"fmt"
	"net"
	"time"
)

func Icmp(host string) {
	conn, err := net.DialTimeout("ip4:icmp", host, 1*time.Second)
	if err != nil {
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	if err := conn.SetDeadline(time.Now().Add(1 * time.Second)); err != nil {
		return
	}
	msg := packet(host)
	if _, err := conn.Write(msg); err != nil {
		return
	}
	var receive = make([]byte, 60)
	if _, err := conn.Read(receive); err != nil {
		return
	}
	//vHosts = append(vHosts, host)
	fmt.Println("[*]", host, " 'is Alive'")
}

func packet(host string) []byte {
	var msg = make([]byte, 40)
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4], msg[5] = host[0], host[1]
	msg[6], msg[7] = byte(1>>8), byte(1&255)
	msg[2] = byte(checksum(msg[0:40]) >> 8)
	msg[3] = byte(checksum(msg[0:40]) & 255)
	return msg
}

func checksum(msg []byte) uint16 {
	var sum = 0
	var length = len(msg)
	for i := 0; i < length-1; i += 2 {
		sum += int(msg[i])*256 + int(msg[i+1])
	}
	if length%2 == 1 {
		sum += int(msg[length-1]) * 256
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum = sum + (sum >> 16)
	return uint16(^sum)
}

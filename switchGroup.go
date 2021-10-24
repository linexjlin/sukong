package sk

import (
	"fmt"
)

func SwitchGroupTo(i int) ([]byte, error) {
	fmt.Printf("%x", i)
	buf := []byte("\x55\xaa\x07\x02\x05\x00\x01\x00\x00\x0d\x0a")
	buf[6] = byte(i)
	crcv := crc(buf[:7])
	h, l := byte(crcv>>8), byte(crcv&0xff)
	buf[7], buf[8] = l, h
	return buf, nil
}

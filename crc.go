package sk

func crc(buf []byte) uint16 {
	var tem, crc uint16
	crc = 0xFFFF
	for i := 0; i < len(buf); i++ {
		crc = uint16(buf[i]) ^ crc
		for j := 0; j < 8; j++ {
			tem = crc & 0x0001
			crc = crc >> 1
			if tem != 0 {
				crc = crc ^ 0xA001
			}
		}
	}
	return crc
}

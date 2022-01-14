package sk

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/linexjlin/simple-log"
)

type SK struct {
	Header        [2]byte //fixed 55 aa
	Length        uint8
	Cmd           [2]byte //81 00
	SiteNo        uint    //01
	Unit          string  //0:kgf.cm 1:N.m
	TorqueValue   uint16  //01 00
	SpinSpeed     uint16  //01 00
	LockAngle     uint16  //01 00
	TwistAngle    uint16  //01 00
	Duration      uint16  //01 00
	Forward       bool    //0: forward 1:backword
	LeftCnt       uint16  //01 00
	GroupFinish   string  //0:unfinsh 1:finish 2:all done
	TwistStatus   string  //0 ing,1 ok, NG
	ErrorReport   string  //1: 滑丝 2:浮锁 3:扭矩不良 4:拧紧角度不良 05 中途提前释放启动信号
	Temp          int8
	MaxTorque     uint16 //00 01
	MinTorque     uint16
	TwistAngleMax uint16
	TwistAngleMin uint16
	LockAngleMax  uint16
	LockAngleMin  uint16
	Mode          string //00
	Batch         byte   //00
	Group         byte   //00
	CRC1          byte
	CRC2          byte
	Tailer        [2]byte
}

func MakePCCmd() []byte {
	hexStr := "55aa070100010072630d0a"
	data, _ := hex.DecodeString(hexStr)
	return data
}

func Parse(data []byte) (SK, error) {
	var sk SK
	if len(data) != 43 {
		return sk, errors.New("frame broken!")
	}

	r := bytes.NewReader(data)
	var b byte
	var b2 [2]byte
	r.Read(sk.Header[:])
	b, _ = r.ReadByte()
	sk.Length = uint8(b)

	r.Read(sk.Cmd[:])

	b, _ = r.ReadByte()
	sk.SiteNo = uint(b)

	b, _ = r.ReadByte()
	switch b {
	case 0x0:
		sk.Unit = "kgf.cm"
	case 0x1:
		sk.Unit = "N.m"
	default:
		log.Errorf("Unknow Unit %x\n", b)
	}
	fmt.Printf("%x", b)

	r.Read(b2[:])
	sk.TorqueValue = binary.BigEndian.Uint16(b2[:])

	r.Read(b2[:])
	sk.SpinSpeed = binary.BigEndian.Uint16(b2[:])

	r.Read(b2[:])
	sk.LockAngle = binary.BigEndian.Uint16(b2[:])

	r.Read(b2[:])
	sk.TwistAngle = binary.BigEndian.Uint16(b2[:])

	r.Read(b2[:])
	sk.Duration = binary.BigEndian.Uint16(b2[:])

	b, _ = r.ReadByte()
	switch b {
	case 0x0:
		sk.Forward = true
	case 0x1:
		sk.Forward = false
	default:
		log.Errorf("Unknow forward %x\n", b)
	}

	r.Read(b2[:])
	sk.LeftCnt = binary.LittleEndian.Uint16(b2[:])

	b, _ = r.ReadByte()
	switch b {
	case 0x0:
		sk.GroupFinish = "整个工件还没打完"
	case 0x1:
		sk.GroupFinish = "当前组别打完了"
	case 0x2:
		sk.GroupFinish = "全部打完了"
	default:
		log.Error("Unknow GroupFinish", b)
	}

	b, _ = r.ReadByte()
	switch b {
	case 0x0:
		sk.TwistStatus = "正在拧紧"
	case 0x1:
		sk.TwistStatus = "拧紧OK"
	case 0x2:
		sk.TwistStatus = "拧紧NG"
	default:
		log.Error("Unknow sk.TwistStatus", b)
	}

	b, _ = r.ReadByte()
	switch b {
	case 0x0:
		sk.ErrorReport = ""
	case 0x1:
		sk.ErrorReport = "滑丝"
	case 0x2:
		sk.ErrorReport = "浮锁"
	case 0x3:
		sk.ErrorReport = "扭矩不良"
	case 0x4:
		sk.ErrorReport = "拧紧角度不良"
	case 0x5:
		sk.ErrorReport = "中途提前释放启动信号"
	default:
		log.Error("Unknow sk.ErrorReport")
	}

	b, _ = r.ReadByte()
	sk.Temp = int8(b)

	r.Read(b2[:])
	sk.MaxTorque = binary.LittleEndian.Uint16(b2[:])

	r.Read(b2[:])
	sk.MinTorque = binary.LittleEndian.Uint16(b2[:])

	r.Read(b2[:])
	sk.TwistAngleMax = binary.LittleEndian.Uint16(b2[:])

	r.Read(b2[:])
	sk.TwistAngleMin = binary.LittleEndian.Uint16(b2[:])

	r.Read(b2[:])
	sk.LockAngleMax = binary.LittleEndian.Uint16(b2[:])

	r.Read(b2[:])
	sk.LockAngleMin = binary.LittleEndian.Uint16(b2[:])

	b, _ = r.ReadByte()
	switch b {
	case 0x0:
		sk.Mode = "单组模式"
	case 0x1:
		sk.Mode = "批次模式"
	default:
		log.Errorf("Unknow sk.Mode %x\n", b)
	}
	b, _ = r.ReadByte()
	sk.Batch = b
	b, _ = r.ReadByte()
	sk.Group = b

	b, _ = r.ReadByte()
	sk.CRC1 = b

	b, _ = r.ReadByte()
	sk.CRC2 = b

	r.Read(sk.Tailer[:])

	log.Debug(sk)

	return sk, nil
}

package sk

import (
	"encoding/hex"
	"testing"
)

func TestSK(t *testing.T) {
	data, _ := hex.DecodeString("55aa27810002011e001d00560000002d0100050000000027b0046400e7030000e010380400000273870d0a")
	data2, _ := hex.DecodeString("55aa278100020100001d00580000009d0200050000020527b0046400e7030000e0103804000002b1930d0a")
	if sk1, err := Parse(data); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(sk1)
	}
	if sk2, err := Parse(data2); err != nil {
		t.Fail()
	} else {
		t.Log(sk2)
	}
}

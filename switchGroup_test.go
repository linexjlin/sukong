package sk

import "testing"
import "log"

func TestSwithTo(t *testing.T) {
	if cmd, err := SwitchGroupTo(20); err != nil {
		t.Fail()
	} else {
		log.Printf("%x", cmd)
	}
}

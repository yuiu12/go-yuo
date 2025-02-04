package main

import (
	"testing"
	"time"
)

func TestWhatsTheClock(t *testing.T) {
	result := whatstheclock()
	if result != time.Now().Format(time.ANSIC) {
		t.Errorf("Function failed!")
	} else {
		t.Log("Function working properly!")
	}
}

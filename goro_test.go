package gorocheck

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	err := CheckForLeaks()
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()

	err = CheckForLeaks()
	if err == nil {
		t.Fatal("Expected check to fail.")
	}
}

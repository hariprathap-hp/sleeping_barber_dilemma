package models

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestNewBarberShop(t *testing.T) {
	shop := NewBarberShop()

	if shop.BarbersEmployed != NumberofBarbers {
		t.Errorf("Expected %d barbers, got %d", NumberofBarbers, shop.BarbersEmployed)
	}

	if cap(shop.WaitingRoom) != WaitingChairs {
		t.Errorf("Expected waiting room capacity %d, got %d", WaitingChairs, cap(shop.WaitingRoom))
	}

	if shop.IsShopOpen != true {
		t.Error("Expected shop to be open, got closed")
	}
}

func TestCloseChannels(t *testing.T) {
	shop := NewBarberShop()
	shop.CloseChannels()

	// Try sending to a closed channel (should panic)
	defer func() { recover() }()
	shop.WaitingRoom <- 1
	t.Error("Failed to close channels")
}

func captureOutput(source func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	source()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	return out
}

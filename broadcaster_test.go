package broadcaster

import (
	"log"
	"testing"
	"time"
)

func TestBroadcast(t *testing.T) {
	b := NewBroadcaster()

	done := make(chan bool, 2)
	b.Wait(func() {
		log.Println("function 1 finished")
		done <- true
	})
	b.Wait(func() {
		log.Println("function 2 finished")
		done <- true
	})

	if len(done) != 0 {
		t.Fatalf("Failed to wait for broadcast")
	}

	b.Broadcast()
	<-done
	<-done
}

func TestLateBroadcast(t *testing.T) {
	b := NewBroadcaster()

	done := make(chan bool, 2)
	b.Wait(func() {
		log.Println("function 1 finished")
		done <- true
	})
	b.Wait(func() {
		log.Println("function 2 finished")
		done <- true
	})

	go func() {
		time.Sleep(100 * time.Millisecond)
		b.Broadcast()
	}()

	<-done
	<-done
}

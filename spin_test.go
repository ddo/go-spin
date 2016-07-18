package spin

import (
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	spinner := New("")

	if spinner.index != -1 {
		t.Error()
	}

	if spinner.sprite != DEFAULT_SPRITE {
		t.Error()
	}
}

func TestSpin(t *testing.T) {
	spinner := New("")

	// spin every 100ms
	ticker := time.NewTicker(time.Millisecond * 100)

	go func() {
		for range ticker.C {
			fmt.Printf("\r %s", spinner.Spin())
		}
	}()

	// stop after 5s
	<-time.After(time.Second * 5)

	// this stop doesnt close the ticker.C channel
	ticker.Stop()

	fmt.Println("done")
}

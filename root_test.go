package log

import (
	"testing"
	"time"
)

func TestInfod1(t *testing.T) {
	root.SetHandler(StdoutHandler)

	for i := 0; i < 30; i++ {
		Infod1("hello", "k", "v")
		time.Sleep(time.Millisecond * 200)
	}
	for i := 0; i < 30; i++ {
		Infod(time.Second*3, "hello2", "k", "v")
		time.Sleep(time.Millisecond * 500)
	}
	time.Sleep(time.Second * 5)
}

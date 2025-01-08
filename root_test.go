package log

import (
	"testing"
	"time"
)

func TestInfod1(t *testing.T) {
	root.SetHandler(DefaultFileHandler())

	for i := 0; i < 30; i++ {
		Infod1("hello", "k", "v")
		time.Sleep(time.Millisecond * 200)
	}
	time.Sleep(time.Second)
}
func TestInfo2(t *testing.T) {
	root.SetHandler(DefaultFileHandler())

	time.Sleep(time.Second)
}

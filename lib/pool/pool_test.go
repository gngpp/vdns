package pool

import (
	"fmt"
	"testing"
	"time"
	"vdns/lib/vlog"
)

func Test(t *testing.T) {
	vlog.SetLevel(vlog.Level.DEBUG)
	pool := NewPool(2, 9, true)
	var d Payload = "hanbi"
	fn := func(payload Payload) {
		fmt.Println("done")
	}
	pool.AddTask(fn, d)
	time.Sleep(time.Hour)
}

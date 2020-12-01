package schedule_job

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

var sum int32 = 0
var N int32 = 300
var tt *Timer

func now() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	atomic.AddInt32(&sum, 1)
	val := atomic.LoadInt32(&sum)
	if val == 2*N {
		tt.Stop()
	}
}

func TestTimer(t *testing.T) {
	// 全局的最低任务执行间隔时间
	timer := New(time.Millisecond)
	tt = timer
	fmt.Println(timer)
	var i int32
	for i = 0; i < N; i++ {
		timer.NewTimer(time.Millisecond*time.Duration(i), now)
		timer.NewTimer(time.Millisecond*time.Duration(i), now)
	}
	timer.Start()
	if sum != 2*N {
		t.Error("failed")
	}
}

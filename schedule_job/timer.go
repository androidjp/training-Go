package schedule_job

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

const (
	TimeNearShift  = 8
	TimeNear       = 1 << TimeNearShift
	TimeLevelShift = 6
	TimeLevel      = 1 << TimeLevelShift
	TimeNearMask   = TimeNear - 1
	TimeLevelMask  = TimeLevel - 1
)

// 标准Linux时间轮实现方式
type Timer struct {
	near [TimeNear]*list.List
	t    [4][TimeLevel]*list.List
	sync.Mutex
	time uint32
	tick time.Duration
	quit chan struct{}
}

type Node struct {
	expire uint32
	f      func()
}

func (n *Node) String() string {
	return fmt.Sprintf("Node:expire,%d", n.expire)
}

// New 新建一个定时调度器，入参是两次调度之间的间隔时间
func New(d time.Duration) *Timer {
	t := new(Timer)
	t.time = 0
	t.tick = d
	t.quit = make(chan struct{})

	var i, j int
	for i = 0; i < TimeNear; i++ {
		t.near[i] = list.New()
	}

	for i = 0; i < 4; i++ {
		for j = 0; j < TimeLevel; j++ {
			t.t[i][j] = list.New()
		}
	}

	return t
}

func (t *Timer) addNode(n *Node) {
	expire := n.expire
	current := t.time
	if (expire | TimeNearMask) == (current | TimeNearMask) {
		t.near[expire&TimeNearMask].PushBack(n)
	} else {
		var i uint32
		var mask uint32 = TimeNear << TimeLevelShift
		for i = 0; i < 3; i++ {
			if (expire | (mask - 1)) == (current | (mask - 1)) {
				break
			}
			mask <<= TimeLevelShift
		}

		t.t[i][(expire>>(TimeNearShift+i*TimeLevelShift))&TimeLevelMask].PushBack(n)
	}

}

func (t *Timer) NewTimer(d time.Duration, f func()) *Node {
	n := new(Node)
	n.f = f
	t.Lock()
	n.expire = uint32(d/t.tick) + t.time
	t.addNode(n)
	t.Unlock()
	return n
}

func (t *Timer) String() string {
	return fmt.Sprintf("Timer:time:%d, tick:%s", t.time, t.tick)
}

func dispatchList(front *list.Element) {
	for e := front; e != nil; e = e.Next() {
		node := e.Value.(*Node)
		go node.f()
	}
}

func (t *Timer) moveList(level, idx int) {
	vec := t.t[level][idx]
	front := vec.Front()
	vec.Init()
	for e := front; e != nil; e = e.Next() {
		node := e.Value.(*Node)
		t.addNode(node)
	}
}

func (t *Timer) shift() {
	t.Lock()
	var mask uint32 = TimeNear
	t.time++
	ct := t.time
	if ct == 0 {
		t.moveList(3, 0)
	} else {
		time := ct >> TimeNearShift
		var i int = 0
		for (ct & (mask - 1)) == 0 {
			idx := int(time & TimeLevelMask)
			if idx != 0 {
				t.moveList(i, idx)
				break
			}
			mask <<= TimeLevelShift
			time >>= TimeLevelShift
			i++
		}
	}
	t.Unlock()
}

func (t *Timer) execute() {
	t.Lock()
	idx := t.time & TimeNearMask
	vec := t.near[idx]
	if vec.Len() > 0 {
		front := vec.Front()
		vec.Init()
		t.Unlock()
		// dispatch_list don't need lock
		dispatchList(front)
		return
	}

	t.Unlock()
}

func (t *Timer) update() {
	// try to dispatch timeout 0 (rare condition)
	t.execute()

	// shift time first, and then dispatch timer message
	t.shift()

	t.execute()

}

func (t *Timer) Start() {
	tick := time.NewTicker(t.tick)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			t.update()
		case <-t.quit:
			return
		}
	}
}

func (t *Timer) Stop() {
	close(t.quit)
}

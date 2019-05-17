package timer

import (
	"log"
	"testing"
	"time"
)

func TestNew(t *testing.T) {

	var timer *Timer
	var exit = make(chan bool, 1)
	var startedAt time.Time

	timer = New(Options{
		Duration:       1000 * time.Millisecond,
		TickerInternal: 100 * time.Millisecond,
		OnRun: func(started bool) {
			if started {
				startedAt = time.Now()
				log.Printf("[STARTED] %v", startedAt)
			} else {
				log.Printf("[RESUMED] %v", time.Now())
			}
		},
		OnPaused: func() {
			log.Printf("[PASUED] (прошло=%v, осталось=%v)", timer.Passed(), timer.Remaining())
		},
		OnDone: func(stopped bool) {
			log.Printf("[DONE] passed=%v", timer.Passed())
			log.Println(time.Now())
			exit <- true
		},
		OnTick: func() {
			log.Printf("[TICKED] %v -> %v", timer.Passed(), timer.Remaining())
		},
	})
	go timer.Run()

	//go func() {
	//	time.Sleep(500 * time.Millisecond)
	//	timer.Pause()
	//}()
	//
	//go func() {
	//	time.Sleep(3000 * time.Millisecond)
	//	timer.Run()
	//}()

	<-exit
}

func Test1(t *testing.T) {

}

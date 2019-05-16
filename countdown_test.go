package timer

import (
	"log"
	"testing"
	"time"
)

func TestNew(t *testing.T) {

	var c *Countdown

	go func() {
		c = New(Options{
			Duration:       1 * time.Second,
			TickerInternal: 100 * time.Millisecond,
			OnRun: func(started bool) {
				if started {
					log.Printf("начат: %v", time.Now())
				} else {
					log.Printf("продолжение: %v", time.Now())
				}
			},
			OnPaused: func(passed, remained Duration) {
				log.Printf("пауза (прошло=%v, осталось=%v)", passed.Millisecondable(), remained.Millisecondable())
			},
			//OnResumed: func(passed, remained countdown.Duration) {
			//	log.Printf("продолжение (прошло=%v, осталось=%v)", passed.Millisecondable(), remained.Millisecondable())
			//},
			OnDone: func(stopped bool) {
				log.Printf("завершен: %v (остановлен=%v)", time.Now(), stopped)
				log.Println(Duration(c.Passed()).Millisecondable())
			},
			OnTick: func(passed, remained Duration) {
				log.Printf("прошло времени: %v, осталось времени: %v", passed.Millisecondable(), remained.Millisecondable())
			},
		})
		go c.Run()
	}()

	//go func() {
	//	time.Sleep(500 * time.Millisecond)
	//	c.Pause()
	//}()
	//
	//go func() {
	//	time.Sleep(1000 * time.Millisecond)
	//	c.Run()
	//}()

	time.Sleep(2 * time.Second)
}

package timer

import (
	"time"
)

// Options represents countdown configuration.
type Options struct {
	Duration       time.Duration
	Passed         time.Duration
	TickerInternal time.Duration
	OnPaused       func(passed, remained Duration)
	OnDone         func(stopped bool)
	OnTick         func(passed, remained Duration)
	OnRun          func(started bool)
}

// Countdown represents back timer.
type Countdown struct {
	options  Options
	ticker   *time.Ticker
	started  bool
	active   bool
	passed   time.Duration
	lastTick time.Time
}

func (c Countdown) Passed() time.Duration {
	return c.passed
}

func (c Countdown) Remained() time.Duration {
	return c.options.Duration - c.passed
}

func (c Countdown) tick() {
	c.options.OnTick(Duration(c.Passed()), Duration(c.Remained()))
}

// Run starts just created countdown and resumes paused.
func (c *Countdown) Run() {
	if !c.started {
		c.start()
	} else {
		c.resume()
	}
}

// start runs just created countdown.
func (c *Countdown) start() {
	c.started = true
	c.active = true
	c.ticker = time.NewTicker(c.options.TickerInternal)
	c.lastTick = time.Now()
	c.options.OnRun(true)
	for tickAt := range c.ticker.C {
		c.passed += tickAt.Sub(c.lastTick)
		c.tick()
		c.lastTick = time.Now()
		if c.Remained() <= 0 {
			stop(c)
		}
	}
}

// resume continues recently paused countdown.
func (c *Countdown) resume() {
	c.active = true
	c.ticker = time.NewTicker(c.options.TickerInternal)
	c.lastTick = time.Now()
	c.options.OnRun(false)
	for tickAt := range c.ticker.C {
		c.passed += tickAt.Sub(c.lastTick)
		c.tick()
		c.lastTick = time.Now()
		if c.Remained() < 0 {
			stop(c)
		}
	}
}

func stop(c *Countdown) {
	c.ticker.Stop()
	c.active = false
	c.options.OnDone(false)
}

// Pause temporarily pauses active countdown.
func (c *Countdown) Pause() {
	c.ticker.Stop()
	c.active = false
	c.passed += time.Now().Sub(c.lastTick)
	c.lastTick = time.Now()
	c.options.OnPaused(Duration(c.passed), Duration(c.Remained()))
}

// Stop finishes the countdown.
func (c *Countdown) Stop() {
	c.ticker.Stop()
	c.active = false
	c.options.OnDone(true)
}

// New creates countdown.
func New(options Options) *Countdown {
	return &Countdown{
		options: options,
	}
}

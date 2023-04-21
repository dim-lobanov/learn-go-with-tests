package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStartingValue = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// We want Countdown to write data somewhere, so it can implement io.Writer

func Countdown(buf io.Writer, sleeper Sleeper) {
	for i := countdownStartingValue; i > 0; i-- {
		fmt.Fprintln(buf, i)
		sleeper.Sleep()
	}
	fmt.Fprint(buf, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
package racer

import (
	"fmt"
	"net/http"
	"time"
)

// <-ch - it's a blocking call
// and select allows us to wait on multiple channels
// The first one to send a value "wins" and the code underneath the case is executed.

const tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a,b,tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout): // time.After return chan and very handy in select expression
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// don't care what type is sent to the channel, we just want to signal we are done
// and chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool
func ping(url string) chan struct{} {
	ch := make(chan struct{}) // For channels the zero value is nil and if you try and send to it with <- it will block forever
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

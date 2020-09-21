package try

import (
	"errors"
	"time"
)

// MaxRetries is the maximum number of retries before bailing.
var MaxRetries = 10

var errMaxRetriesReached = errors.New("exceeded retry limit")

// Func represents functions that can be retried.
type Func func(attempt int) (retry bool, err error)
type DelayFunc func(sleepTime time.Duration) (time.Duration, error)

// Do keeps trying the function until the second argument
// returns false, or no error is returned.
func Do(fn Func) error {
	var err error
	var cont bool
	attempt := 1
	for {
		cont, err = fn(attempt)
		if !cont || err != nil {
			break
		}
		attempt++
		if attempt > MaxRetries {
			return errMaxRetriesReached
		}
	}
	return err
}

// IsMaxRetries checks whether the error is due to hitting the
// maximum number of retries or not.
func IsMaxRetries(err error) bool {
	return err == errMaxRetriesReached
}

// DelayDo delay some time.
// If delayFn is nil that will use defaultDelay.
func DelayDo(fn Func, delayFn ...DelayFunc) error {
	var err error
	var cont bool
	var sleep time.Duration
	attempt := 1
	for {
		cont, err = fn(attempt)
		if !cont || err != nil {
			break
		}
		attempt++
		if attempt > MaxRetries {
			return errMaxRetriesReached
		}

		var delayFunc DelayFunc
		if delayFn == nil || len(delayFn) == 0 {
			delayFunc = defaultDelay
		} else {
			delayFunc = delayFn[0]
		}
		sleep, err = delayFunc(sleep)
		if err != nil {
			return err
		}
		time.Sleep(sleep)
	}
	return err
}

func defaultDelay(delayTime time.Duration) (time.Duration, error) {
	if delayTime == 0 {
		return time.Second,nil
	}

	maxDelayTime := 2 * time.Minute
	delayTime = 2 * delayTime
	if delayTime > maxDelayTime {
		return maxDelayTime,nil
	}

	return delayTime,nil
}

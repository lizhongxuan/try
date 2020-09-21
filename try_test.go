package try

import (
	"errors"
	"fmt"
	"testing"
)

func TestTryExample(t *testing.T) {
	SomeFunction := func() (string, error) {
		panic("something went badly wrong")
	}
	var value string
	err := Do(func(attempt int) (retry bool, err error) {
		retry = attempt < 5 // try 5 times
		defer func() {
			if r := recover(); r != nil {
				err = errors.New(fmt.Sprintf("panic: %v", r))
			}
		}()
		value, err = SomeFunction()
		return
	})
	if err != nil {
		//log.Fatalln("error:", err)
	}
}


func TestDelayTryExample(t *testing.T) {
	SomeFunction := func() (string, error) {
		panic("something went badly wrong")
	}
	var value string
	err := DelayDo(func(attempt int) (retry bool, err error) {
		retry = attempt < 5 // try 5 times
		defer func() {
			if r := recover(); r != nil {
				err = errors.New(fmt.Sprintf("panic: %v", r))
			}
		}()
		value, err = SomeFunction()
		return
	})
	if err != nil {
		//log.Fatalln("error:", err)
	}
}
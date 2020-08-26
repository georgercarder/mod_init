package mod_init

import (
	"time"
)

type Callable func() interface{}

func NewModInit(fn Callable, timeout time.Duration,
	timeoutErrMsg error) (m *modInitialzer) {
	return &modInitialzer{mod_CH: fnCallToChannel(fn),
		timeout:       timeout,
		timeoutErrMsg: timeoutErrMsg}
}

func (m *modInitialzer) Get() (ret interface{}, err error) {
	if m.mod == nil {
		to := make(chan bool, 1)
		go func() {
			time.Sleep(m.timeout)
			to <- true
		}()
		select {
		case m.mod = <-m.mod_CH:
			go func() {
				// pass it to the others who are waiting
				// so they won't hang forever
				m.mod_CH <- m.mod
			}()
			break
		case _ = <-to:
			err = m.timeoutErrMsg
			return
		}
	}
	ret = m.mod
	return
}

type modInitialzer struct {
	mod           interface{}
	mod_CH        chan interface{}
	timeout       time.Duration
	timeoutErrMsg error
}

func fnCallToChannel(fn Callable) chan interface{} {
	ch := make(chan interface{})
	go func() {
		ch <- fn()
	}()
	return ch
}

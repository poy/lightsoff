package lightsoff

import "sync/atomic"

type LightsOff struct {
	count    int32
	callback func()
}

func New(count int, callback func()) *LightsOff {
	argValidation(count, callback)
	return &LightsOff{
		count:    int32(count),
		callback: callback,
	}
}

func (l *LightsOff) TurnOff() {
	if atomic.AddInt32(&l.count, -1) == 0 {
		l.callback()
	}
}

func argValidation(count int, callback func()) {
	if count <= 0 {
		panic("count is less than 0")
	}

	if callback == nil {
		panic("callback is nil")
	}
}

package waitfor

import "time"

type Waiter func() bool

func Func(waiter Waiter) {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			if waiter() {
				return
			}
		}
	}
}

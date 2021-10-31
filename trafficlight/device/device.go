package device

import (
	"fmt"
)

const (
	GREEN = iota
	YELLOW
	RED
)

type Light struct {
	status chan int
	handle func(int)
}

func (light *Light) runDevice(interrupt chan struct{}, status int) {
	for {
		select {
		case <-interrupt:
			light.handle(2)
			fmt.Println("Light is RED, stop moving")
			return
		default:
			light.handle(status)
			fmt.Println("Light is: ", status, ", keep moving")
			//time.Sleep(1 * time.Second)
		}
	}
}

func (light *Light) initDevice() {
	interrupt := make(chan struct{})
	for {
		select {
		case status := <-light.status:
			if status == GREEN || status == YELLOW {
				go light.runDevice(interrupt, status)
			}
			if status == RED {
				interrupt <- struct{}{}
			}
		}
	}
}

func (light *Light) KeepMoving() {
	light.status <- GREEN
}

func (light *Light) SlowMoving() {
	light.status <- YELLOW
}

func (light *Light) StopMoving() {
	light.status <- RED
}

func NewLight(h func(x int)) *Light {
	light := &Light{
		status: make(chan int),
		handle: h,
	}
	go light.initDevice()
	return light
}

func CloseLight(ligh *Light) {
	close(ligh.status)
}

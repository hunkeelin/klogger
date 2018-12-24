package klogger

import (
	"fmt"
)

func (k klogger) Loglevel(l int) {
	k.loglevel = l
}

func (k klogger) Log(loglvl int, msg string) {
	if loglvl <= k.loglevel {
		fmt.Println(msg)
	}
}

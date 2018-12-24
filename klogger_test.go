package klogger

import (
	"testing"
)

var klog = klogger{}

func TestLog(t *testing.T) {
	klog.loglevel = 2
	klog.Log(3, "fuck")
}

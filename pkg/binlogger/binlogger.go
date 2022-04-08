package binlogger

import (
	"time"
)

func Listen() {
	go binlogListener()

	for {
		time.Sleep(2 * time.Second)
	}
}

package gprint

import (
	"fmt"
	"time"
)

func Lblsl(text string, t time.Duration, nl bool) {

	for _, v := range text {
		fmt.Printf("%v", string(v))
		time.Sleep(t * time.Millisecond)
	}

	if nl == true {
		fmt.Printf("\n")
	}
}

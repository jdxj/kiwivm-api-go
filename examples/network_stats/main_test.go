package main

import (
	"fmt"
	"testing"
	"time"
)

func TestToday0oClock(t *testing.T) {
	ts := today0oClock()
	fmt.Printf("%s\n", time.Unix(ts, 0).Format(time.RFC3339))
}

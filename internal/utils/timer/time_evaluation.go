package timer

import (
	"fmt"
	"time"
)

func Evaluate(name string) func() {
	start := time.Now()

	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

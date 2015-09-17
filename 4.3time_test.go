package sleep

import (
	"fmt"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	current := time.Now()
	sleep(5)
	elapsedTime := time.Since(current)
	t.Log("elapsed time is", elapsedTime)
}

func sleep(n int) {
	select {
	case <-time.After(time.Duration(n) * time.Second):
		fmt.Println("slept for", n)
	}
}

func main() {
	var n int
	fmt.Println("Enter time to sleep")
	fmt.Scanf("%d", &n)
	sleep(5)

	var input string
	fmt.Scanln(&input)

}

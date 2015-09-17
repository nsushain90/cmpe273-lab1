package main




import (

	"fmt"

	"time"

)




func sleep(n int) {

	select {

	case <-time.After(time.Duration(n) * time.Second):

		fmt.Println("slept for" ,n)

	}

}




func main() {

	var n int

	fmt.Println("Enter time to sleep")

	fmt.Scanf("%d", &n)

	sleep(n)


        var input string

	fmt.Scanln(&input)




}
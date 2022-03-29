//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func myFunc() {
//	fmt.Println("Inside my goroutine")
//}
//
//func main() {
//	fmt.Println("Hello World")
//	go myFunc()
//	fmt.Println("Finished Execution")
//}

// run above, you won't see one of the print statements, it is like await in node.js

//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func myFunc(waitGroup *sync.WaitGroup) {
//	fmt.Println("Inside my goroutine")
//	waitGroup.Done()
//}
//
//func main() {
//	fmt.Println("Hello World")
//
//	var waitGroup sync.WaitGroup
//	waitGroup.Add(1)
//	go myFunc(&waitGroup)
//	waitGroup.Wait()
//
//	fmt.Println("Finished Execution")
//}

// now it works as charm

//Apply above goroutine with anonymous function

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello World")

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		fmt.Println("Inside my goroutine")
		waitGroup.Done()
	}()
	waitGroup.Wait()

	fmt.Println("Finished Execution")
}

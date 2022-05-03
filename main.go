package main

import (
	"fmt"

	"besimgurbuz.com/concurrency"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipAddr IPAddr) String() string {
	address := ""
	for _, num := range ipAddr {
		if address == "" {
			address = fmt.Sprint(num)
			continue
		}
		address = fmt.Sprint(address, ".", num)
	}

	return address
}

func main() {
	// hosts := map[string]IPAddr{
	// 	"loopback":  {127, 0, 0, 1},
	// 	"googleDNS": {8, 8, 8, 8},
	// }
	// for name, ip := range hosts {
	// 	fmt.Printf("%v: %v\n", name, ip)
	// }
	// myLinkedList := &list.List[string]{
	// 	Next: &list.List[string]{
	// 		Next: &list.List[string]{
	// 			Next: &list.List[string]{
	// 				Next: &list.List[string]{
	// 					Next: nil,
	// 					Val:  "fifth item",
	// 				},
	// 				Val: "fourth item",
	// 			},
	// 			Val: "thirth item",
	// 		},
	// 		Val: "second item",
	// 	},
	// 	Val: "first item",
	// }

	// myLinkedList.Print()

	// fmt.Println("REMOVING BY INDEX")
	// _, err := myLinkedList.RemoveByIndex(0)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// myLinkedList.Print()

	// fmt.Println("Removing by pointer")
	// _, err = myLinkedList.Remove(myLinkedList)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// myLinkedList.Print()

	// fmt.Println("Adding!")

	// myLinkedList.Add(list.List[string]{
	// 	Next: nil,
	// 	Val:  "ayo ayo ayo this is new item",
	// }, 0)

	// myLinkedList.Print()

	// // Convert to list
	// for i, item := range myLinkedList.ConvertToSlice() {
	// 	fmt.Printf("Index: %d Val: %v\n", i, item)
	// }

	// go concurrency.Say("another thread: parallel job 1")
	// concurrency.Say("main thread: parallel job 2")

	// concurrency.SumMain()

	// concurrency.BufferedMain()

	// concurrency.FibonacciMain()

	// concurrency.FibonacciMainWithSelect()

	// concurrency.DefaultSelectMain()

	concurrency.ConcurrencyExercise()
}

package main

import (
	"fmt"
	"log"

	"besimgurbuz.com/list"
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
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	myLinkedList := &list.List[string]{
		&list.List[string]{
			&list.List[string]{
				&list.List[string]{
					&list.List[string]{
						nil,
						"fifth item",
					},
					"fourth item",
				},
				"thirth item",
			},
			"second item",
		},
		"first item",
	}

	myLinkedList.Print()

	fmt.Println("REMOVING BY INDEX")
	_, err := myLinkedList.RemoveByIndex(0)

	if err != nil {
		log.Fatal(err)
	}

	myLinkedList.Print()

	fmt.Println("Removing by pointer")
	_, err = myLinkedList.Remove(myLinkedList)

	if err != nil {
		log.Fatal(err)
	}

	myLinkedList.Print()

	fmt.Println("Adding!")

	myLinkedList.Add(list.List[string]{
		nil,
		"ayo ayo ayo this is new item",
	}, 0)

	myLinkedList.Print()

	fmt.Println(myLinkedList.ConvertToSlice())
}

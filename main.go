package main

import (
	"requester/internal"
)

func main() {
	wc := &internal.Requester{
		NumOfRoutines: 10,
		Urls:          nil,
	}
	wc.Run()
}

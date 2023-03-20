package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

var (
	jsonData1 = []byte(`["1", "2", "3"]`)
	jsonData2 = []byte(`["4", "5", "6"]`)
	gData     []*string
	mutex     sync.RWMutex
)

func setNewData(newData []*string) {
	mutex.Lock()
	defer mutex.Unlock()
	gData = newData
}

func getData() []*string {
	mutex.RLock()
	defer mutex.RUnlock()
	// results := make([]*string, len(gData))
	// copy(results, gData)
	return gData
}

func main() {
	// 1. g1 call getData, and cache data[0]
	// 2. g2 call setData
	// 3. g1 call getData, and compare cache with data[0]

	a := make(chan int)
	b := make(chan int)

	d := []*string{}
	json.Unmarshal(jsonData1, &d)
	setNewData(d)
	fmt.Println("Orignal data", d)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		data1 := getData()
		a <- 1
		<-b
		data2 := getData()
		fmt.Printf("data1: %p and data2: %p\n", data1, data2)
		fmt.Printf("Two value %p: %v,   %p:%v\n", data1[0], *data1[0], data2[0], *data2[0])
	}()

	go func() {
		defer wg.Done()
		<-a
		d := []*string{}
		json.Unmarshal(jsonData2, &d)
		fmt.Println("Original data in g2", d)
		setNewData(d)
		fmt.Println(len(getData()), getData())
		fmt.Printf("Data in g2 %v\n", getData()[0])
		b <- 1
	}()

	wg.Wait()
}

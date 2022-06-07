package main

import "fmt"

type Data int

type Driver struct {
	PubData []Data
}

// bad
func (c *Driver) SetData1(d []Data) {
	c.PubData = d
}

func (c *Driver) SetData2(d []Data) {
	c.PubData = make([]Data, len(d))
	copy(c.PubData, d)
}

func main() {
	driver1 := &Driver{}

	d1 := []Data{0, 0, 0, 0, 0}

	driver1.SetData1(d1)

	// 对d1的修改会影响driver1的内部状态,反之亦然
	d1[0] = 5
	driver1.PubData[1] = 10
	fmt.Printf("d1 %v\n", d1)
	fmt.Printf("data in driver1 %v\n", driver1.PubData)

	driver2 := &Driver{}
	d2 := []Data{0, 0, 0, 0, 0}
	driver2.SetData2(d2)

	// 两者不会相互影响
	d2[0] = 5
	driver2.PubData[1] = 10

	fmt.Printf("d1 %v\n", d2)
	fmt.Printf("data in driver1 %v\n", driver2.PubData)
}

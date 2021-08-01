package main

import (
	"fmt"
)

type Address struct {
	village string
	state   string
	hno     int
}
type Customer struct {
	custname string
	sal      int
	Address
}

func main() {
	initialize()
}
func objreference(cust *Customer) {
	fmt.Println("cust name:", (*cust).custname, cust.sal, cust.Address)
}
func valueref(cust Customer) {
	fmt.Println("cust data", cust.custname)
}
func initialize() {
	cust1 := Customer{}
	cust1.custname = "subbu"
	cust1.sal = 13500
	cust1.hno = 3 - 24
	cust1.village = "kammapalli"
	cust1.state = "AndhraPradesh"
	fmt.Println(cust1)
	fmt.Println()
	add := Address{
		hno:     3 - 24,
		village: "kammapalli",
		state:   "AndhraPradesh",
	}
	cust2 := Customer{
		custname: "srinu",
		sal:      15000,
		Address:  add,
	}
	fmt.Println(cust2)
	objreference(&cust2)
	valueref(cust2)

	data := make(map[int]Customer)
	data[100] = cust1
	data[101] = cust1
	data[103] = cust2

	fmt.Println(data)
	
	var add1 = data[100]
	fmt.Println("cust info based on key 100: ", add1)
	
	delete(data, 100)
	fmt.Println("deleted key is 100: ", data)
	
	data[103] = cust1
	fmt.Println("updated key 103: ", data)
}

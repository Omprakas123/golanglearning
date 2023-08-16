/*Create Company struct Id,Name, Website, Telehone, Fax, Address // Address must be a promoted filed

Create Address struct City, Line1, Street, State, Country, PinCode , Map // Map must be a promoted field

Create Map struct Lan, Lat

-> Can you access Lan and Lap with the object of Company*/

package main

import "fmt"

type Company struct {
	Id, Fax, Teliphone int
	Name               string
	Website            string
	Address
}

type Address struct {
	city, Line1, Street, State, Country string
	PinCode                             int
	map1
}

type map1 struct {
	Lan, Lat int
}

func main() {

	object1 := Company{1, 5766, 575788, "Vs&CO", "VICTORIA.com", Address{"banglore", "34434", "manayata Tech park", "banglore", "India", 824123, map1{7676, 8889}}}

	obj1 := object1.Address.map1
	fmt.Println("the value of lan and lat", obj1.Lan, obj1.Lat)

}

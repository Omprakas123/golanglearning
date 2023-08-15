/* Q8-*Implement using switch case
State : Karnataka , AP, Delhi ,UP
Gender: M,F
Age: >0
Height:

|State     |Gender |Height|   Age |Ticket Status|
|----------|-------|------|-------|-------------|
|Karnataka |  F    |      |       | No ticket   |
|AP        |  F    |<110cm|  <5y  | No ticket   |
|Delhi     |  F    |      |       | No Ticket   |
|UP        |  F    |<120cm|  <6y  | No ticket   |
|Karnataka |  M    |<110cm|  <5y  | No ticket   |
|AP        |  M    |<110cm|  <5y  | No ticket   |
|Delhi     |  M    |<130cm|  <7y  | No Ticket   |
|UP        |  M    |<120cm|  <6y  | No ticket   |
|Other than the above table ,It is a full ticket|
*/

package main

import "fmt"

func main() {
	CheckBusTicketPerson("Karnataka", "F", 114, 8)

}
func CheckBusTicketPerson(State, Gender string, Height, Age int) {

	switch {

	case State == "Karnataka" && (Gender == "F" || Gender == "M") && Height < 110 && Age < 5:
		fmt.Println("This Guys are : -> Not Applicable for Bus Ticket")
	case State == "AP" && (Gender == "F" || Gender == "M") && Height < 110 && Age < 5:
		fmt.Println("This Guys are : -> Not Applicable for Bus Ticket")
	case State == "Delhi" && (Gender == "F" || Gender == "M") && Height < 130 && Age < 7:
		fmt.Println("This Guys are : -> Not Applicable for Bus Ticket")
	case State == "UP" && (Gender == "F" || Gender == "M") && Height < 120 && Age < 6:
		fmt.Println("This Guys are : -> Not Applicable for BusTicket")
	default:
		fmt.Println("This Guys are : ->  Applicable for Bus Ticket")

	}
}

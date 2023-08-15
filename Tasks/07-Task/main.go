/* Q7-> tate : Karnataka , AP, Delhi ,UP
Gender: M,F
Age: >0
Height:

            Gender  Height   Age   Ticket Status
Karnataka    F                     No ticket
AP           F      <110cm   <5y   No ticket
Delhi        F                     No Ticket
UP           F      <120cm   <6y   No ticket

Karnataka    M      <110cm    <5y  No ticket
AP           M      <110cm    <5y  No ticket
Delhi        M      <130cm    <7y  No Ticket
UP           M      <120cm    <6y  No ticket

Other than the above table , It is full ticket */

package main

import "fmt"

func main() {
	CheckTicket("Karnataka", "F", 105, 4)
	CheckTicket("AP", "M", 145, 26)
	CheckTicket("Delhi", "M", 106, 6)
	CheckTicket("UP", "M", 235, 30)
	CheckTicket("Karnatka", "M", 157, 20)

}

func CheckTicket(State, Gender string, Height, Age int) {

	if State == "Karnataka" && (Gender == "F" || Gender == "M") && Height < 110 && Age < 5 {

		fmt.Println("Ticket Status: -> No Ticket")

	} else if State == "AP" && (Gender == "F" || Gender == "M") && Height < 110 && Age < 5 {

		fmt.Println("Ticket Status: -> No Ticket")

	} else if State == "Delhi" && (Gender == "F" || Gender == "M") && Height < 130 && Age < 7 {

		fmt.Println("Ticket Status: -> No  Ticket")

	} else if State == "UP" && (Gender == "F" || Gender == "M") && Height < 120 && Age < 6{

		fmt.Println("Ticket Status: -> No Ticket")

	} else {
		fmt.Println("Ticket Status: -> Full Ticket")
	}

}

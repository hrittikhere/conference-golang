package main

import (
	"fmt"
)

func main() {
	const confTicket = 50
	var remainingTickets = 50
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var conferenceName = "GoLang Conf"

	fmt.Printf("%v of %v are available \n", remainingTickets, confTicket)
	fmt.Printf("Welcome to our %s Ticket Booking System \n", conferenceName)


	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}



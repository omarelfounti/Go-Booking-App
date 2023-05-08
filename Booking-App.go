package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("conferencetickets is %T, remainingTickets is %T, conferenceName is %T\n.", conferenceName, conferenceTickets, remainingTickets)
	
	fmt.printf("Welcome to %v booking application\n", conferenceName)
	fmt-Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.println("Get your tickets here to attend")

	
	var firstName string
	var lastName string
	var emailAdress string
	var numberOfTickets int

	fmt.Println("enter your first name: ")
	fmt.Scan(&firstName)

	
	fmt.Println("enter your last name: ")
	fmt.Scan(&lastName)


	
	fmt.Println("enter your email address: ")
	fmt.Scan(&emailAdress)


	fmt.Println("enter the number of tickets: ")
	fmt.Scan(&numberOfTickets)


	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("the whole array: %v\n.", bookings)
	fmt.Printf("the first value: %v\n.", bookings[0])
	fmt.Printf("Array type: %v\n.", bookings)
	fmt.Printf("Array length: %v\n.", len(bookings))

	fmt.Printf("Thank you %v %v for buying %v tickets. You will receive a confirmation in your email %v\n.", firstName, lastName, numberofTickets, emailAdress)
	fmt.println("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Println("These are the bookings from the application: %v\n.", bookings)


}


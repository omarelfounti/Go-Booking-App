package main

import (
	"fmt"
	"strings"
    )

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("conferencetickets is %T, remainingTickets is %T, conferenceName is %T\n.", conferenceName, conferenceTickets, remainingTickets)
	
	fmt.printf("Welcome to %v booking application\n", conferenceName)
	fmt-Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.println("Get your tickets here to attend")

	for {

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

	isValidName := len(firstName) >= 2 && len(lastName) >= 2 
	isValidEmail := strings.Contains(email, "@")
    isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

    if isValidName && isValidEmail && isValidTicketNumber{
		remainingTickets = remainingTickets - userTickets
	    bookings = append(bookings, firstName + " " + lastName)
		fmt.Printf("Thank you %v %v for buying %v tickets. You will receive a confirmation in your email %v\n.", firstName, lastName, numberofTickets, emailAdress)
	    fmt.println("%v tickets remaining for %v\n", remainingTickets, conferenceName)

        firstNames := []string{}
	    for index, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
 	   }
		
	} 
	fmt.Printf("The first names of bookings are: %v\n", firstNames)

    
	if remainingTickets == 0 {
		fmt.Println("conference is booked out.")
		break
	}else {
		if !isValidName{
			fmt.Println("first name or las name you entered is too short")
		}
		if !isValidEmail{
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber{
			fmt.Println("number of tickets you entered is invalid")
		}
        fmt.Printf("you input data is invalid. Try again")
	  }


   }

	
	
}


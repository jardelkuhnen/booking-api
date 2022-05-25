package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go conferece"
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings []string

	greetUsers(conferenceName, remainingTickets, conferenceTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter your first name: ")
		//&variabelName indica a pasagem do ponteiro de memoria onde sera alocado o valor que sera lido no input
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter quantity of tickets do you wanna: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v, tickets, so ou cann't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		isValidName, isValidEmail, isValidTicketNumber := helper.IsValidInputData(firstName, lastName, email, remainingTickets, userTickets)
		if !isValidName || !isValidEmail || !isValidTicketNumber {
			fmt.Println("Invalid data. Try again")
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)
		fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)

		printFirstNames(getFirstNames(bookings))

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come bak next year.")
			break
		}

	}

}

func printFirstNames(firstNames []string) {
	fmt.Printf("The first names of bookings are: %v\n", firstNames)

}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	/**
	_ usado para mostrar que o parametro nao sera usado.
	Em Go precisa se deixar explicito parametros que nao serao usados.
	*/
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func greetUsers(conferenceName string, remainingTickets int, conferenceTickets int) {
	fmt.Println("Seja bem vindo a", conferenceName, "booking app")
	fmt.Println("Temos no total", conferenceTickets, "tickets e", remainingTickets, "ainda estão disponiveis!")
}

package main

import (
	"BOOKING-APP/helper"
	"fmt"
	"time"
	"sync"
)

const TotalConferenceTickets int = 50
var conferenceName = "GO Conference"
var remainingConferenceTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg =  sync.WaitGroup{}

func main() {

	greetUsers()

	// fmt.Printf("conferenceName is %T, totoalConferenceTickets is %T, remainingConferenceTickets is %T ", conferenceName, TotalConferenceTickets, remainingConferenceTickets)

	fmt.Println("Get your tickets here to attend")

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingConferenceTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTicket(userTickets, firstName, lastName, email)
			// call function print first Name
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of booking are:%v\n", firstNames)

			if remainingConferenceTickets == 0 {
				// end Program
				fmt.Println("Our conference is booked out. Come back next year.")
				// break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address u entered doesn't contain @ sign")
			}
			if !isValidTickets {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Printf("Your input data is invalid try again\n")

		}
		wg.Wait()

	}


func greetUsers() {

	fmt.Printf("Welcome to %v booking application \n", conferenceName)

	fmt.Printf("We have total %v tickets and %v are still available\n", TotalConferenceTickets, remainingConferenceTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for thier name

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingConferenceTickets = remainingConferenceTickets - userTickets

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingConferenceTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("################")

	fmt.Printf("Sending ticket:\n %v\n to email address: %v\n", ticket, email)

	fmt.Println("################")

	wg.Done()

}

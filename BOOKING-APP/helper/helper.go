package helper

import "strings"

var Myvar = "somevalue"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint,remainingConferenceTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingConferenceTickets

	return isValidName, isValidEmail, isValidTickets
}

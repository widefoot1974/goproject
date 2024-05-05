package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers(conferenceTickets, remainingTickets)

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("\nOur conference is booked out. Come back next year.")
			//break
		}
	} else {
		// fmt.Printf(
		// 	"We only have %v tickets remaing, so you can't book %v tickets\n",
		// 	remainingTickets, userTickets)
		fmt.Println("[WARNING] Your input data is invalid, try again.")

		if !isValidName {
			fmt.Println("## first name or last name you entered is too short ##")
		}
		if !isValidEmail {
			fmt.Println("## email address you entered doesn't contain @ sign ##")
		}
		if !isValidTicketNumber {
			fmt.Println("## number of tickets you entered is invalid ##")
		}
	}

	wg.Wait()
}

func greetUsers(confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, userData := range bookings {
		// names := strings.Fields(booking)
		firstNames = append(firstNames, userData.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// // create a map for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf(
		"\n- Thank you (%v %v) for booking (%v) tickets. You will received a confirmation email at (%v)\n",
		firstName, lastName, userTickets, email)
	fmt.Printf(
		"- (%v) tickets remaining for %v\n",
		remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("[%v tickets for %v %v]", userTickets, firstName, lastName)
	fmt.Println("##################################################")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("##################################################")
	wg.Done()
}

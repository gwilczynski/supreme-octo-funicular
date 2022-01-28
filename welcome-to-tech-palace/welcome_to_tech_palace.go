package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func createBorder(numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return createBorder(numStarsPerLine) + "\n" + welcomeMsg + "\n" + createBorder(numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newStr := strings.Replace(oldMsg, "*", "", -1)
	newStr = strings.Replace(newStr, "\n", "", -1)

	return strings.Trim(newStr, " ")
}

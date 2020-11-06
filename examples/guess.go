// guess - it a game where player should guess random number between 1 and 100
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// get current date as seconds number
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Program picks a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	// bufio.reader to read from command prompt
	reader := bufio.NewReader(os.Stdin)
	// by default programm assumes that player lose
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Println("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// Deletion of new line simbol
		input = strings.TrimSpace(input)
		// input string converted to int number
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Ooops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Ooops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good game! Easy win for you!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess programs number. It was:", target)
	}
}

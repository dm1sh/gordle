package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type CharacterStatus byte

const (
	RIGHT CharacterStatus = iota
	CONTAINS
	WRONG
)

func main() {
	rightOutput := color.New(color.BgGreen, color.FgBlack)
	containsOutput := color.New(color.BgYellow, color.FgBlack)
	wrongOutput := color.New(color.BgWhite, color.FgBlack)

	var nFlag = flag.Int("n", 5, "Word size")
	flag.Parse()
	nChar := *nFlag

	fmt.Println("Welcome to Gordle - go implementation of Wordle game")

	file, err := os.Open(fmt.Sprintf("./dictionary/%d.txt", nChar))

	if err != nil {
		log.Fatal("There is no dictionary with such amount of letters")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	words := []string{}

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	fmt.Printf("You are going to guess a %d letter word in %d tries from %d dictionary\n", nChar, nChar+1, len(words))

	rand.Seed(time.Now().Unix())

	ind := rand.Intn(len(words))

	chosen := words[ind]

	inpScanner := bufio.NewScanner(os.Stdin)

	var nTries = 1

	for nTries <= nChar+1 {
		inpScanner.Scan()
		input := inpScanner.Text()

		moveConsoleCursorUp()

		if len(input) != nChar {
			fmt.Printf("Wrong number of letters in your input. Remember, you need to input %d character words.\n", nChar)
			continue
		}

		if input == chosen {
			color.Green(input)
			fmt.Printf("Congratulations, you have found the right word on %d try!\n", nTries)
			os.Exit(0)
		}

		if !BinSearch(words, input) {
			fmt.Printf("The word you entered: %s is not contained in the dictionary\n", input)
			continue
		}

		nTries++

		comparitionResult := CompareStrings(input, chosen)

		for pos, chr := range input {
			switch comparitionResult[pos] {
			case RIGHT:
				rightOutput.Print(string(chr))
			case CONTAINS:
				containsOutput.Print(string(chr))
			case WRONG:
				wrongOutput.Print(string(chr))
			}
		}

		fmt.Print("\n")
	}

	fmt.Println("Unfortunately, you lost. Puzzled word was", chosen)
}

func moveConsoleCursorUp() {
	fmt.Print("\033[1A")
}

func CompareStrings(input, chosen string) []CharacterStatus {
	inputReader := strings.NewReader(input)
	chosenReader := strings.NewReader(chosen)

	result := make([]CharacterStatus, len(input))

	for i := 0; ; i++ {
		inputRune, _, err1 := inputReader.ReadRune()
		chosenRune, _, err2 := chosenReader.ReadRune()

		if err1 != nil || err2 != nil {
			return result
		}

		if inputRune == chosenRune {
			result[i] = RIGHT
		} else if strings.ContainsRune(chosen, inputRune) {
			result[i] = CONTAINS
		} else {
			result[i] = WRONG
		}
	}
}

func BinSearch(arr []string, el string) bool {
	begin, end := 0, len(arr)-1

	for begin <= end {
		middle := begin + (end-begin)/2

		if el > arr[middle] {
			begin = middle + 1
		} else if el < arr[middle] {
			end = middle - 1
		} else {
			return true
		}
	}

	return false
}

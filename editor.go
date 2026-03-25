package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Color(text string) string {
	red := "\033[31m"
	reset := "\033[0m"
	text = red + text + reset
	return text
}

func LastChar(text string) string {
	return string(text[len(text)-1])
}

func Capitalize(text string) string {
	return strings.ToUpper(text)
}

func DeleteIndex(text string, index int) string {

	return text[:index] + text[index+1:]

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\t TEXT MANIPULATOR")
	fmt.Println()
	fmt.Println()

	var recode string

top:

	for {

		fmt.Print("🔸 INPUT A WORD : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(Color("please input a word"))
			fmt.Println()

			continue
		}

		digits, _ := regexp.MatchString(`\d`, input)
		punc, _ := regexp.MatchString(`[[:punct:]]`, input)

		if digits || punc {

			fmt.Println(Color("Error alphabetical format only"))
			fmt.Println()
			continue
		}

		recode = input

		break
	}
	fmt.Println()
	fmt.Println("\t OPERATION FUNCTIONS⚙️⚙️")
	fmt.Println()
	fmt.Println("🔸(lastChar) ➡️   Returns the last character of the word.")
	fmt.Println()
	fmt.Println("🔸(capitalize) ➡️   Returns the word with every letter capitalized.")
	fmt.Println()
	fmt.Println("🔸(deleteIndex) ➡️   Prompts the user to input an index. Deletes the character at that index and returns the result. If the index is greater than or equal to the length of the word, display an error message and restart from the beginning.")
	fmt.Println()
	fmt.Println()

bottom:
	fmt.Println()
	fmt.Println()
	fmt.Println("\t PICK A OPERATION TO USE ⚙️")
	fmt.Println()
	fmt.Print("🟢 FOR (lastChar) INPUT [1]", "\n", "🟢 FOR (capitalize) INPUT [2]", "\n", "🟢 FOR (deleteIndex) INPUT [3]", "\n")

	var operation string

	for {

		fmt.Print("\n", "🔸ENTER A OPERATION NUMBER : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(Color("Error.. Empty Operation"))
			fmt.Println()
			continue
		}

		operation = input

		switch operation {
		case "1":
			fmt.Println("(lastChar)⭐")
			fmt.Println()
		case "2":
			fmt.Println("(capitalize)⭐")
			fmt.Println()
		case "3":
			fmt.Println("(deleteIndex)⭐")
			fmt.Println()
		default:
			fmt.Println(Color("Error!! Operation Not Found"))
			fmt.Println()
			continue

		}

		break

	}

	if operation == "1" {
		result := LastChar(recode)
		fmt.Println("RESULT :", result, "✅")

	}

	if operation == "2" {
		result := Capitalize(recode)
		fmt.Println("RESULT :", result, "✅")
	}

	if operation == "3" {

		var index int

		for {

			fmt.Print("🔸INPUT A INDEX: ")

			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			input = strings.ReplaceAll(input, " ", "")

			if input == "" {
				fmt.Println(Color("Error Input can't be EMPTY"))
				fmt.Println()
				continue
			}

			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(Color("Index can't be in Alphabetical format"))
				fmt.Println()
				continue
			}

			if num < 0 {
				fmt.Println(Color("Error!! index can't be less than zero"))
				fmt.Println()
				continue
			}

			if num > len(recode) {
				fmt.Println(Color("Error.. Index shouldn't be greater than the length of the string"))
				fmt.Println()
				continue
			}

			index = num

			break
		}

		result := DeleteIndex(recode, index)
		fmt.Println("RESULT :", result, "✅")
	}

	fmt.Println()
	fmt.Println()

	var option string

	for {

		fmt.Print("🔴 DO YOU WANNA END THE PROGRAM? (YES/NO) : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(Color("please select an option"))
			fmt.Println()

			continue
		}

		option = Capitalize(input)

		switch option {
		case "YES":
			return
		case "NO":
			fmt.Println()
			fmt.Println("🟢 GREAT!!.. you can continue the program 🎊")
			fmt.Println()

		default:
			fmt.Println(Color("invalid option.. Choose between(YES/NO)"))
			fmt.Println()
			continue

		}

		break
	}

	if option == "NO" {

		fmt.Println()
		fmt.Println("🔸INPUT [1] TO USE PREVIOUS INPUT TO TEST ANOTHER OPERATION")
		fmt.Println("🔸INPUT [2] TO USE A NEW INPUT TO TEST THE OPERATOR")
		fmt.Println()

		var pick string

		for {

			fmt.Print("🔸ENTER YOUR OPTION NUMBER : ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "" {
				fmt.Println(Color("please enter an option.."))
				fmt.Println()
				continue
			}

			pick = input

			switch pick {
			case "1":
				goto bottom

			case "2":
				fmt.Println()
				fmt.Println()
				goto top

			default:
				fmt.Println(Color("Error invaild Option"))
				fmt.Println()

			}
		}
	}
}

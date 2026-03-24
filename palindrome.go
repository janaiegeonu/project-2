package main 


import (
	"fmt"
	"bufio"
	"os"
	"strings"
	
)

func cap(word string) string {
  
	word = strings.ToLower(word)

	left := 0
	right := len(word) -1

	for left < right {
		if word[left] != word[right] {
			return "NOT A PALINDROME"
		}
		left++
		right-- 
		
	}
	return "A PALINDROME"
}

	func main() {

		reader := bufio.NewReader(os.Stdin)

		fmt.Println("        PALINDROME CHECKER       ")
		fmt.Println()

			var tester string

     for {

		fmt.Print("Input text : ")

		input,_ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Error.. can't be EMPTY")
			fmt.Println()
			continue
		}

		tester = input

		break

	 }
	 fmt.Print("Result : ")
	 fmt.Println(cap(tester))

	}
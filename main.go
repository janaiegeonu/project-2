package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func Case(text string) string {
	parts := strings.Fields(text)

	for index, value := range parts{
		value = strings.ToLower(value)
		parts[index] = strings.ToUpper(value[:1]) + (value[1:])
	}
	return strings.Join(parts, " ")
}

func Upper(text string) string {

	return strings.ToUpper(text)
}

func main() {

	fmt.Println("🔸      ENYI OKW'OLA SPOT🍵🫗      🔸")
	fmt.Println()
	fmt.Println("-- Customer Details🪪")

	reader := bufio.NewReader(os.Stdin)

	var fullname string

	for {
		
		fmt.Print("\n","🔸 ENTER FULL NAME : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, _ := regexp.MatchString(`\d`, input)

		if num {
			fmt.Println("Error no numerical format in name")
			fmt.Println()
			continue
		}
		if input == "" {
			fmt.Println("please input a name")
			fmt.Println()
			continue
		}
		if !strings.Contains(input, " ") {
			fmt.Println("Error.. enter your lastname too")
			fmt.Println()
			continue
		}

		fullname = Case(input)

		break

	}
	fmt.Println()
	fmt.Println("🎊 Welcome to our store",fullname,"🎊")
	fmt.Println()


	var contact int

	for {

		fmt.Print("🔸 ENTER YOUR PHONE-NUMBER🔢 : (+234)")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Phone-number can't be empty")
			fmt.Println()
			continue
		}

		input = strings.ReplaceAll(input, " ","") 

		if len(input) != 10 {
			fmt.Println("Error.. Phone-number must be 10 digits only")
			fmt.Println()
			continue
		}

		digits, err := strconv.Atoi(input) 
		if err != nil {
			fmt.Println("Error.. Phone-number can't contain alphabetic formats")
			fmt.Println()
			continue
		}


		contact = digits

		break

	}

	fmt.Println(fullname, "your Phone-Number is (+234)", contact,"✅")
	fmt.Println()
	fmt.Println()


  fmt.Println("       AVALIBALE PRODUCTS 🍱📃 ")
  fmt.Println()
  fmt.Println(" 1. 🟢 ENYI OKW'OLA  ", "\n", "2. 🟢 ENYI OKW'OLA & BREAD ", "\n", "3. 🟢 ENYI OKW'OLA & HOT OKPA ", "\n")

	fmt.Println()
   fmt.Println("       PRODUCT PRICE(₦)🍱 💹  ")
  fmt.Println()
  fmt.Println(" 1. 🟢 ENYI OKW'OLA = ₦500  ", "\n", "2. 🟢 ENYI OKW'OLA & BREAD = ₦1200", "\n", "3. 🟢 ENYI OKW'OLA & HOT OKPA = ₦1000", "\n")
  fmt.Println("\n","        ⭐PLACE YOUR ORDER⭐       ","\n")
  fmt.Println(" ▫️ Input [1] to Buy ENYI OKW'OLA ", "\n", "▫️ Input [2] to Buy ENYI OKW'OLA & BREAD ", "\n", "▫️ Input [3] to Buy ENYI OKW'OLA & HOT OKPA","\n")
  
    var order string
	var amount int

start :

	for {
		fmt.Print("🔸 🔢 ORDER OPTION : ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("Error.. Can't be EMPTY!", "\n")
		continue
	}

		order = input

		switch order {
		case "1" :
			fmt.Println("Great choice💫!!.. (ENYI OKW'OLA)","\n")

		case "2" :
			fmt.Println("Great choice💫!!.. (ENYI OKW'OLA & BREAD)","\n")

		case "3" :
			fmt.Println("Great choice💫!!.. (ENYI OKW'OLA & HOT OKPA)","\n")

		default :
			fmt.Println("Error.. INVALID Option!","\n") 
				continue
			
		}

		break
	}

	for {

		fmt.Print("🔸 🧮 QUANTITY OF PRODUCT : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ReplaceAll(input, " ","")

		if input == "" {
			fmt.Println("Error.. Can't be empty","\n")
			continue
		}

		quant, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error.. Numeric format only", "\n")
			continue
		}

		if quant <= 0 {
			fmt.Println("Error.. quantity must be more than Zero", "\n")
			continue
		}

		if quant > 50 {
			fmt.Println("can't deliver huge order in one delivery!!","\n")
			continue
		}
		
		amount = quant
		
		break
	}


	

	if order == "1" {
		fmt.Print("\n","💵 Price = ₦")
		price := (500 * amount)
		fmt.Println(price)
		fmt.Println("\n")
	}

	if order == "2" {
		fmt.Print("\n","💵 Price = ₦")
		price := (1200 * amount)
		fmt.Println(price)
		fmt.Println("\n")

	}

	if order == "3" {
		fmt.Print("\n","💵 Price = ₦")
		price := (1000 * amount)
		fmt.Println(price)
		fmt.Println("\n")
	}

	var reorder string

	for {

	fmt.Print("DO YOU WANNA PLACE ANOTHER OTHER?(YES/NO) : ")
	

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("Error.. Can't be empty","\n")
			continue
	}

	reorder = Upper(input)

	switch reorder {

	case "YES": 
	   fmt.Println("\n", "Okay you can order Again!!🎊", "\n")
	   goto start

	case "NO":
		fmt.Println("\n", "Okay..",fullname," THANKS FOR YOUR PATRONAGE ","\n")
		fmt.Println("Delivering 1 hrs from NOW..🛵📦")


	}

	break

	}
}

	
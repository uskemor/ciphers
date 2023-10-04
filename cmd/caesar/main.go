package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	var plain string
	var shiftStr string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input text: ")
	scanner.Scan()
	plain = scanner.Text()
	for {
		fmt.Print("Right shift of ?(1-25): ")
		scanner.Scan()
		shiftStr = scanner.Text()
		shift, err := strconv.Atoi(shiftStr)
		if err != nil {
			fmt.Printf("Invalid input: %s\n", err)
		} else if shift < 1 || shift > 25 {
			fmt.Println("Right shift must be integer between 1 and 25.")
		} else {
			fmt.Printf("Encrypted text: %s\n", encrypt(plain, shift))
			break
		}
	}
}

func encrypt(plain string, shift int) (encrypted string) {
	encrypted = ""
	for _, c := range plain {
		var e rune
		if unicode.IsLetter(c) {
			if unicode.IsLower(c) {
				e = rune((int(c)-'a'+shift)%26 + 'a')
			} else {
				e = rune((int(c)-'A'+shift)%26 + 'A')
			}
		} else {
			e = c
		}
		encrypted += string(e)
	}
	return
}

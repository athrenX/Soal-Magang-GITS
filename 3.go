package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isBalanced(s string) string {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	fmt.Print("Masukkan string bracket: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		input = strings.ReplaceAll(input, " ", "") 
		fmt.Println(isBalanced(input))
	}
}

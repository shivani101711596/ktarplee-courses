package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ops = []string{"+", "-", "*", "/"}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		puzzleStr := strings.Split(scanner.Text(), " ")
		puzzle := make([]int, len(puzzleStr))
		for i, numStr := range puzzleStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Fprintln(os.Stderr, err) 
				os.Exit(1)
			}
			puzzle[i] = num
		}

		var solutions []string
		dp(puzzle, 2, puzzle[0], strconv.Itoa(puzzle[0]), &solutions)

		if len(solutions) == 0 {
			fmt.Println()
		} else {
			fmt.Println(strings.Join(solutions, ", "))
		}
	}
}

func dp(puzzle []int, index int, value int, s string, solutions *[]string) {
	if index == len(puzzle) {
		if value == puzzle[len(puzzle)-1] {
			*solutions = append(*solutions, s+" = "+strconv.Itoa(value))
		}
	} else if index <= len(puzzle) {
		var ele1, ele2 int
		if index == 2 {
			ele1 = puzzle[0]
			ele2 = puzzle[1]
		} else {
			ele1 = value
			ele2 = puzzle[index-1]
		}
		for _, op := range ops {
			var val int
			switch op {
			case "+":
				val = ele1 + ele2
			case "-":
				val = ele1 - ele2
			case "*":
				val = ele1 * ele2
			case "/":
				if ele2 != 0 && ele1%ele2 == 0 {
					val = ele1 / ele2
				} else {
					continue
				}
			}
			dp(puzzle, index+1, val, s+" "+op+" "+strconv.Itoa(ele2), solutions)
		}
	}
}


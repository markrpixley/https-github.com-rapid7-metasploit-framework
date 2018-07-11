package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processStack(e []string) float64 {
	result := 0.0
	for _, v := range e {
		c := strings.Split(v, " ")
		switch c[1] {
		case "*":
			num1, err := strconv.ParseFloat(c[0], 64)
			if err != nil {
				fmt.Println(err)
			}
			num2, err := strconv.ParseFloat(c[2], 64)
			if err != nil {
				fmt.Println(err)
			}
			result = num1 * num2
		case "/":
			num1, err := strconv.ParseFloat(c[0], 64)
			if err != nil {
				fmt.Println(err)
			}
			num2, err := strconv.ParseFloat(c[2], 64)
			if err != nil {
				fmt.Println(err)
			}
			if num2 == 0 {
				fmt.Println("error: you tried to divide by zero.")
				return 0.0
			}
			result = num1 / num2
		case "+":
			num1, err := strconv.ParseFloat(c[0], 64)
			if err != nil {
				fmt.Println(err)
			}
			num2, err := strconv.ParseFloat(c[2], 64)
			if err != nil {
				fmt.Println(err)
			}
			result = num1 + num2
		case "-":
			num1, err := strconv.ParseFloat(c[0], 64)
			if err != nil {
				fmt.Println(err)
			}
			num2, err := strconv.ParseFloat(c[2], 64)
			if err != nil {
				fmt.Println(err)
			}
			result = num1 - num2
		}
	}
	return result
}

func main() {
	expressions := make([]string, 0)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("gocalc>")
		for scanner.Scan() {
			expressions = append(expressions, scanner.Text())
			res := processStack(expressions)
			fmt.Println(res)
			fmt.Print("gocalc>")
		}
	}
}


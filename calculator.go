package mycalculator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (c calc) parser(entry string) (int, error) {
	value, errorParse := strconv.Atoi(entry)
	return value, errorParse
}

func (c calc) operate(input string, operator string) (int, error) {
	values := strings.Split(input, operator)
	valueA, errorParseA := c.parser(values[0])
	valueB, errorParseB := c.parser(values[1])

	if errorParseA != nil {
		return 0, errorParseA
	}

	if errorParseB != nil {
		return 0, errorParseB
	}

	switch operator {
	case "+":
		return valueA + valueB, nil
	case "-":
		return valueA - valueB, nil
	case "*":
		return valueA * valueB, nil
	case "/":
		return valueA / valueB, nil
	default:
		log.Println(operator, "operation is not supported!")
		return 0, nil
	}
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func processResult(input string, operator string) {
	c := calc{}
	value, err := c.operate(input, operator)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result of", input, "equals to", value)
	}
}

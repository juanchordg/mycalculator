package mycalculator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Calc struct
type Calc struct{}

func (c Calc) parser(entry string) (int, error) {
	value, errorParse := strconv.Atoi(entry)
	return value, errorParse
}

// Operate es usado para calcular las operaciones aritméticas
func (c Calc) Operate(input string, operator string) (int, error) {
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

//ReadInput lee una entrada
func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

//ProcessResult usa los datos de la entrada y los procesa
func ProcessResult(input string, operator string) {
	c := Calc{}
	value, err := c.Operate(input, operator)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result of", input, "equals to", value)
	}
}

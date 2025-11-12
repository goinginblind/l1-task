package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"os"
	"unicode/utf8"
)

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).
//
// Комментарий: в Go тип int справится с такими числами,
// но обратите внимание на возможное переполнение для ещё больших значений.
// Для очень больших чисел можно использовать math/big.

// Operation holds the 'big.Float's to perform an operation on,
// which is too, stored in this same struct - op field of type rune.
type Operation struct {
	a, b *big.Float
	op   rune
}

// ParseFromLine splits the given input, creating an operation.
// If the input contains too few arguments, it returns an ErrNotEnoughArgs.
// If the input couldn't be parsef into big.Float, it returns ErrNonNumericArg.
func ParseFromLine(input []byte) (*Operation, error) {
	tokens := bytes.Fields(input)
	if len(tokens) < 3 {
		return nil, ErrNotEnoughArgs
	}

	a, b := new(big.Float), new(big.Float)
	a, okA := a.SetString(string(tokens[0]))
	b, okB := b.SetString(string(tokens[2]))
	if !okA || !okB {
		return nil, ErrNonNumericArg
	}

	op, _ := utf8.DecodeRune(tokens[1])
	return &Operation{
		a:  a,
		b:  b,
		op: op,
	}, nil
}

var (
	ErrUnsupportedOp = errors.New("unsupported operation")
	ErrNonNumericArg = errors.New("at least one of the arguments contains non-numeric characters")
	ErrNotEnoughArgs = errors.New("too few arguments")
)

// GetResult performs the operation itself.
func GetResult(o *Operation) (*big.Float, error) {
	res := new(big.Float).Copy(o.a)
	switch o.op {
	case add:
		return res.Add(o.a, o.b), nil
	case sub:
		return res.Sub(o.a, o.b), nil
	case div:
		return res.Quo(o.a, o.b), nil
	case mult:
		return res.Mul(o.a, o.b), nil

	default:
		return nil, fmt.Errorf("fail performing '%c': %w", o.op, ErrUnsupportedOp)
	}
}

const (
	useMsg = "use: [VALUE1] [OPERATION: +, -, / or *] [VALUE2]"

	add, sub, div, mult = '+', '-', '/', '*'
	quit                = "quit"
)

// RunCalculator scans user input and displays the results - either for
// performed operations or for the errors. Also allows the user to quit
// without a <C-c> sigint - via 'quit' cmd.
func RunCalculator() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput := scanner.Bytes()
		if bytes.HasPrefix(userInput, []byte(quit)) {
			return
		}
		op, err := ParseFromLine(userInput)
		if err != nil {
			fmt.Printf("%s\n%s\n\n", err, useMsg)
			continue
		}
		res, err := GetResult(op)
		if err != nil {
			fmt.Printf("%s\n%s\n\n", err, useMsg)
			continue
		}
		fmt.Println(res.Text('g', 10))
	}
}

func main() {
	fmt.Println("Hi! This simple single-operation calculator supports: ")
	fmt.Println("addition, subtraction, division and multiplication for pretty big numbers!")
	fmt.Println("Example use: ", useMsg)
	RunCalculator()
	fmt.Println("Bye!")
}

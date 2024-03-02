package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("filename required!")
		os.Exit(1)
	}

	code, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("cant read source file:", err.Error())
		os.Exit(1)
	}

	in := bufio.NewReader(os.Stdin)

	var mem [30_000]byte

	var memIx int

	var execStack []int

	for pos := 0; pos < len(code); pos++ {
		sym := code[pos]

		switch sym {
		case '>': 
			memIx++
		case '<':
			memIx--
		case '+':
			mem[memIx]++
		case '-':
			mem[memIx]--
		case '.':
			fmt.Printf("%c", mem[memIx])
		case ',':
			data, err := in.ReadByte()
			if err != nil {
				fmt.Println("runtime error - cant read from in:", err.Error())
				os.Exit(1)
			}

			mem[memIx] = data
		case '[':
			execStack = append(execStack, pos-1)
			if mem[memIx] != 0 {
				continue
			}

			for len(execStack) > 0 {
				pos++
				if code[pos] == ']' {
					execStack = execStack[:len(execStack) - 1]
				} else if code[pos] == '[' {
					execStack = append(execStack, pos-1)
				}
			}

		case ']':
			if len(execStack) == 0 {
				fmt.Println("runtime error - unexpected jump operation at", pos + 1)
				os.Exit(1)
			}

			if mem[memIx] != 0 {
				pos = execStack[len(execStack) - 1]
			}

			execStack = execStack[:len(execStack) - 1]
		}
	}

}

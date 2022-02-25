package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//s := "++++++++++[->++++++++++<]>."
	s := ">,<+++[->+++<]>."

	header()
	body(s)
	footer()
}

func header() {
	fmt.Println("package main")
	fmt.Println("import (\"fmt\")")
	fmt.Println("import (\"os\")")
	fmt.Println("func main(){")
	fmt.Println("ptr:=0")
	fmt.Println("var mem [10]int")
	fmt.Println("b := []byte{1}")

}

/*
 */

func body(code string) {
	cs := trim(code)
	result := ""
	for _, c := range cs {
		switch c {
		case "+":
			result += "mem[ptr]++\n"
		case "-":
			result += "mem[ptr]--\n"
		case "<":
			result += "ptr--\n"
		case ">":
			result += "ptr++\n"
		case "[":
			result += "for mem[ptr] != 0 {\n"
			//result += "fmt.Printf(\"%c\",mem[ptr])\n"
		case "]":
			result += "}\n"
		case ".":
			result += "fmt.Printf(\"%c\",mem[ptr])\n"
		case ",":
			result += "_, err := os.Stdin.Read(b)\n"
			result += "mem[ptr] = 0\n"
			result += "if err == nil {\n"
			result += "mem[ptr] = int(b[0])\n   }\n"
			//result += "fmt.Scanf(\"%c\",&(mem[ptr]))\n"
		default:
			result += fmt.Sprintf("!!!%s!!!\n", c)
		}
	}
	fmt.Println(result)
}

func footer() {
	fmt.Println("}")
}

func trim(code string) []string {
	result := []string{}
	for _, r := range code {
		if contains(r) {
			result = append(result, string(r))
		}
	}
	return result
}

func contains(x rune) bool {
	for _, r := range "+-<>[].," {
		if x == r {
			return true
		}
	}
	return false
}

func readLines() string {
	str := ""
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		str += s.Text()
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
	return str
}

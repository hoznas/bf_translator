package main
import "fmt"
import "bufio"
import "os"
import "log"

func main(){
	//s := "++++++++++[->++++++++++<]>."
	s := ">,<+++[->+++<]>."
	
	header()
	body(s)
	footer()
}


func header(){
	fmt.Println("package main")
	fmt.Println("import \"fmt\"")
	fmt.Println("func main(){")
	fmt.Println("ptr:=0")
	fmt.Println("var mem [10]int")
	
}

func body(code string){
	cs := trim(code)
	result := ""
	for _,c := range cs {
		switch c{
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
			result += "fmt.Println(ptr,mem)\n"
		case "]":
			result += "}\n"
		case ".":
			result += "fmt.Printf(\"%c\",mem[ptr])\n"
		case ",":
			result += "fmt.Scanf(\"%c\",&(mem[ptr]))\n"
		default:
			result += fmt.Sprintf("!!!%s!!!\n",c)
		}
	}
	fmt.Println(result)
}

func footer(){
	fmt.Println("}")
}


func trim(code string) []string{
	result := []string{}
	for _,r := range code {
		if contains(r) {
			result = append(result,string(r))
		}
	}
	return result
}

func contains(x rune) bool{
	for _,r := range "+-<>[].," {
		if x == r {
			return true
		}
	}
	return false
}

func readLines()string{
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


package main
import (
"bufio"
"fmt"
"os"
"strings"
)
func main() {

fmt.Print("Enter your name: ")

reader := bufio.NewReader(os.Stdin)
name, _ := reader.ReadString('\n')

name = strings.TrimSpace(name)

fmt.Println("Hello, " + name + "!")

var age int
fmt.Print("Enter your age: ")

fmt.Scanf("%d", &age)
fmt.Println("You are", age, "years old.")
fmt.Println("Next year, you will be", age + 1)
}
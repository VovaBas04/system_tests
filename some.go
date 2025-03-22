package all

import "fmt"

var a = func() int {
	fmt.Println("var") //Исполняется в первую очередь
	return 0
}()

func init() {
	fmt.Println("init") //Исполняется во вторую очередь
}
func main() {

	fmt.Println("main") //Исполняется в последнюю очередь
}

package main
import (
	"fmt"
)
func  main()  {
	var a int = 100
	var b int = 200
	b, a = a, b
	fmt.Println(a, b)

}

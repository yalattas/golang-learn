package main

import (
	"fmt"
	"myproject/module1"
	"myproject/module2"
	"myproject/moduleLast"
)

func main() {
	fmt.Println("Main program starting...")
	module1.Module1Function()
	module2.Module2Function()
	moduleLast.ModuleTest()
}

package arrange

import "fmt"

func moveZeroes() {

}

func RunArrFunc(funcName string) {

	functions := make(map[string]func())
	functions["moveZeroes"] = moveZeroes

	if function, ok := functions[funcName]; ok {
		function()
		return
	}

	fmt.Printf("function %v doesn't exist\n", funcName)
}

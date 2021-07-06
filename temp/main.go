package main

import (
	"fmt"

	"github.com/deanveloper/gridspech-go"
	"github.com/deanveloper/gridspech-go/solve"
)

func main() {
	const levelAaa = `
[   ] [   ] [   ] [   ] [   ] [   ] [   ] [   ] 
[   ] [g /] [   ] [   ] [gA/] [   ] [   ] [   ] 
[   ] [   ] [   ] [   ] [   ] [   ] [   ] [   ] 
[   ] [   ] [   ] [   ] [   ] [ A/] [   ] [   ] 
[   ] [   ] [   ] [   ] [   ] [g  ] [   ] [   ] 
[   ] [   ] [   ] [   ] [   ] [   ] [   ] [   ] 
[   ] [   ] [   ] [   ] [   ] [   ] [   ] [   ] 
[   ] [   ] [g /] [   ] [   ] [   ] [   ] [   ] 
`
	grid := gridspech.MakeGridFromString(levelAaa)
	ch := solve.Goals(solve.Grid{Grid: grid}, 2)
	for solvedGrid := range ch {
		fmt.Println(solvedGrid)
		fmt.Println("=============")
	}
}
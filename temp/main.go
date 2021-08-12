package main

import (
	"fmt"

	"github.com/deanveloper/gridspech-go"
	"github.com/deanveloper/gridspech-go/solve"
)

func testDotsLevel() {
	const level = `
[1   ] [1   ] [2   ] [1   ] [1   ]
[1   ] [1   ] [2   ] [1   ] [1   ]
[1   ] [1   ] [    ] [1   ] [1   ]
[    ] [2   ] [1   ] [2   ] [    ]
[2   ] [1   ] [2   ] [1   ] [2   ]
`
	grid := gridspech.MakeGridFromString(level, 2)

	solver := solve.NewGridSolver(grid)
	for solution := range solve.Dots(solver) {
		newGrid := grid.Clone()
		newGrid.ApplyTileSet(solution)
		fmt.Println(newGrid)
		fmt.Println()
	}
}

func pathTest() {

	const levelAaa = `
[    ] [    ] [    ] [    ] [    ] [    ]
[    ] [gA/ ] [    ] [    ] [gA/ ] [    ]
[    ] [ A/ ] [    ] [    ] [ A/ ] [    ]
[    ] [ A/ ] [    ] [  / ] [    ] [ A/ ]
[    ] [ A/ ] [    ] [  / ] [    ] [gA/ ]
[    ] [ A/ ] [    ] [    ] [    ] [    ]
[    ] [    ] [ A/ ] [    ] [    ] [    ]
[    ] [    ] [gA/ ] [    ] [    ] [    ]
`
	grid := gridspech.MakeGridFromString(levelAaa, 2)
	ch := solve.Goals(solve.NewGridSolver(grid))
	for solvedGrid := range ch {
		fmt.Println(solvedGrid.MultiLineString())
		fmt.Println("=============")
	}
}

func main() {
	testDotsLevel()
}

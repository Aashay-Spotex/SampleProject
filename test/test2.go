package main

import (
	"sampleProject/ritesh"
)

func main() {
	ritesh.Test("Ritesh")

	// Blank Identifier
	_ = ritesh.Test1("Pratik")
	//y := ritesh.Test1("Pratik")

	// Anonymous Function
	x := func() {
		ritesh.Test1("Aashay")
	}
	x()
}

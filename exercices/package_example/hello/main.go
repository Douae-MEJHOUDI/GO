package main

import (
	"log"
	"package_example/hello/mathutils"
	"simplemath/utils"

	"github.com/pkg/math"
)

func main() {
	log.Println("Hello World")

	a, b := 10, 20
	log.Println(mathutils.Add(a, b))
	log.Println(mathutils.Multiply(a, b))
	log.Println(utils.Square(a))
	log.Println(math.Max(a, b))

}

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
	res, err := mathutils.Devide(5, 0)
	if err != nil {
		log.Println(err)
		log.Fatal("devision was not good")
	} else {
		log.Println(res)
	}

	log.Println(mathutils.Devide(5, 2))

}

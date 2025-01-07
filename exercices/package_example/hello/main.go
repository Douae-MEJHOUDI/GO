package main

import (
	"errors"
	"fmt"
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
	} else {
		log.Println(res)
	}

	str, err := formatDivision(5, 0)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(str)
	}
	log.Println(mathutils.Devide(5, 2))

}

func formatDivision(a, b int) (string, error) {
	res, err := mathutils.Devide(a, b)
	if err != nil {
		return "", errors.New("error deviding :" + err.Error())
	}
	return fmt.Sprintf("%d / %d = %f", a, b, res), nil
}

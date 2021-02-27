package main

import (
	"fmt"
	"math"
	"time"
)

func imprime(any interface{}) {
	fmt.Println(any)
}

func vals() (int, float32, string) {
	return 1, 0.1, "oi"
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func fakeAbreFechaArquivo() {
	fmt.Println("abre")
	defer func() { fmt.Println("fechado") }()

	fmt.Println("escreve")
	fmt.Println("apaga")
}

func main() {
	// hello world
	fmt.Println("### hello world")
	fmt.Println("hello world")

	// arithmetic
	fmt.Println("\n### arithmetic")
	fmt.Println(1 + 1)
	fmt.Println(1 * 2)
	fmt.Println(2 - 1)
	fmt.Println(3 / 2)
	fmt.Println(53 % 7)
	fmt.Println(math.Sqrt(5))

	// string
	fmt.Println("\n### string")
	a := "geleia"
	fmt.Println(a + " talk")
	fmt.Println(a[3])
	fmt.Println(string(a[3]))

	// array's
	fmt.Println("\n### array's")
	l := []int{1, 2, 3, 4, 5}
	for i, v := range l {
		fmt.Println(i, v)
	}

	mat := [][]float64{
		[]float64{0.1, 0.2, 0.3},
		[]float64{0.4, 0.5, 0.6},
		[]float64{0.7, 0.8, 0.9},
	}
	fmt.Println(mat)

	// struct
	fmt.Println("\n### struct")
	type point struct {
		x, y float64
	}
	p1 := point{x: 0, y: 0}
	p2 := point{
		x: 1,
		y: 1,
	}
	fmt.Println(p1, p2)

	// for
	fmt.Println("\n### for")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// if-else
	fmt.Println("\n### if-else")
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// switch
	fmt.Println("\n### switch")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	/// map
	fmt.Println("\n### map")
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))
	delete(m, "k2")
	fmt.Println("map:", m)

	// func
	fmt.Println("\n### func")
	imprime(1)
	imprime("teste")
	imprime(map[string]string{"foo": "bar"})
	imprime(true)
	v1, v2, v3 := vals()
	imprime(v1)
	imprime(v2)
	imprime(v3)
	imprime(func() int { return 8 })
	imprime(func() int { return 8 }())

	// methods
	fmt.Println("\n### methods")
	c := circle{radius: 5}
	fmt.Println(c.area())
	fmt.Println(c.perim())

	// defer
	fmt.Println("\n### defer")
	fakeAbreFechaArquivo()
}

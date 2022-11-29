package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
)

// type triangle struct {
// 	size int
// }

// type coloredTriangle struct {
// 	triangle
// 	color string
// }

// type upperstring string

// func (s upperstring) Upper() string {
// 	return strings.ToUpper((string(s)))
// }

//	func (t triangle) perimeter() int {
//		return t.size * 3
//	}
//
//	func (t coloredTriangle) perimeter() int {
//		return t.size * 3 * 2
//	}
//
//	func (t *triangle) doubleSize() {
//		t.size *= 2
//	}
func main() {
	// t := tringle{size: 3}
	// t.doubleSize()
	// fmt.Println("size:", t.size)
	// fmt.Println("Perimeter:", t.perimeter())
	// s := upperstring("Learning Go!")
	// fmt.Println(s)
	// fmt.Println(s.Upper())

	// t := geometry.Triangle{}
	// t.SetSize(4)
	// fmt.Println("Perimeter", t.Perimeter())

	// t := geometry.Triangle{}
	// t.SetSize(3)
	// fmt.Println("Size", t.size)
	// fmt.Println("Perimeter", t.Perimeter())

	// var s Shape = Square{3}
	// printInformation(s)

	// c := Circle{6}
	// printInformation(c)
	// rs := Person{"JON", "USA"}
	// var ab Stringer = Person{"VIA", "KOREA"}
	// fmt.Printf("%s\n%s\n", rs, ab)

	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println(resp)
	var writer customWriter = customWriter{}
	io.Copy(writer, resp.Body)
}

type customWriter struct{}
type GithubResponse []struct {
	FullName string `json:"full_name"`
	NodeId   string `json:"node_id"`
}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GithubResponse
	json.Unmarshal(p, &resp)
	fmt.Println(p)
	for _, r := range resp {
		fmt.Println(r.FullName, r.NodeId)
	}
	return len(p), nil
}

func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	size float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Stringer interface {
	String() string
}

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

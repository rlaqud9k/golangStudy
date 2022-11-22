package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>heqwerwewllo</bofddy></html>\n")
}

// 리턴값 명시적으로 구분가능
func sum(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	result = int1 + int2
	result2 := 34
	fmt.Println(result2)
	return
}

func deferTest() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular1", i)
	}
}

// 포인터 &주소 *주소에 있는 값 엑세스
func updateName(name *string) {
	*name = "David"
}

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {

	employee, err := getInformation(1000)
	if errors.Is(err, ErrorNotFound) {
		fmt.Printf("NOT FOUND: %v\n", err)
	} else {
		fmt.Print(employee)
	}
}

var ErrorNotFound = errors.New("Employee not found!")

func getInformation(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrorNotFound
	}
	employee := Employee{LastName: "die", FirstName: "fdsfd"}
	return &employee, nil
	// for tries := 0; tries < 3; tries++ {
	// 	_, err := apiCallEmployee(1000)
	// 	if err != nil {
	// 		// return nil, err
	// 		return nil, fmt.Errorf("GOT an error when getting the employee", err)
	// 	}
	// 	fmt.Println("Server is not responding, retrying ...")
	// 	time.Sleep(time.Second * 2)
	// }

	// return nil, fmt.Errorf("server has failed to respond to get the employee information")
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "eee", FirstName: "John"}
	return &employee, nil
}

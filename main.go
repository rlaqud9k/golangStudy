package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/rlaqud9k/golangStudy/calculator"
	"rsc.io/quote"
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

// 포인터 &주소 *주소에 있는 값 엑세스
func updateName(name *string) {
	*name = "David"
}

func main() {
	// var httpServer http.Server
	// http.HandleFunc("/", handler)
	// log.Println("start http listening :18888")
	// httpServer.Addr = ":18888"
	// log.Println(httpServer.ListenAndServe())
	// fmt.Println(sum("2", "54"))
	// firstName := "John"
	// updateName(&firstName)
	// fmt.Println(firstName)
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)
	total2 := calculator.Sum(3, 1113)
	fmt.Println(total2)
	fmt.Println("Version: ", calculator.Version)
	fmt.Println(quote.Hello())
}

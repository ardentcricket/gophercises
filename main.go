package main

import (
	// "io"
	"fmt"
)

// type MyStringData struct {
// 	str string
// 	readIndex int
// }

// func (myStringData *MyStringData) printSomething(p []byte, words string) {
// 	fmt.Println("thing thang: ", p, words)
// }

// func (myStringData *MyStringData) Read(p []byte) (n int, err error) {
// 	strBytes := []byte(myStringData.str)

// 	if myStringData.readIndex >= len(strBytes) {
// 		return 0, io.EOF 
// 	}

// 	fmt.Println("IN here")

// 	return
// }

func main() {
	var variable []int
	a := []int{1,2,3,4,5,6,7,8,9}
	s := a[2:4]
	newS := append(s, 55,66)
	fmt.Println(variable == nil, cap(a), a, newS)
	// src := MyStringData{str: "Hello Amazing World",}
	// src.printSomething([]byte(src.str), src.str)
}




















// func handleHello(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "hello world - nn \n")
// }

// func main() {
// 	http.HandleFunc("/hello", handleHello)
// 	var b bytes.Buffer
// 	fmt.Fprintf(&b, "It is number: %d, This is a string: %v", 10, "Bridge")
// 	b.WriteString("[DONE]")
// 	fmt.Println("The string buffer output is",b.String(), b.Len())
// 	log.Println("Listening @ localhost:8000/hello")
// 	http.ListenAndServe(":8000", nil)
// }
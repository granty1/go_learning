package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	file, err := os.OpenFile("/Users/Grant/go/src/github.com/granty1/go_learning/command/a.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
	fmt.Println(err)
}
	defer file.Close()
	//n, _ := fmt.Fprintf(file, "Grant")

	str := "Grant 22"
	var grant Person
	_, _ = fmt.Sscanf(str, "%s %d", &grant.Name, &grant.Age)
	resp, err := http.Get("http://c.m.163.com/nc/article/headline/T1348647853363/0-40.html")
	if err != nil {
		os.Exit(0)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	for i := 0; i < 20; i++ {
		_, _ = file.WriteString(string(body))
	}
}

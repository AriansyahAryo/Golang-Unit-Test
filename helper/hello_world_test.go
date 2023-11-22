package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Aryo")
	if result != "Hello Aryo" {
		t.Fail()
		//panic("Result is not Hello Aryo")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldiky(t *testing.T) {
	result := HelloWorld("iky")
	if result != "Hello Risky" {
		t.FailNow()
		//panic("Result is not Hello Risky")
	}
	fmt.Println("TestHelloWorldiky Done")
}

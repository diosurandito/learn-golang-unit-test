package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldDio(t *testing.T) {
	result := HelloWorld("Dio")
	if result != "Hello Dio" {
		// unit test failed
		t.Error("Result must be 'Hello Dio'")
	}
	fmt.Println("TestHelloWorldDio Done")
}

func TestHelloWorldDito(t *testing.T) {
	result := HelloWorld("Dito")
	if result != "Hello Dito" {
		// unit test failed
		t.Fatal("Result must be 'Hello Dito'")
	}
	fmt.Println("TestHelloWorldDito Done")
}

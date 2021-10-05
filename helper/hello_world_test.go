package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Dio")
	if result != "Hello Dio" {
		// unit test failed
		panic("Result is not Hello Dio")
	}
}

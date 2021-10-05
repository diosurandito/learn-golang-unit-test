package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Dio")
	require.Equal(t, "Hello Dio", result, "Result must be 'Hello Dio'")
	fmt.Println("TestHelloWorld with Require done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Dio")
	assert.Equal(t, "Hello Dio", result, "Result must be 'Hello Dio'")
	fmt.Println("TestHelloWorld with Assertion done")
}

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

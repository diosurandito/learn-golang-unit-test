package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run() // eksekusi semua unit test

	// after
	fmt.Println("AFTER UNIT TEST")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}
	result := HelloWorld("Dio")
	require.Equal(t, "Hello Dio", result, "Rseult must be 'Hello Dio'")
}

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

package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouteNameResolver(t *testing.T) {
	type testCase struct {
		input  string
		output string
	}

	testcases := []testCase{
		{input: "hello.Hello", output: "/"},
	}

	for _, test := range testcases {
		result := RouteNameResolver(test.input)
		assert.Equal(t, test.output, result)
	}

}

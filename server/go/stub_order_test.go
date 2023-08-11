package server

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSanitySunglasses(t *testing.T) {
	order := OrderWithSingleProduct("whatever")
	prettyPrint(order)
	AssertSanity(t, order)
}

func prettyPrint(i interface{}) {
	data, _ := json.MarshalIndent(i, "", "    ")
	fmt.Println(string(data))
}

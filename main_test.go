package main

import (
	"fmt"
	"testing"
)

func TestPrintSomething(t *testing.T) {
	fmt.Println("Say hi")
	t.Log("Say bye")
}

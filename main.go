package main

import (
	"fmt"

	"github.com/elucid503/Sprout-API-Go/Functions"
)

func main() {

	Error := Functions.Log("ePOSnX6SS26BpLxh", Functions.LogLevelInfo, "Test from Go", "Test log from the new Go module!")

	fmt.Println("Error/Success (Nil): ", Error)

}
// This package is for define App Name, Version App And Database URL.
// Read data from json file
package define

import (
    "encoding/json"
	"fmt"
	"os"
)

// Define struct for data container
type Define struct {
    AppName string
    Version string
    Dburl string
}

// Init func
// Return pointer Define struct
func Init() *Define {
    file, _ := os.Open("define/define.json")
    
	decoder := json.NewDecoder(file)
	define := Define{}
	err := decoder.Decode(&define)
	if err != nil {
		fmt.Println("error Define:", err)
	}

	return &define    
}
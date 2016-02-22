package define

import (
    "encoding/json"
	"fmt"
	"os"
)

type Define struct {
    AppName string
    Version string
    Dburl string
}

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
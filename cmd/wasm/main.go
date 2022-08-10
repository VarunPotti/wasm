package main

import (
	"fmt"
)
import "github.com/spf13/afero"


func main() {
	var AppFs = afero.NewOsFs();
	
	fmt.Println("Go Web Assembly")
	<-make(chan bool)
}

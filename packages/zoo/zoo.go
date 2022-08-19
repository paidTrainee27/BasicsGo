package zoo

import "fmt"

var ZooGlobal string = func() string {
	fmt.Println("zoo global initialisation")
	return "zoo global"
}()

func init() {
	ZooGlobal = "just zoo"
	fmt.Println("zoo init")
}

func Bar() {
	fmt.Println("This function lives in an another file!")
}

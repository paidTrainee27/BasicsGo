package main

import "fmt"


func main()  {
	handleBydefer()
	fmt.Println("Panic handled successfully")
}

func handleBydefer()  {
	defer  func ()  {
		if err := recover(); err != nil {
			fmt.Println("Panic occured",err)
		}
	}()
	dividByZero(54,0)
}

func dividByZero(a,b int)int  {
	return a/b
}

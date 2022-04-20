package main

import (
	"go-basics/logs"
	"time"
)

var users []user

type user struct {
	id   int
	name string
}

func main() {

	logger := logs.NewLogger(logs.Info, time.RFC3339)
	//for basic
	// for i:=0;i<5;i++ { //initialize | condition | increment
	// 	fmt.Println("i is:",i)
	// }

	//basic 2
	// i := 0 //initialize
	// for i<2{ // condition
	// 	fmt.Println(fmt.Sprintf("i is:%d",i))
	// 	i++ //increment
	// }

	//while
	// i := 0 //initialize
	// for {  // condition
	// 	fmt.Println(fmt.Sprintf("i is:%d", i))
	// 	i++ //increment
	// 	if i > 10 {
	// 		break
	// 	}
	// }

	//range
	// x := []int{1,4,6}
	// x1 := map[string]int{"first":1}
	// for i,v := range x1 {
	// 	fmt.Println("index is:",i,"value is:",v)
	// }

	users = append(users, user{1, "ryaaz"})
	users = append(users, user{2, "faraz"})
	users = append(users, user{3, "shayjaz"})

	for _, user := range users {
		// logs.PrintJson(user)
		logger.Log(user.name)
	}

}

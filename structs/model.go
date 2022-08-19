package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

/*Camel casing fields properly requires that the first character be lower-cased.
While JSON doesn’t care how you name your fields, Go does, as it indicates the visibility of the field
outside of the package. Since the encoding/json package is a separate package from the main package we’re using,
we must uppercase the first character in order to make it visible to encoding/json.
It would seem that we’re at an impasse, and we need some way to convey to the JSON encoder what we would like this
field to be named.
*/

type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"-"`//private
	PreferredFish []string  `json:"preferredFish,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
}

func marshalStruct() {
	u := &User{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}
	//won't be able to use json.marshal if fleids of struct are not capitalized
	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}


func marshallingByte(){
	data := []byte(`{"k1":"val1","k2":4,"k3": ["apples","oranges",34]}`)
	var dataMap map[string]interface{}

	if err := json.Unmarshal(data,&dataMap);err != nil {
		fmt.Println(err)
	}

	fmt.Println(dataMap)

	resp ,err := json.Marshal(dataMap)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s",resp)
}
package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	ocCourses := []course{
		{"Reactjs Bootcamp", 299, "OnlineCoder.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 299, "OnlineCoder.in", "abc123", []string{"full-stack", "js"}},
		{"Python Bootcamp", 299, "OnlineCoder.in", "abc123", []string{"cp", "js"}},
		{"Angular Bootcamp", 299, "OnlineCoder.in", "abc123", []string{}},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(ocCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Reactjs Bootcamp",
		"Price": 299,
		"website": "OnlineCoder.in",
		"tags": ["web-dev","js"]
	}
	`)
	var ocCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &ocCourse)
		fmt.Printf("%#v\n", ocCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is : %T\n", k, v, v)
	}
}

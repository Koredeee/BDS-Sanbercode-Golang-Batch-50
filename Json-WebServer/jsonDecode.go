package main

import (
	"encoding/json"
	f "fmt"
)

// untuk membuat variabel baru penampung hasil decode json string.
type User struct {
	// tag json: "Name" digunakan untuk mapping informasi json ke property yg bersangkutan

	FullName string `json:"Name"`
	// kebetulan penulisan Age pada data json dan pada struktur struct adalah "sama", jadinya ga perlu mapping
	Age int
}

type User_2 struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {

	// JSON to struct
	var jsonString = `{"Name": "John Wick", "Age": 20}`

	// func Unmarshal hanaya menerima data json dalam bentuk []byte, so harus di-casting ke []byte dulu
	var jsonData = []byte(jsonString)

	var data User

	// argument ke-2 func unmarshal harus diisi dgn pointer dari object yang bakal nampung hasilnya
	err := json.Unmarshal(jsonData, &data)

	if err != nil {
		f.Println(err.Error())
		return
	}

	f.Println("Ini JSON to Struct")
	f.Println("user : ", data.FullName)
	f.Println("age : ", data.Age)
	f.Println()

	// JSON to map
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	f.Println("ini JSON to map")
	f.Println("user : ", data1["Name"])
	f.Println("Age :", data1["Age"])
	f.Println()

	// JSON to interface{}
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	// harus di-casting dulu
	var decodedData = data2.(map[string]interface{})
	f.Println("ini JSON to interface{}")
	f.Println("user : ", decodedData["Name"])
	f.Println("Age : ", decodedData["Age"])
	f.Println()

	// Array JSON -> Array Object
	jsonString2 := `[
		{"Name": "yusup manjur", "Age": 50},
		{"Name": "soleh solihan", "Age": 30},
		{"Name": "tretan non-muslim", "Age": 28}
	]`

	jsonData2 := []byte(jsonString2)

	var data3 []User_2

	err_2 := json.Unmarshal(jsonData2, &data3)
	if err_2 != nil {
		f.Println(err_2.Error())
		return
	}

	f.Println("ini JSON to Array Object")
	// f.Println("user 1: ", data3[0].FullName)
	// f.Println("user 2: ", data3[1].FullName)
	// f.Println("user 3: ", data3[2].FullName)

	// for i, user := range data3 {
	// 	f.Printf("user %d:  %s\n", i+1, user.FullName)
	// }

	for i, user := range data3 {
		f.Printf("user %d: %s, umur: %d\n", i+1, user.FullName, user.Age)
	}
	f.Println()

	//object to JSON string
	var object = []User{
		{"john doe", 27},
		{"doe john", 32},
	}

	var jsonData3, err_3 = json.Marshal(object)

	if err_3 != nil {
		f.Println(err_3.Error())
	}

	// encode object to JSON string
	var jsonString3 = string(jsonData3)
	f.Println(jsonString3)
	f.Println()

}

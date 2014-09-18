// This source is taken from a tutorial 
// Hope its help for some one

package main

import(
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit Fruits
	Veggi Vegitables
}

type Fruits map[string]int

type Vegitables map[string]int

func serveRest (w http.ResponseWriter, r *http.Request) {

	response, err := getJsonResponse()

	if err != nil {
		panic(err)	
	}

	fmt.Fprintf(w, string(response))
}

func main(){
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:9090", nil)
}

func getJsonResponse()([]byte, error){
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 20

	vegitables := make(map[string]int)
	vegitables["Beatroot"] = 15
	vegitables["Carrot"] = 10

	d := Data {fruits, vegitables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")


}

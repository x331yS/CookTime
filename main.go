package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Recipes struct which contains
// an array of Recipes
type Recipes struct {
	Recipes []Recipe `json:"Recipes"`
}

// Recipe struct which contains a name
// a type and a list of social links
type Recipe struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
	Social Social `json:"social"`
}


// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("Recipes.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Recipes.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Recipes array
	var Recipes Recipes

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'Recipes' which we defined above
	json.Unmarshal(byteValue, &Recipes)

	// we iterate through every Recipe within our Recipes array and
	// print out the Recipe Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(Recipes.Recipes); i++ {
		fmt.Println("Recipe Id: " + strconv.Itoa(Recipes.Recipes[i].ID))
		fmt.Println("Recipe Name: " + Recipes.Recipes[i].Name)
		fmt.Println("Facebook Url: " + Recipes.Recipes[i].Social.Facebook)
		fmt.Println("Image" + Recipes.Recipes[i].Image)
	}

}

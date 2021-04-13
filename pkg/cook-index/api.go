package cook_index

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Recipe []struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Saison         []string `json:"saison"`
	Ingredient     []string `json:"ingredient"`
	Prix           int      `json:"prix"`
	TpsCuisson     int      `json:"tpscuisson"`
	TpsPreparation int      `json:"tpspréparation"`
	Preparation    []string `json:"preparation"`
}

type About struct {
	ID             int
	Image          string
	Suiv           int
	Prec           int
	Aleatoire      int
	Name           string
	Saison         []string
	Ingredient     []string
	Prix           int
	TpsCuisson     int
	TpsPreparation int
	Preparation    []string
}

var Recipes Recipe

var Api = "recipes.json"

func Parse() {
	file, err := ioutil.ReadFile(Api)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	json.Unmarshal(file, &Recipes)

}

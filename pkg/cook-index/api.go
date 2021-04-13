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
	Type           string   `json:"type"`
	Name           string   `json:"name"`
	Saison         []string `json:"saison"`
	Ingredient     []string `json:"ingredient"`
	Prix           string   `json:"prix"`
	TpsCuisson     string   `json:"tpscuisson"`
	TpsPreparation string   `json:"tpspréparation"`
	Preparation    []string `json:"preparation"`
	Link           string   `json:"link"`
	Personne       int      `json:"personne"`
}

type About struct {
	Type           string
	ID             int
	Image          string
	Suiv           int
	Prec           int
	Aleatoire      int
	Name           string
	Saison         []string
	Ingredient     []string
	Prix           string
	TpsCuisson     string
	TpsPreparation string
	Personne       int
	Preparation    []string
	Link           string
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

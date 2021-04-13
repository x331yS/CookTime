package cook_index

import (
	"../cook-error"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Parse()
	if r.URL.Path != "/" {
		//Error 404 IndexHandler
		cook_error.StatusNotFound(w, r)
		return
	}
	t, _ := template.ParseFiles("static/templates/index.html")
	_ = t.Execute(w, Recipes)
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	Parse()
	//Add ID of the artist after the page about
	id, _ := strconv.Atoi(r.URL.Path[len("/about/"):])
	suiv := id + 1
	prec := id - 1
	max := 8
	min := 1
	aleatoire := rand.Intn(max - min) + min
	if suiv == 53 {
		suiv = 1
	}
	if prec == 0 {
		prec = 52
	}
	if id >= 1 && id <= len(Recipes) {
		id = id - 1
		recipe := Recipes[id]
		//Get Data from API
		about := &About{ID: id, Suiv: suiv, Prec: prec, Aleatoire: aleatoire, Image: recipe.Image, Name: recipe.Name, Saison: recipe.Saison, Ingredient: recipe.Ingredient, TpsPreparation: recipe.TpsPreparation, TpsCuisson: recipe.TpsCuisson, Preparation: recipe.Preparation}
		t, _ := template.ParseFiles("./static/templates/about.html")
		_ = t.Execute(w, about)
	} else {
		//Error 400 Handler
		cook_error.StatusNotFound(w, r)
	}
}
func FrIndexHandler(w http.ResponseWriter, r *http.Request) {
	//Handle 404 error
	Parse()
	if r.URL.Path != "/fr/" {
		cook_error.StatusNotFound(w, r)
		return
	}
	//Search Template
	t, err := template.ParseFiles("./static/templates/fr.html")
	if err != nil {
		cook_error.InternalServerError(w, r)
		return
	}
	_ = t.Execute(w, Recipes)
}

func FrAboutHandler(w http.ResponseWriter, r *http.Request) {
	Parse()
	//Add ID of the artist after the page about
	id, _ := strconv.Atoi(r.URL.Path[len("/frabout/"):])
	suiv := id + 1
	prec := id - 1
	max := 8
	min := 1
	aleatoire := rand.Intn(max - min) + min
	if suiv == 53 {
		suiv = 1
	}
	if prec == 0 {
		prec = 52
	}
	if id >= 1 && id <= len(Recipes) {
		id = id - 1
		recipe := Recipes[id]
		//Get Data from API
		about := &About{ID: id, Suiv: suiv, Prec: prec, Aleatoire: aleatoire, Image: recipe.Image, Name: recipe.Name, Saison: recipe.Saison, Ingredient: recipe.Ingredient, TpsPreparation: recipe.TpsPreparation, TpsCuisson: recipe.TpsCuisson, Preparation: recipe.Preparation}
		t, _ := template.ParseFiles("./static/templates/frabout.html")
		_ = t.Execute(w, about)
	} else {
		//Error 400 Handler
		cook_error.StatusNotFound(w, r)
	}
}

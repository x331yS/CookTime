package cook_index

import (
	"../cook-error"
	"net/http"
	"text/template"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	//Handle 404 error
	Parse()
	if r.URL.Path != "/search/" {
		cook_error.StatusNotFound(w, r)
		return
	}
	//Search Template
	t, err := template.ParseFiles("./static/templates/search.html")
	if err != nil {
		cook_error.InternalServerError(w, r)
		return
	}
	_ = t.Execute(w, Recipes)
}

func FrSearchHandler(w http.ResponseWriter, r *http.Request) {
	//Handle 404 error
	Parse()
	if r.URL.Path != "/frsearch/" {
		cook_error.StatusNotFound(w, r)
		return
	}
	//Search Template
	t, err := template.ParseFiles("./static/templates/frsearch.html")
	if err != nil {
		cook_error.InternalServerError(w, r)
		return
	}
	_ = t.Execute(w, Recipes)
}

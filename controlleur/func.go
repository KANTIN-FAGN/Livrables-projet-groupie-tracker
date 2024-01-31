package controlleur

import (
	"net/http"
	"pokemon/temp"
)

func CardPage(w http.ResponseWriter, r *http.Request) {
	temp.Temp.ExecuteTemplate(w, "accueil", r)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	err := temp.Temp.ExecuteTemplate(w, "erreur", nil)
	if err != nil {
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
}
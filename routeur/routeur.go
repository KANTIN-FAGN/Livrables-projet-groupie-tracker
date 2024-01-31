package routeur

import (
	"fmt"
	"log"
	"net/http"
	"pokemon/controlleur"
)

func Initserv() {

	css := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	http.HandleFunc("/accueil", controlleur.CardPage)

	// http.HandleFunc("/", controlleur.DefaultHandler)

	// Démarrage du serveur
	log.Println("[✅] Serveur lancé !")
	fmt.Println("[🌐] http://localhost:8080/accueil")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package routeur

import (
	"fmt"
	"log"
	"net/http"
	"pokemon/backend"
	"pokemon/controlleur"
)

func Initserv() {

	css := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	http.HandleFunc("/accueil", controlleur.IndexPage)
	http.HandleFunc("/login", controlleur.LoginPage)
	http.HandleFunc("/login_credits", controlleur.GetCreds)
	http.HandleFunc("/create-account", controlleur.CreateAccountPage)
	http.HandleFunc("/mail_verif", controlleur.MailVerifPage)
	http.HandleFunc("/verifycode", controlleur.VerifCode)
	http.HandleFunc("/success_code", controlleur.SuccessPage)
	http.HandleFunc("/picture-profil", controlleur.DisplayProfil)
	http.HandleFunc("/treatement/picture/profil", controlleur.SubmitPicture)

	// affichage des pokemon celon leurs types
	http.HandleFunc("/TypeFire", backend.PokemonProfilFire)
	http.HandleFunc("/TypeBug", backend.PokemonProfilBug)
	http.HandleFunc("/TypeDragon", backend.PokemonProfilDragon)
	http.HandleFunc("/TypeElectric", backend.PokemonProfilElectric)
	http.HandleFunc("/TypeFairy", backend.PokemonProfilFairy)
	http.HandleFunc("/TypeFighting", backend.PokemonProfilFighting)
	http.HandleFunc("/TypeGhost", backend.PokemonProfilGhost)
	http.HandleFunc("/TypeGrass", backend.PokemonProfilGrass)
	http.HandleFunc("/TypeGround", backend.PokemonProfilGround)
	http.HandleFunc("/TypeIce", backend.PokemonProfilIce)
	http.HandleFunc("/TypeNormal", backend.PokemonProfilNormal)
	http.HandleFunc("/TypePoison", backend.PokemonProfilPoison)
	http.HandleFunc("/TypePsychic", backend.PokemonProfilPsychic)
	http.HandleFunc("/TypeRock", backend.PokemonProfilRock)
	http.HandleFunc("/TypeWater", backend.PokemonProfilWater)

	http.HandleFunc("/", controlleur.DefaultHandler)

	// Démarrage du serveur
	log.Println("[✅] Serveur lancé !")
	fmt.Println("[🌐] http://localhost:8080/accueil")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

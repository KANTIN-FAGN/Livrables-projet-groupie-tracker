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

	http.HandleFunc("/welcome", controlleur.IndexPage)

	// route relative connexion du compte
	http.HandleFunc("/login", controlleur.LoginPage)

	// route relative à la création du compte
	http.HandleFunc("/login_credits", controlleur.GetCreds)
	http.HandleFunc("/create-account", controlleur.CreateAccountPage)
	http.HandleFunc("/mail_verif", controlleur.MailVerifPage)
	http.HandleFunc("/verifycode", controlleur.VerifCode)
	http.HandleFunc("/success_code", controlleur.SuccessPage)
	http.HandleFunc("/picture-profil", controlleur.DisplayProfil)
	http.HandleFunc("/treatement/picture/profil", controlleur.SubmitPicture)
	http.HandleFunc("/pokecount", controlleur.ProfilPage)
	http.HandleFunc("/deconnexion", controlleur.Deconnexion)

	// route pour la recherche
	http.HandleFunc("/search/cards", backend.SearchPokemonName)
	http.HandleFunc("/cards/pokemon", backend.SearchPokemonAll)
	http.HandleFunc("/cards/pokemon_rarity", backend.SearchPokemonRarity)
	http.HandleFunc("/cards/pokemon_releaseDate", backend.SearchPokemonReleaseDate)
	http.HandleFunc("/cards/trainer", backend.SearchTrainerAll)
	http.HandleFunc("/cards/trainer_rarity", backend.SearchTrainerRarity)
	http.HandleFunc("/cards/trainer_releaseDate", backend.SearchTrainerReleaseDate)
	http.HandleFunc("/cards/energy", backend.SearchEnergyAll)
	http.HandleFunc("/cards/energy_rarity", backend.SearchEnergyRarity)
	http.HandleFunc("/cards/energy_releaseDate", backend.SearchEnergyReleaseDate)

	// Ajout d'une carte en favorie
	http.HandleFunc("/treatement/fav", backend.AddToFav)

	// Filtre pour la recherche des pokemon
	http.HandleFunc("/cards/filtre_rarity", backend.FiltreRarity)
	http.HandleFunc("/cards/filtre_releaseDate", backend.FiltreReleaseDate)

	// route pour display les pokemon
	http.HandleFunc("/card/", backend.DisplayPokemonCards)

	// route de chargement
	http.HandleFunc("/loading", controlleur.LoadingPage)

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

	// Route Error 404
	http.HandleFunc("/", controlleur.DefaultHandler)

	// Démarrage du serveur
	log.Println("[✅] Serveur lancé !")
	fmt.Println("[🌐] http://localhost:8080/welcome")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

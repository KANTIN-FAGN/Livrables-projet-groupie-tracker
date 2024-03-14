package backend

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"pokemon/structure"
)

var ToSend []structure.TypePokemon

func PokemonProfilFire(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgFire int
	nbImgFire = 12
	for i := 1; i <= nbImgFire; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "fire"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilBug(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgbug int
	nbImgbug = 12
	for i := 1; i <= nbImgbug; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "bug"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilWater(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgwater int
	nbImgwater = 30
	for i := 1; i <= nbImgwater; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "water"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilDragon(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgdragon int
	nbImgdragon = 3
	for i := 1; i <= nbImgdragon; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "dragon"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilElectric(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgelectric int
	nbImgelectric = 9
	for i := 1; i <= nbImgelectric; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "electric"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilFairy(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgfairy int
	nbImgfairy = 4
	for i := 1; i <= nbImgfairy; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "fairy"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilFighting(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgfighting int
	nbImgfighting = 7
	for i := 1; i <= nbImgfighting; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "fighting"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilGhost(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgghost int
	nbImgghost = 3
	for i := 1; i <= nbImgghost; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "ghost"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilGrass(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImggrass int
	nbImggrass = 12
	for i := 1; i <= nbImggrass; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "grass"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilGround(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgground int
	nbImgground = 8
	for i := 1; i <= nbImgground; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "ground"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilIce(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgice int
	nbImgice = 1
	for i := 1; i <= nbImgice; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "ice"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilNormal(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgnormal int
	nbImgnormal = 20
	for i := 1; i <= nbImgnormal; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "normal"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilPoison(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgpoison int
	nbImgpoison = 14
	for i := 1; i <= nbImgpoison; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "poison"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilPsychic(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgpsychic int
	nbImgpsychic = 9
	for i := 1; i <= nbImgpsychic; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "psychic"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}
func PokemonProfilRock(w http.ResponseWriter, r *http.Request) {
	ToSend = ToSend[:0]
	var nbImgrock int
	nbImgrock = 7
	for i := 1; i <= nbImgrock; i++ {
		var toGive structure.TypePokemon
		toGive.Type = "rock"
		toGive.Id = i
		ToSend = append(ToSend, toGive)
	}

	fmt.Println(ToSend)
	http.Redirect(w, r, "/picture-profil", http.StatusSeeOther)
}

// GetAccountState récupère le statut d'un utilisateur par son pseudonyme dans le fichier
func GetAccountState(username string) string {
	file, _ := os.ReadFile("./json/accounts.json")

	var accounts Accounts
	json.Unmarshal(file, &accounts)

	for _, account := range accounts.Comptes {
		if account.Username == username {
			return account.State
		}
	}

	return ""
}
// GetAccountMail récupère le mail d'un utilisateur par son pseudonyme dans le fichier
func GetAccountMail(username string) string {
	file, _ := os.ReadFile("./json/accounts.json")

	var accounts Accounts
	json.Unmarshal(file, &accounts)

	for _, account := range accounts.Comptes {
		if account.Username == username {
			return account.Email
		}
	}

	return ""
}
// SetSession paramètre la session utilisateur globale active sur le site
func SetSession(session Session) {
	GlobalSession = session
}
// GetSession renvoie la session utilisateur globale active sur le site
func GetSession() Session {
	return GlobalSession
}

// ClearSession vide la session en cours sur le site
func ClearSession() {
	GlobalSession = Session{}
}
// ClearAccount vide la variable temporaire GlobalAccount pour la création de compte
func ClearAccount() {
	GlobalAccount = AccountCreation{}
}

// AddAccountToFile ajoute une variable temporaire de création de compte à un fichier json
func AddAccountToFile(account AccountCreation, filePath string) error {
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier JSON")
	}

	var data map[string][]map[string]interface{}
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		fmt.Println("Erreur lors du parsing du JSON")
	}

	salt, err := GenerateSalt()
	hashedPassword := HashPassword(account.Password, salt)

	newAccount := map[string]interface{}{
		"picture":  account.Picture,
		"username": account.Username,
		"email":    account.Email,
		"password": hashedPassword,
		"state":    "member",
		"salt":     salt,
		"favCard":  account.FavCards,
	}

	accounts, ok := data["comptes"]
	if !ok {
		accounts = make([]map[string]interface{}, 0)
	}
	data["comptes"] = append(accounts, newAccount)

	newJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("Erreur lors de la conversion en JSON : %v", err)
	}

	err = os.WriteFile(filePath, newJSON, 0644)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'écriture dans le fichier JSON : %v", err)
	}

	return nil
}

// GetUsernameByEmail permet de récupérer le nom d'utilisateur d'un compte par son mail dans la base de données
func GetUsernameByEmail(emailToFind string) string {
	jsonFile, _ := os.ReadFile("./json/accounts.json")

	var data Accounts
	json.Unmarshal(jsonFile, &data)

	for _, account := range data.Comptes {
		if account.Email == emailToFind {
			return account.Username
		}
	}

	return ""
}
// GetEmailsFromJSON permet de récupérer la liste des emails présents dans le fichier json
func GetEmailsFromJSON(filePath string) []string {
	fileContent, _ := os.ReadFile(filePath)

	var comptes Accounts

	json.Unmarshal(fileContent, &comptes)

	var emails []string
	for _, compte := range comptes.Comptes {
		emails = append(emails, compte.Email)
	}

	return emails
}
// GetUsersFromJSON permet de récupérer la liste des noms d'utilisateurs présents dans le fichier json
func GetUsersFromJSON(filePath string) []string {
	fileContent, _ := os.ReadFile(filePath)

	var comptes Accounts

	json.Unmarshal(fileContent, &comptes)

	var usernames []string
	for _, compte := range comptes.Comptes {
		usernames = append(usernames, compte.Username)
	}

	return usernames
}

// GenerateSalt permet de générer un sel afin de hasher un mot de passe
func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(salt), nil
}
// HashPassword permet de hasher un mot de passe avec un sel prédéfini
func HashPassword(password string, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password + salt))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

// CheckRememberStatus permet de lire un fichier json afin de savoir si une session a été sauvegardé par l'utilisateur
func CheckRememberStatus(filename string) (bool, string) {
	content, _ := os.ReadFile(filename)

	var data RememberData
	json.Unmarshal(content, &data)

	if data.Remember.Active == "True" {
		return true, data.Remember.Username
	}

	return false, ""
}
// SetRememberActive permet d'ajouter une sauvegarde de session avec le nom d'utilisateur de la session active
func SetRememberActive(username string, filename string) error {
	content, _ := os.ReadFile(filename)

	var data RememberData
	json.Unmarshal(content, &data)

	data.Remember.Active = "True"
	data.Remember.Username = username

	newContent, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de la création du nouveau contenu JSON")
	}

	os.WriteFile(filename, newContent, 0644)
	return nil
}
// ClearRemember permet de supprimer la sauvegarde de session active
func ClearRemember(filename string) error {
	content, _ := os.ReadFile(filename)

	var data RememberData
	json.Unmarshal(content, &data)

	data.Remember.Active = "False"
	data.Remember.Username = ""

	newContent, _ := json.MarshalIndent(data, "", "  ")

	os.WriteFile(filename, newContent, 0644)

	return nil
}
func IsAdmin() bool {
	return GlobalSession.State == "member"
}

func AddToFav(w http.ResponseWriter, r *http.Request) {
	// Récupérez les informations du formulaire
	pokemonID := r.FormValue("pokemonID")
	pokemonName := r.FormValue("pokemonName")
	pokemonImage := r.FormValue("pokemonImage")

	// Obtenez l'utilisateur de la session
	user := GetSession()

	// Chargez la liste des utilisateurs depuis le fichier JSON
	var accounts struct {
		Comptes []AccountCreation `json:"comptes"`
	}
	dataFile, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier JSON :", err.Error())
		return
	}

	err = json.Unmarshal(dataFile, &accounts)
	if err != nil {
		fmt.Println("Erreur lors du parsing du JSON : ", err.Error())
		return
	}

	// Trouvez le bon utilisateur dans la liste (son index)
	var userIndex int = -1
	for i, u := range accounts.Comptes {
		if u.Username == user.Username {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		fmt.Println("Utilisateur non trouvé dans la liste.")
		// Gérez l'erreur ici (par exemple, renvoyez une réponse d'erreur à l'utilisateur)
		http.Redirect(w, r, "/error-page", http.StatusSeeOther)
		return
	}

	// Ajoutez la nouvelle carte Pokemon à la liste des favoris de l'utilisateur
	accounts.Comptes[userIndex].FavCards = append(accounts.Comptes[userIndex].FavCards, FavoriteCard{pokemonID, pokemonName, pokemonImage})

	// Encodez les données mises à jour au format JSON
	newData, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON :")
		// Gérez l'erreur ici (par exemple, renvoyez une réponse d'erreur à l'utilisateur)
		return
	}

	// Écrivez les données mises à jour dans le fichier JSON
	err = os.WriteFile("./json/accounts.json", newData, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier JSON :")
		// Gérez l'erreur ici (par exemple, renvoyez une réponse d'erreur à l'utilisateur)
		return
	}

	// Redirigez l'utilisateur vers son profil
	http.Redirect(w, r, "/pokecount", http.StatusSeeOther)
}

func TreatementDeleteCard(username, cardID string) error {
	filepath := "./json/accounts.json"

	// Charger les données depuis le fichier JSON
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	var accounts Accounts
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return err
	}

	// Trouver l'utilisateur dans la liste
	for i, acc := range accounts.Comptes {
		if acc.Username == username {
			// Trouver la carte dans les favoris de l'utilisateur
			for j, favCard := range acc.FavCard {
				if favCard.ID == cardID {
					// Supprimer la carte
					accounts.Comptes[i].FavCard = append(acc.FavCard[:j], acc.FavCard[j+1:]...)
					// Encoder les données au format JSON
					newData, err := json.MarshalIndent(accounts, "", "  ")
					if err != nil {
						return err
					}
					// Écrire les données mises à jour dans le fichier JSON
					err = os.WriteFile(filepath, newData, 0644)
					if err != nil {
						return err
					}
					return nil // Carte supprimée avec succès
				}
			}
		}
	}

	return fmt.Errorf("Carte non trouvée pour l'utilisateur %s avec l'ID %s", username, cardID)
}
func DeleteCardHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer les paramètres nécessaires (par exemple, l'ID de la carte et le nom d'utilisateur)
	cardID := r.FormValue("cardID")
	username := r.FormValue("username")

	// Supprimer la carte
	err := TreatementDeleteCard(username, cardID)
	if err != nil {
		// Gérer l'erreur (par exemple, renvoyer une réponse d'erreur à l'utilisateur)
		http.Redirect(w, r, "/error-page", http.StatusSeeOther)
		return
	}

	// Rediriger l'utilisateur vers "/pokecount"
	http.Redirect(w, r, "/pokecount", http.StatusSeeOther)
}
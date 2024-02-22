package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"pokemon/temp"
	"strings"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

var CurrentPage = 1 //Variable globale pour définir à quelle page on se trouve

var CurrentPagePokeAll = 1         // Variable pour la categorie pokemon all
var CurrentPagePokeRarity = 1      // Variable pour la categorie pokemon rarity
var CurrentPagePokeReleaseDate = 1 // Variable pour la categorie pokemon Release Date

var CurrentPageTrainAll = 1         // Variable pour la categorie trainer all
var CurrentPageTrainRarity = 1      // Variable pour la categorie trainer rarity
var CurrentPageTrainReleaseDate = 1 // Variable pour la categorie trainer Release Date

var CurrentPageEnerAll = 1         // Variable pour la categorie energie all
var CurrentPageEnerRarity = 1      // Variable pour la categorie energie rarity
var CurrentPageEnerReleaseDate = 1 // Variable pour la categorie energie Release Date

var CurrentPageSet = 1 // Variable pour les Sets
var CurrentPageSetsRarity = 1 // Variable pour les Sets par rareté
var CurrentPageSetsReleaseDate = 1 // Variable pour les Sets par date de vente

type TempTestResult struct {
	Cards       []*tcg.PokemonCard
	Sets        []*tcg.Set
	DataSearch  string
	DataCompte  IndexData
	CurrentPage int
}

func DisplayPokemonCards(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	PokemonID := strings.TrimPrefix(r.URL.Path, "/card/")

	var lstsearch []*tcg.PokemonCard
	cards, err := c.GetCardByID(PokemonID)
	if err != nil {
		log.Fatal(err)
	}
	lstsearch = append(lstsearch, cards)
	data := TempTestResult{Cards: lstsearch, DataSearch: PokemonID, DataCompte: datacompte}

	temp.Temp.ExecuteTemplate(w, "cards", data)
}

func FiltreRarity(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {
		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPage++
	} else if move == "Moins" {
		CurrentPage--
	}

	fmt.Println("currentPage")

	// Récupérer le nom du Pokémon à partir des paramètres de la requête
	PokemonName := r.URL.Query().Get("q")

	fmt.Println(PokemonName)

	cards, err := c.GetCards(
		request.Query("name:"+PokemonName),
		request.OrderBy("rarity"),
		request.PageSize(10),
		request.Page(CurrentPage),
	)
	if err != nil {
		log.Fatal(err)
	}

	data := TempTestResult{Cards: cards, DataSearch: PokemonName, DataCompte: datacompte, CurrentPage: CurrentPage}

	temp.Temp.ExecuteTemplate(w, "result", data)
}
func FiltreReleaseDate(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {
		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPage++
	} else if move == "Moins" {
		CurrentPage--
	}

	fmt.Println("currentPage")

	// Récupérer le nom du Pokémon à partir des paramètres de la requête
	PokemonName := r.URL.Query().Get("q")

	fmt.Println(PokemonName)

	cards, err := c.GetCards(
		request.Query("name:"+PokemonName),
		request.OrderBy("set.releaseDate"),
		request.PageSize(10),
		request.Page(CurrentPage),
	)
	if err != nil {
		log.Fatal(err)
	}

	data := TempTestResult{Cards: cards, DataSearch: PokemonName, DataCompte: datacompte, CurrentPage: CurrentPage}

	temp.Temp.ExecuteTemplate(w, "result", data)
}

func SearchPokemonName(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae") //token
	session := GetSession() != Session{}                       //est connecté ou pas
	isAdmin := IsAdmin()                                       // a les droits ou pas
	user := GetSession()                                       //pour l'img

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {
		if i.Username == user.Username {
			picture = i.Picture //récupère la photo de profil
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPage++
	} else if move == "Moins" {
		CurrentPage--
	}
	// Récupérer le nom du Pokémon à partir des paramètres de la requête
	PokemonName := r.FormValue("name")

	// Effectuer la requête API avec la page actuelle
	cards, err := c.GetCards(
		request.Query("name:"+PokemonName),
		request.OrderBy("name"),
		request.PageSize(10),
		request.Page(CurrentPage),
	)
	fmt.Println(cards)
	if err != nil {
		log.Fatal(err)
	}

	// Afficher les résultats
	data := TempTestResult{Cards: cards, DataSearch: PokemonName, DataCompte: datacompte, CurrentPage: CurrentPage}
	temp.Temp.ExecuteTemplate(w, "result", data)
}

func SearchPokemonAll(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPagePokeAll++
	} else if move == "Moins" {
		CurrentPagePokeAll--
	}

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	fmt.Println(query)

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("name"),
		request.PageSize(10),
		request.Page(CurrentPagePokeAll),
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Pokémon cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte, CurrentPage: CurrentPagePokeAll}

	temp.Temp.ExecuteTemplate(w, "ResultPokemonAll", data)
}
func SearchPokemonRarity(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPagePokeRarity++
	} else if move == "Moins" {
		CurrentPagePokeRarity--
	}

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("rarity"),
		request.PageSize(10),
		request.Page(CurrentPagePokeRarity),
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Pokémon cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte, CurrentPage: CurrentPagePokeRarity}

	temp.Temp.ExecuteTemplate(w, "ResultPokemonRarity", data)
}
func SearchPokemonReleaseDate(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPagePokeReleaseDate++
	} else if move == "Moins" {
		CurrentPagePokeReleaseDate--
	}

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("set.releaseDate"),
		request.PageSize(10),
		request.Page(CurrentPagePokeReleaseDate),
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Pokémon cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte, CurrentPage: CurrentPagePokeReleaseDate}

	temp.Temp.ExecuteTemplate(w, "ResultPokemonRelease", data)
}

func SearchTrainerAll(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPageTrainAll++
	} else if move == "Moins" {
		CurrentPageTrainAll--
	}

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	fmt.Println(query)

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("name"),
		request.PageSize(10),
		request.Page(CurrentPageTrainAll),
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Pokémon cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte, CurrentPage: CurrentPageTrainAll}

	temp.Temp.ExecuteTemplate(w, "ResultTrainerAll", data)
}
func SearchTrainerRarity(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPageTrainRarity++
	} else if move == "Moins" {
		CurrentPageTrainRarity--
	}

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("rarity"),
		request.PageSize(10),
		request.Page(CurrentPageTrainRarity),
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Pokémon cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte, CurrentPage: CurrentPageTrainRarity}

	temp.Temp.ExecuteTemplate(w, "ResultTrainerRarity", data)
}
func SearchTrainerReleaseDate(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPageTrainReleaseDate++
	} else if move == "Moins" {
		CurrentPageTrainReleaseDate--
	}

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("set.releaseDate"),
		request.PageSize(10),
		request.Page(CurrentPageTrainReleaseDate),
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Pokémon cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte, CurrentPage: CurrentPageTrainReleaseDate}

	temp.Temp.ExecuteTemplate(w, "ResultTrainerRelease", data)
}

func SearchEnergyAll(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPageEnerAll++
	} else if move == "Moins" {
		CurrentPageEnerAll--
	}

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	fmt.Println(query)

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("name"),
		request.PageSize(10),
		request.Page(CurrentPageEnerAll),
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Pokémon cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte, CurrentPage: CurrentPageEnerAll}

	temp.Temp.ExecuteTemplate(w, "ResultEnergyAll", data)
}
func SearchEnergyRarity(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPageEnerRarity++
	} else if move == "Moins" {
		CurrentPageEnerRarity--
	}

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("rarity"),
		request.PageSize(10),
		request.Page(CurrentPageEnerRarity),
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Pokémon cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte, CurrentPage: CurrentPageEnerRarity}

	temp.Temp.ExecuteTemplate(w, "ResultEnergyRarity", data)
}
func SearchEnergyReleaseDate(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPageEnerReleaseDate++
	} else if move == "Moins" {
		CurrentPageEnerReleaseDate--
	}

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("set.releaseDate"),
		request.PageSize(10),
		request.Page(CurrentPageEnerReleaseDate),
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Pokémon cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte, CurrentPage: CurrentPageEnerReleaseDate}

	temp.Temp.ExecuteTemplate(w, "ResultEnergyRelease", data)
}

func SetsPokemon(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPageSet++
	} else if move == "Moins" {
		CurrentPageSet--
	}

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	sets, err := c.GetSets(
		request.PageSize(10),
		request.Page(CurrentPageSet),
	)
	if err != nil {
		log.Fatal(err)
	}

	data := TempTestResult{Sets: sets, DataSearch: query, DataCompte: datacompte, CurrentPage: CurrentPageSet}

	temp.Temp.ExecuteTemplate(w, "ResultSets", data)
}
func SetsPokemonRarity(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPageSet++
	} else if move == "Moins" {
		CurrentPageSet--
	}

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	sets, err := c.GetSets(
		request.OrderBy("rarity"),
		request.PageSize(10),
		request.Page(CurrentPageSetsRarity),
	)
	if err != nil {
		log.Fatal(err)
	}

	data := TempTestResult{Sets: sets, DataSearch: query, DataCompte: datacompte, CurrentPage: CurrentPageSetsRarity}

	temp.Temp.ExecuteTemplate(w, "ResultSetsRarity", data)
}
func SetsPokemonReleaseDate(w http.ResponseWriter, r *http.Request) {
	c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")
	session := GetSession() != Session{}
	isAdmin := IsAdmin()
	user := GetSession()

	fmt.Println(user.Username)

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	datacompte := IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	// Récupérer la valeur de la page à partir des paramètres de la requête
	CurrentPageStr := r.URL.Query().Get("page")
	fmt.Println("dede", CurrentPageStr)

	move := r.URL.Query().Get("move")

	if move == "Plus" {
		CurrentPageSet++
	} else if move == "Moins" {
		CurrentPageSet--
	}

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	sets, err := c.GetSets(
		request.OrderBy("releaseDate"),
		request.PageSize(10),
		request.Page(CurrentPageSetsReleaseDate),
	)
	if err != nil {
		log.Fatal(err)
	}

	data := TempTestResult{Sets: sets, DataSearch: query, DataCompte: datacompte, CurrentPage: CurrentPageSetsReleaseDate}

	temp.Temp.ExecuteTemplate(w, "ResultSetsRelease", data)
}
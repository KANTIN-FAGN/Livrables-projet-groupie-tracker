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

type Request struct {
	endpoint string
	options  map[string]string
}

type TempTestResult struct {
	Cards      []*tcg.PokemonCard
	DataSearch string
	DataCompte IndexData
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

	// Récupérer le nom du Pokémon à partir des paramètres de la requête
	PokemonName := r.URL.Query().Get("q")

	fmt.Println(PokemonName)
	cards, err := c.GetCards(
		request.Query("name:"+PokemonName),
		request.OrderBy("rarity"),
		request.PageSize(10),
	)
	if err != nil {
		log.Fatal(err)
	}

	data := TempTestResult{Cards: cards, DataSearch: PokemonName, DataCompte: datacompte}

	temp.Temp.ExecuteTemplate(w, "result", data)
}
func FiltreReleaseDate(w http.ResponseWriter, r *http.Request) {
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

	// Récupérer le nom du Pokémon à partir des paramètres de la requête
	PokemonName := r.URL.Query().Get("q")

	fmt.Println(PokemonName)
	cards, err := c.GetCards(
		request.Query("name:"+PokemonName),
		request.OrderBy("set.releaseDate"),
		request.PageSize(10),
	)
	if err != nil {
		log.Fatal(err)
	}

	data := TempTestResult{Cards: cards, DataSearch: PokemonName, DataCompte: datacompte}

	temp.Temp.ExecuteTemplate(w, "result", data)
}

func SearchPokemonName(w http.ResponseWriter, r *http.Request) {
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

	// Récupérer le nom du Pokémon à partir des paramètres de la requête
	PokemonName := r.FormValue("name")
	fmt.Println(PokemonName)

	cards, err := c.GetCards(
		request.Query("name:"+PokemonName),
		request.PageSize(10),
		request.Page(2),
	)
	if err != nil {
		log.Fatal(err)
	}

	data := TempTestResult{Cards: cards, DataSearch: PokemonName, DataCompte: datacompte}
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

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("name"),
		request.PageSize(10), // Adjust the order by parameter as needed
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Pokémon cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte}

	temp.Temp.ExecuteTemplate(w, "resultType", data)
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

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("rarity"),
		request.PageSize(10), // Adjust the order by parameter as needed
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Pokemon cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte}

	temp.Temp.ExecuteTemplate(w, "resultType", data)
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

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("set.releaseDate"),
		request.PageSize(10), // Adjust the order by parameter as needed
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Pokemon cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte}

	temp.Temp.ExecuteTemplate(w, "resultType", data)
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

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("name"),
		request.PageSize(10), // Adjust the order by parameter as needed
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Trainer cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte}

	temp.Temp.ExecuteTemplate(w, "resultType", data)
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

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("rarity"),
		request.PageSize(10), // Adjust the order by parameter as needed
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Trainer cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte}

	temp.Temp.ExecuteTemplate(w, "resultType", data)
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

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("set.releaseDate"),
		request.PageSize(10), // Adjust the order by parameter as needed
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Trainer cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte}

	temp.Temp.ExecuteTemplate(w, "resultType", data)
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

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("name"),
		request.PageSize(10), // Adjust the order by parameter as needed
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Energy cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte}

	temp.Temp.ExecuteTemplate(w, "resultType", data)
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

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("rarity"),
		request.PageSize(10), // Adjust the order by parameter as needed
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Trainer cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte}

	temp.Temp.ExecuteTemplate(w, "resultType", data)
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

	// Retrieve the query parameter from the URL
	query := r.URL.Query().Get("q")

	// Modify the query to include the order by parameter
	cards, err := c.GetCards(
		request.Query(query),
		request.OrderBy("set.releaseDate"),
		request.PageSize(10), // Adjust the order by parameter as needed
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching Trainer cards", http.StatusInternalServerError)
		return
	}

	data := TempTestResult{Cards: cards, DataSearch: query, DataCompte: datacompte}

	temp.Temp.ExecuteTemplate(w, "resultType", data)
}

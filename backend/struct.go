package backend

var GlobalSession Session
var GlobalAccount AccountCreation

// Account est une structure qui stock toutes les données d'un compte
type Account struct {
	Picture  string `json:"picture"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	State    string `json:"state"`
	Salt     string `json:"salt"`
}

// Accounts est une structure qui stock une liste d'Account
type Accounts struct {
	Comptes []Account `json:"comptes"`
}

// Session est une structure qui stock les données d'une session
type Session struct {
	Username string
	State    string
	Mail     string
}

// IndexData est une structure qui gère les données envoyées à la page index
type IndexData struct {
	Picture    string
	Username   string
	Email      string
	IsLoggedIn bool
	AsAdmin    bool
}


// MailCodeData est une structure qui gère les données envoyées à la page de vérification mail
type MailCodeData struct {
	Success bool
	Message string
}

// ArticleData est une structure qui gère les données envoyées à la page index
type ArticleData struct {
	IsLoggedIn bool
	AsAdmin    bool
	Data       map[string]interface{}
}

// LoginStatus est une structure qui gère l'état de la connexion active sur le site
type LoginStatus struct {
	IsLoggedIn bool
	AsAdmin    bool
}

type FavoriteCard struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

// AccountCreation est une structure qui stock des données temporaires liées à la création d'un compte
type AccountCreation struct {
	Picture  string         `json:"picture"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	State    string         `json:"state"`
	Salt     string         `json:"salt"`
	FavCards []FavoriteCard `json:"favCard"`
	MailCode string         `json:"-"`
}

// AccountsCreation est une structure qui stock une liste d'AccountCreation
type AccountsCreation struct {
	Comptes []AccountCreation `json:"comptes"`
}

// RememberData est une structure qui gère l'état de la sauvegarde de session
type RememberData struct {
	Remember struct {
		Active   string `json:"Active"`
		Username string `json:"Username"`
	} `json:"remember"`
}

type PokemonInfo struct {
	PokemonID    string `json:"pokemonID"`
	PokemonName  string `json:"pokemonName"`
	PokemonImage string `json:"pokemonImage"`
}
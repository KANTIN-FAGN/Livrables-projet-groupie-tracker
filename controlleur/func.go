package controlleur

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"pokemon/backend"
	"pokemon/structure"
	"pokemon/temp"
	"strconv"
	"time"
)

var SelectedPokemon string

// IndexPage est la fonction handel de la page d'accueil
func IndexPage(w http.ResponseWriter, r *http.Request) {
	session := backend.GetSession() != backend.Session{}
	isAdmin := backend.IsAdmin()
	user := backend.GetSession()

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode backend.Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	data := backend.IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	temp.Temp.ExecuteTemplate(w, "accueil", data)
}
func AboutPage(w http.ResponseWriter, r *http.Request) {
	session := backend.GetSession() != backend.Session{}
	isAdmin := backend.IsAdmin()
	user := backend.GetSession()

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode backend.Accounts

	json.Unmarshal(filedata, &DataDecode)

	var picture string

	for _, i := range DataDecode.Comptes {

		if i.Username == user.Username {
			picture = i.Picture
		}
	}

	data := backend.IndexData{
		Picture:    picture,
		IsLoggedIn: session,
		AsAdmin:    isAdmin,
	}

	temp.Temp.ExecuteTemplate(w, "about", data)
}
func ProfilPage(w http.ResponseWriter, r *http.Request) {
	user := backend.GetSession()

	filedata, err := os.ReadFile("./json/accounts.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}

	var DataDecode backend.Accounts

	err = json.Unmarshal(filedata, &DataDecode)
	if err != nil {
		fmt.Println("Erreur lors du parsing du JSON :", err.Error())
		return
	}

	var picture string
	var username string
	var email string
	var favCards []backend.FavoriteCard

	for _, i := range DataDecode.Comptes {
		if i.Username == user.Username {
			picture = i.Picture
			username = i.Username
			email = i.Email
			favCards = i.FavCard
		}
	}

	data := backend.IndexData{
		Picture:    picture,
		Username:   username,
		Email:      email,
		FavCards:   favCards,
	}

	temp.Temp.ExecuteTemplate(w, "profil", data)
}
func DisplayProfil(w http.ResponseWriter, r *http.Request) {
	if backend.ToSend == nil { //Tosend type fire de base
		var nbImgFire int
		nbImgFire = 12
		for i := 1; i <= nbImgFire; i++ {
			var toGive structure.TypePokemon
			toGive.Type = "fire"
			toGive.Id = i
			backend.ToSend = append(backend.ToSend, toGive)
		}
	}
	temp.Temp.ExecuteTemplate(w, "picture-profil", backend.ToSend)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	err := temp.Temp.ExecuteTemplate(w, "erreur", nil)
	if err != nil {
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
}

// LoginPage est la fonction qui permet d'exécuter la page de login
func LoginPage(w http.ResponseWriter, r *http.Request) {
	errMessage := r.URL.Query().Get("error")
	temp.Temp.ExecuteTemplate(w, "login", errMessage)
}

func CreateAccountPage(w http.ResponseWriter, r *http.Request) {
	errMessage := r.URL.Query().Get("error")
	temp.Temp.ExecuteTemplate(w, "create-profil", errMessage)
}

// GetCreds est la fonction qui permet de récupérer les données de login et de les traiter
func GetCreds(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données du formulaire :", err)
		return
	}

	mail := r.Form.Get("email")
	password := r.Form.Get("password")
	remember := r.FormValue("remember")

	file, _ := os.ReadFile("./json/accounts.json")

	var accounts backend.Accounts
	json.Unmarshal(file, &accounts)

	valid := false
	for _, account := range accounts.Comptes {
		if account.Email == mail {
			if backend.HashPassword(password, account.Salt) == account.Password {
				valid = true
				break
			}
		}
	}

	if valid {
		username := backend.GetUsernameByEmail(mail)
		session := backend.Session{Username: username, State: backend.GetAccountState(username), Mail: mail}
		if remember == "on" {
			backend.SetRememberActive(username, "./json/rememberSession.json")
		}
		backend.SetSession(session)
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	} else {
		fmt.Println("Identifiants incorrects")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// MailVerifPage permet de vérifier les champs rentrés lors de la création d'un compte et envoie un mail pour valider le compte
func MailVerifPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		data := backend.MailCodeData{
			Success: false,
			Message: "Le code est incorrect !",
		}
		temp.Temp.ExecuteTemplate(w, "mailverif", data)
		return
	}

	r.ParseForm()

	emailDestinataire := r.FormValue("email")
	username := r.FormValue("username")
	passwordAccount := r.FormValue("password")

	for _, element := range backend.GetEmailsFromJSON("./json/accounts.json") {
		if element == emailDestinataire {
			fmt.Println("Mail déjà utilisé")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	}

	if len(username) < 5 {
		fmt.Println("Nom d'utilisateur trop court")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	} else {
		for _, element := range backend.GetUsersFromJSON("./json/accounts.json") {
			if element == username {
				fmt.Println("Nom d'utilisateur déjà utilisé")
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}
	}

	if len(passwordAccount) < 8 {
		fmt.Println("Mot de passe trop court")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	rand.Seed(time.Now().UnixNano())

	codeMail := rand.Intn(899999) + 100000
	codeMailString := strconv.Itoa(codeMail)

	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	email := "pokecenter.verified@gmail.com"
	password := "rkjq rwfu mhby zwyh"

	to := []string{emailDestinataire}
	subject := "PokeCenter verification code"
	body := "Hello " + username + ",\nHere is your verification code : " + codeMailString

	auth := smtp.PlainAuth("", email, password, smtpHost)

	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, email, to, msg)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("E-mail envoyé avec succès!")

	tempaccount := backend.AccountCreation{Picture: SelectedPokemon, Username: username, Email: emailDestinataire, Password: passwordAccount, MailCode: codeMailString}
	backend.GlobalAccount = tempaccount

	data := backend.MailCodeData{
		Success: true,
		Message: "",
	}

	temp.Temp.ExecuteTemplate(w, "mailverif", data)
}

// VerifCode vérifie si le code reçu et le code envoyé par mail correspondent et redirige en fonction
func VerifCode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	codeRecu := r.FormValue("verificationCode")
	codeEnvoye := backend.GlobalAccount.MailCode

	if codeRecu == codeEnvoye {
		http.Redirect(w, r, "/success_code", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/mail_verif", http.StatusSeeOther)
}

// SuccessPage est la fonction handler de la page de succès après la création d'un compte
func SuccessPage(w http.ResponseWriter, r *http.Request) {
	backend.AddAccountToFile(backend.GlobalAccount, "./json/accounts.json")
	backend.ClearAccount()
	temp.Temp.ExecuteTemplate(w, "success", nil)
}

func SubmitPicture(w http.ResponseWriter, r *http.Request) {
	SelectedPokemon = r.FormValue("picture")
	http.Redirect(w, r, "/create-account", http.StatusMovedPermanently)
}

func Deconnexion(w http.ResponseWriter, r *http.Request) {
	backend.ClearSession()
	backend.ClearRemember("./json/rememberSession.json")
	http.Redirect(w, r, "/welcome", http.StatusSeeOther)
}

func LoadingPage(w http.ResponseWriter, r *http.Request) {
	errMessage := r.URL.Query().Get("error")
	temp.Temp.ExecuteTemplate(w, "loading", errMessage)
}

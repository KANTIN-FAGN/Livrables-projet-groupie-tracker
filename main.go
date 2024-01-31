package main

import (
	"fmt"
	"pokemon/backend"
	"pokemon/routeur"
	"pokemon/temp"
)

func main() {
	active, user := backend.CheckRememberStatus("rememberSession.json")
	if active {
		fmt.Println("Une session a été sauvegardée")
		backend.GlobalSession = backend.Session{Username: user, State: backend.GetAccountState(user), Mail: backend.GetAccountMail(user)}
		fmt.Println("Session initialisée")
	}
	temp.InitTemplate()
	routeur.Initserv()
}

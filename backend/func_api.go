package backend

import (
	"log"
	"net/http"
	"pokemon/temp"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

type ResultData struct {
	Cards []tcg.PokemonCard
}

func SearchPokemon(w http.ResponseWriter, r *http.Request) {
    c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")

    // Récupérer le nom du Pokémon à partir des paramètres de la requête
    PokemonName := r.FormValue("name")

    cards, err := c.GetCards(
        request.Query("name:" + PokemonName),
        request.OrderBy("+name"),
    )
    if err != nil {
        log.Fatal(err)
    }

    temp.Temp.ExecuteTemplate(w, "result", cards)
}

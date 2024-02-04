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

func SearchPokemonName(w http.ResponseWriter, r *http.Request) {
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

func SearchPokemonAll(w http.ResponseWriter, r *http.Request) {
    c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")

    // Retrieve the query parameter from the URL
    query := r.URL.Query().Get("q")

    // Modify the query to include the order by parameter
    cards, err := c.GetCards(
        request.Query(query),
        request.OrderBy("name"), // Adjust the order by parameter as needed
    )
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Error fetching Pokémon cards", http.StatusInternalServerError)
        return
    }

    temp.Temp.ExecuteTemplate(w, "result", cards)
}
func SearchPokemonRarity(w http.ResponseWriter, r *http.Request) {
    c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")

    // Retrieve the query parameter from the URL
    query := r.URL.Query().Get("q")

    // Modify the query to include the order by parameter
    cards, err := c.GetCards(
        request.Query(query),
        request.OrderBy("rarity"), // Adjust the order by parameter as needed
    )
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Error fetching Pokemon cards", http.StatusInternalServerError)
        return
    }

    temp.Temp.ExecuteTemplate(w, "result", cards)
}
func SearchPokemonReleaseDate(w http.ResponseWriter, r *http.Request) {
    c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")

    // Retrieve the query parameter from the URL
    query := r.URL.Query().Get("q")

    // Modify the query to include the order by parameter
    cards, err := c.GetCards(
        request.Query(query),
        request.OrderBy("set.releaseDate"), // Adjust the order by parameter as needed
    )
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Error fetching Pokemon cards", http.StatusInternalServerError)
        return
    }

    temp.Temp.ExecuteTemplate(w, "result", cards)
}

func SearchTrainerAll(w http.ResponseWriter, r *http.Request) {
    c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")

    // Retrieve the query parameter from the URL
    query := r.URL.Query().Get("q")

    // Modify the query to include the order by parameter
    cards, err := c.GetCards(
        request.Query(query),
        request.OrderBy("name"), // Adjust the order by parameter as needed
    )
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Error fetching Trainer cards", http.StatusInternalServerError)
        return
    }

    temp.Temp.ExecuteTemplate(w, "result", cards)
}
func SearchTrainerRarity(w http.ResponseWriter, r *http.Request) {
    c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")

    // Retrieve the query parameter from the URL
    query := r.URL.Query().Get("q")

    // Modify the query to include the order by parameter
    cards, err := c.GetCards(
        request.Query(query),
        request.OrderBy("rarity"), // Adjust the order by parameter as needed
    )
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Error fetching Trainer cards", http.StatusInternalServerError)
        return
    }

    temp.Temp.ExecuteTemplate(w, "result", cards)
}
func SearchTrainerReleaseDate(w http.ResponseWriter, r *http.Request) {
    c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")

    // Retrieve the query parameter from the URL
    query := r.URL.Query().Get("q")

    // Modify the query to include the order by parameter
    cards, err := c.GetCards(
        request.Query(query),
        request.OrderBy("set.releaseDate"), // Adjust the order by parameter as needed
    )
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Error fetching Trainer cards", http.StatusInternalServerError)
        return
    }

    temp.Temp.ExecuteTemplate(w, "result", cards)
}

func SearchEnergyAll(w http.ResponseWriter, r *http.Request) {
    c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")

    // Retrieve the query parameter from the URL
    query := r.URL.Query().Get("q")

    // Modify the query to include the order by parameter
    cards, err := c.GetCards(
        request.Query(query),
        request.OrderBy("name"), // Adjust the order by parameter as needed
    )
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Error fetching Energy cards", http.StatusInternalServerError)
        return
    }

    temp.Temp.ExecuteTemplate(w, "result", cards)
}
func SearchEnergyRarity(w http.ResponseWriter, r *http.Request) {
    c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")

    // Retrieve the query parameter from the URL
    query := r.URL.Query().Get("q")

    // Modify the query to include the order by parameter
    cards, err := c.GetCards(
        request.Query(query),
        request.OrderBy("rarity"), // Adjust the order by parameter as needed
    )
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Error fetching Trainer cards", http.StatusInternalServerError)
        return
    }

    temp.Temp.ExecuteTemplate(w, "result", cards)
}
func SearchEnergyReleaseDate(w http.ResponseWriter, r *http.Request) {
    c := tcg.NewClient("f8165ff9-ad83-41ea-ba42-6fb0cc2835ae")

    // Retrieve the query parameter from the URL
    query := r.URL.Query().Get("q")

    // Modify the query to include the order by parameter
    cards, err := c.GetCards(
        request.Query(query),
        request.OrderBy("set.releaseDate"), // Adjust the order by parameter as needed
    )
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Error fetching Trainer cards", http.StatusInternalServerError)
        return
    }

    temp.Temp.ExecuteTemplate(w, "result", cards)
}
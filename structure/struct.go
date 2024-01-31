package structure

type PokemonCard struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Supertype    string   `json:"supertype"`
	Subtypes     []string `json:"subtypes"`
	Level        string   `json:"level"`
	Hp           string   `json:"hp"`
	Types        []string `json:"types"`
	EvolvesFrom  string   `json:"evolvesFrom"`
	EvolvesTo    []string `json:"evolvesTo"`
	Rules        []string `json:"rules"`
	AncientTrait *struct {
		Name string `json:"name"`
		Text string `json:"text"`
	} `json:"ancientTrait"`
	Abilities []struct {
		Name string `json:"name"`
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"abilities"`
	Attacks []struct {
		Name                string   `json:"name"`
		Cost                []string `json:"cost"`
		ConvertedEnergyCost int      `json:"convertedEnergyCost"`
		Damage              string   `json:"damage"`
		Text                string   `json:"text"`
	} `json:"attacks"`
	Weaknesses []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"weaknesses"`
	Resistances []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"resistances"`
	RetreatCost          []string `json:"retreatCost"`
	ConvertedRetreatCost int      `json:"convertedRetreatCost"`
	Set                  struct {
		Name         string `json:"name"`
		Images      struct {
			Symbol string `json:"symbol"`
			Logo   string `json:"logo"`
		} `json:"images"`
	} `json:"set"`
	Number                 string `json:"number"`
	Artist                 string `json:"artist"`
	Rarity                 string `json:"rarity"`
	FlavorText             string `json:"flavorText"`
	NationalPokedexNumbers []int  `json:"nationalPokedexNumbers"`
	Images struct {
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"images"`
	TCGPlayer struct {
		URL       string `json:"url"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"tcgplayer"`
	CardMarket struct {
		URL       string `json:"url"`
		UpdatedAt string `json:"updatedAt"`
		Prices    struct {
			AverageSellPrice *float64 `json:"averageSellPrice"`
		} `json:"prices"`
	} `json:"cardmarket"`
}

type Set struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Series       string `json:"series"`
	PrintedTotal int    `json:"printedTotal"`
	Total        int    `json:"total"`
	Legalities   struct {
		Unlimited string `json:"unlimited"`
	} `json:"legalities"`
	PtcgoCode   string `json:"ptcgoCode"`
	ReleaseDate string `json:"releaseDate"`
	UpdatedAt   string `json:"updatedAt"`
	Images      struct {
		Symbol string `json:"symbol"`
		Logo   string `json:"logo"`
	} `json:"images"`
}

type apiClient struct {
	apiKey string
}
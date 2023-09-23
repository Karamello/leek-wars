package leekwars

type LeekInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Farmer struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		AvatarChanged int    `json:"avatar_changed"`
	} `json:"farmer"`
	Talent      int    `json:"talent"`
	TalentMore  int    `json:"talent_more"`
	Level       int    `json:"level"`
	Xp          int    `json:"xp"`
	UpXp        int    `json:"up_xp"`
	DownXp      int    `json:"down_xp"`
	RemainingXp int    `json:"remaining_xp"`
	Life        int    `json:"life"`
	Strength    int    `json:"strength"`
	Wisdom      int    `json:"wisdom"`
	Agility     int    `json:"agility"`
	Resistance  int    `json:"resistance"`
	Science     int    `json:"science"`
	Magic       int    `json:"magic"`
	Frequency   int    `json:"frequency"`
	Tp          int    `json:"tp"`
	Mp          int    `json:"mp"`
	Victories   int    `json:"victories"`
	Draws       int    `json:"draws"`
	Defeats     int    `json:"defeats"`
	Ratio       string `json:"ratio"`
	Chips       []struct {
		Template int `json:"template"`
		ID       int `json:"id"`
	} `json:"chips"`
	Weapons []struct {
		Template int `json:"template"`
		ID       int `json:"id"`
	} `json:"weapons"`
	Ai struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Level      int    `json:"level"`
		Valid      bool   `json:"valid"`
		Version    int    `json:"version"`
		TotalLines int    `json:"total_lines"`
	} `json:"ai"`
	MaxWeapons int  `json:"max_weapons"`
	MaxChips   int  `json:"max_chips"`
	InGarden   bool `json:"in_garden"`
	Fights     []struct {
		ID           int    `json:"id"`
		Leeks1       []any  `json:"leeks1"`
		Leeks2       []any  `json:"leeks2"`
		Winner       int    `json:"winner"`
		Status       int    `json:"status"`
		Date         int    `json:"date"`
		Context      int    `json:"context"`
		Type         int    `json:"type"`
		Farmer1      int    `json:"farmer1"`
		Farmer2      int    `json:"farmer2"`
		Team1        int    `json:"team1"`
		Team2        int    `json:"team2"`
		Composition1 int    `json:"composition1"`
		Composition2 int    `json:"composition2"`
		Result       string `json:"result"`
		Farmer1Name  string `json:"farmer1_name,omitempty"`
		Farmer2Name  string `json:"farmer2_name,omitempty"`
	} `json:"fights"`
	Skin int `json:"skin"`
	Hat  struct {
		ID       int `json:"id"`
		Template int `json:"template"`
		Quantity int `json:"quantity"`
	} `json:"hat"`
	TalentHistory []int `json:"talent_history"`
	Tournaments   []struct {
		ID     int `json:"id"`
		Date   int `json:"date"`
		Rounds struct {
			Num1 []int `json:"1"`
		} `json:"rounds"`
	} `json:"tournaments"`
	Weapon    int   `json:"weapon"`
	Title     []int `json:"title"`
	Ranking   int   `json:"ranking"`
	Metal     bool  `json:"metal"`
	Face      int   `json:"face"`
	XpBlocked bool  `json:"xp_blocked"`
}
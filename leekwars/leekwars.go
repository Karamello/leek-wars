package leekwars

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type LeekWars struct {
	baseURL string
	token string
}

type Leek struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Talent  int    `json:"talent"`
	Skin    int    `json:"skin"`
	Hat     int    `json:"hat"`
	Weapon  any    `json:"weapon"`
	Metal   bool   `json:"metal"`
	Face    int    `json:"face"`
	Fights  int    `json:"fights"`
	Country string `json:"country"`
}

type GardenLeeks struct {
	Opponents []Leek `json:"opponents"`
}

func New(baseURL string) *LeekWars {
	username := os.Getenv("LEEK_WARS_USERNAME")
	password := os.Getenv("LEEK_WARS_PASSWORD")

	v := url.Values{}
	v.Add("login", username)
	v.Add("password", password)

	resp, err := http.DefaultClient.PostForm(baseURL +"farmer/login-token", v)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	var data interface{}

	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	return &LeekWars{baseURL: baseURL, token: data.(map[string]interface{})["token"].(string)}
}

func (lw LeekWars) GetFarmerOpponents() {
	endpoint := lw.baseURL + "garden/get-farmer-opponents"

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", lw.token))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	var data interface{}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if jsonerr := json.Unmarshal(body, &data); jsonerr != nil {
		panic(jsonerr)
	}

	fmt.Println(data)
}

func (lw LeekWars) GetLeekOpponents(leekid int) []Leek {
	endpoint := fmt.Sprintf("%s%s/%d", lw.baseURL, "garden/get-leek-opponents", leekid)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", lw.token))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	var data GardenLeeks
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if jsonerr := json.Unmarshal(body, &data); jsonerr != nil {
		panic(jsonerr)
	}

	return data.Opponents
}

func (lw LeekWars) GetLeekInfo(leekid int) LeekInfo {
	endpoint := fmt.Sprintf("%s%s/%d", lw.baseURL, "leek/get", leekid)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", lw.token))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	var data LeekInfo
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if jsonerr := json.Unmarshal(body, &data); jsonerr != nil {
		panic(jsonerr)
	}

	return data
}
package profile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Response struct {
	Status string `json:"status"`
	Data   struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Surname     string `json:"surname"`
		Description string `json:"description"`
		Image       string `json:"image"`
		Linkedin    string `json:"linkedin"`
		Facebook    string `json:"facebook"`
		Twitter     string `json:"twitter"`
		Github      string `json:"github"`
	} `json:"data"`
}

func JsonGet(i interface{}) string {
	resp, err := http.Get("http://127.0.0.1:8000/restapp/profile-api/1?format=json")
	if err != nil {
		fmt.Println("Profile Server İs Not Connect")

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)

}

func ConvertJSON() {
	var result Response
	if err := json.Unmarshal([]byte(JsonGet(result)), &result); err != nil {
		fmt.Println("JSON HATA!")

	}
	ıd, name, surname, description, linkedin, facebook, twitter, github := result.Data.ID, result.Data.Name, result.Data.Surname, result.Data.Description, result.Data.Linkedin, result.Data.Facebook, result.Data.Twitter, result.Data.Github
	fmt.Printf("Kimlik Sırası:%d\n", ıd)
	fmt.Printf("Adı:--%s\n", name)
	fmt.Printf("SoyAdı:--%s\n", surname)
	fmt.Printf("Uzmanlık Alanı:--%s\n", description)
	fmt.Printf("Linkedin:--%s\n", linkedin)
	fmt.Printf("Facebook:--%s\n", facebook)
	fmt.Printf("twitter:--%s\n", twitter)
	fmt.Printf("Github:--%s\n", github)

	f, err := os.Create("CV.txt")
	if err != nil {
		fmt.Println("dosya hatası")
	}
	defer f.Close()

	c := strconv.Itoa(ıd)
	f.WriteString("Kimlik Sırası: " + c)
	f.WriteString("\nAd: " + name)
	f.WriteString("\nSoyAd: " + surname)
	f.WriteString("\nUzmanlık Alanı: " + description)
	f.WriteString("\nLinkedin: " + linkedin)
	f.WriteString("\nFacebook: " + facebook)
	f.WriteString("\nTwitter: " + twitter)
	f.WriteString("\nGithub: " + github)

}

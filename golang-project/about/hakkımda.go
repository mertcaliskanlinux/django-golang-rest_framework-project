package about

import (
	"encoding/json"
	"fmt"
	"go-portofillo/goodat"
	"io/ioutil"
	"net/http"
	"os"
)

type AboutResponse struct {
	Status string `json:"status"`
	Data   []struct {
		TitleL       string `json:"titleL"`
		DescriptionL string `json:"descriptionL"`
		TitleR       string `json:"titleR"`
		DescriptionR string `json:"descriptionR"`
	} `json:"data"`
}

func jsonGet(i interface{}) string {

	resp, err := http.Get("http://127.0.0.1:8000/restapp/profile-api/about?format=json")
	if err != nil {
		fmt.Println("About Server İs Not Connect")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func AboutConvertJson() {
	var About AboutResponse
	f, err := os.OpenFile(goodat.Txtname(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("About Dosya Yazma Hatası")
	}
	defer f.Close()

	if err := json.Unmarshal([]byte(jsonGet(About)), &About); err != nil {
		fmt.Println("ABOUT JSON CONVERT HATA")

	}

	for _, i := range About.Data {
		fmt.Printf("Hakkımda:--%s\n", i.TitleL)
		f.WriteString("\nHakkımda: " + i.TitleL)
		fmt.Printf("Açıklama:-- %s,%s\n", i.DescriptionL, i.DescriptionR)
		f.WriteString("\n" + i.DescriptionL)
		f.WriteString("\n" + i.DescriptionR)
	}

}

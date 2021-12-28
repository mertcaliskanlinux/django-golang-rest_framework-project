package goodat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type GoodAtResponse struct {
	Status string `json:"status"`
	Data   []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"data"`
}

func jsonget(i interface{}) string {
	resp, err := http.Get("http://127.0.0.1:8000/restapp/profile-api/goodat")
	if err != nil {
		fmt.Println("GoodAT Server İs Not Connect")

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func GoodATConvertJson() {
	var GoodAt GoodAtResponse

	f, err := os.OpenFile(Txtname(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("About Yazma Hatası")
	}
	defer f.Close()

	if err := json.Unmarshal([]byte(jsonget(GoodAt)), &GoodAt); err != nil {
		fmt.Println("GoodAT JSON CONVERT HATA")

	}

	for _, i := range GoodAt.Data {
		if i.ID == 1 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("PROJE AÇIKLAMASI:%s\n", i.Description)

			f.WriteString("\nTeknoloji: " + i.Title)
			f.WriteString("\nPROJE AÇIKLAMASI: " + i.Description)

		} else if i.ID == 2 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("PROJE AÇIKLAMASI:%s\n", i.Description)

			f.WriteString("\nTeknoloji: " + i.Title)
			f.WriteString("\nPROJE AÇIKLAMASI: " + i.Description)
		} else if i.ID == 3 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("PROJE AÇIKLAMASI:%s\n", i.Description)

			f.WriteString("\nTeknoloji: " + i.Title)
			f.WriteString("\nPROJE AÇIKLAMASI: " + i.Description)
		} else if i.ID == 4 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("PROJE AÇIKLAMASI:%s\n", i.Description)

			f.WriteString("\nTeknoloji: " + i.Title)
			f.WriteString("\nPROJE AÇIKLAMASI: " + i.Description)
		} else if i.ID == 5 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("PROJE AÇIKLAMASI:%s\n", i.Description)

			f.WriteString("\nTeknoloji: " + i.Title)
			f.WriteString("\nPROJE AÇIKLAMASI: " + i.Description)
		} else if i.ID == 6 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("PROJE AÇIKLAMASI:%s\n", i.Description)

			f.WriteString("\nTeknoloji: " + i.Title)
			f.WriteString("\nPROJE AÇIKLAMASI: " + i.Description)
		} else if i.ID == 7 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("PROJE AÇIKLAMASI:%s\n", i.Description)

			f.WriteString("\nTeknoloji: " + i.Title)
			f.WriteString("\nPROJE AÇIKLAMASI: " + i.Description)
		} else if i.ID == 8 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("PROJE AÇIKLAMASI:%s\n", i.Description)

			f.WriteString("\nTeknoloji: " + i.Title)
			f.WriteString("\nPROJE AÇIKLAMASI: " + i.Description)
		} else if i.ID == 9 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("PROJE AÇIKLAMASI:%s\n", i.Description)

			f.WriteString("\nTeknoloji: " + i.Title)
			f.WriteString("\nPROJE AÇIKLAMASI: " + i.Description)
		} else if i.ID == 10 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("PROJE AÇIKLAMASI:%s\n", i.Description)

			f.WriteString("\nTeknoloji: " + i.Title)
			f.WriteString("\nPROJE AÇIKLAMASI: " + i.Description)
		} else {
			fmt.Println("Maksimum Seviyeye Ulaştınız En Fazla 10 Tane DATA Ekleye Bilirsiniz!!")
		}

	}

}

func Txtname() string {

	return "CV.txt"
}

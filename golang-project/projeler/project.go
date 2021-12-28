package projeler

import (
	"encoding/json"
	"fmt"
	"go-portofillo/goodat"
	"io/ioutil"
	"net/http"
	"os"
)

type ProjectResponse struct {
	Status string `json:"status"`
	Data   []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"data"`
}

func jsonget(i interface{}) string {
	resp, err := http.Get("http://127.0.0.1:8000/restapp/profile-api/projects")
	if err != nil {
		fmt.Println("Profile URL Server İs not Connect")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(body)
}

func ProfileConvertJson() {
	var Project ProjectResponse
	f, err := os.OpenFile(goodat.Txtname(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Project Dosya Yazma Hatası")
	}
	defer f.Close()

	if err := json.Unmarshal([]byte(jsonget(Project)), &Project); err != nil {
		fmt.Println("PROFİLE JSON CONVERT HATA")
	}

	for _, i := range Project.Data {
		if i.ID == 1 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("%s\n", i.Description)
			fmt.Println("******************")

			f.WriteString("\nProje: " + i.Title)
			f.WriteString("\nProje Açıklaması: " + i.Description)
		} else if i.ID == 2 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("%s\n", i.Description)
			fmt.Println("******************")

			f.WriteString("\nProje: " + i.Title)
			f.WriteString("\nProje Açıklaması: " + i.Description)
		} else if i.ID == 3 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("%s\n", i.Description)
			fmt.Println("******************")

			f.WriteString("\nProje: " + i.Title)
			f.WriteString("\nProje Açıklaması: " + i.Description)

		} else if i.ID == 4 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("%s\n", i.Description)
			fmt.Println("******************")

			f.WriteString("\nProje: " + i.Title)
			f.WriteString("\nProje Açıklaması: " + i.Description)

		} else if i.ID == 5 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("%s\n", i.Description)
			fmt.Println("******************")

			f.WriteString("\nProje: " + i.Title)
			f.WriteString("\nProje Açıklaması: " + i.Description)

		} else if i.ID == 6 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("%s\n", i.Description)
			fmt.Println("******************")

			f.WriteString("\nProje: " + i.Title)
			f.WriteString("\nProje Açıklaması: " + i.Description)

		} else if i.ID == 7 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("%s\n", i.Description)
			fmt.Println("******************")

			f.WriteString("\nProje: " + i.Title)
			f.WriteString("\nProje Açıklaması: " + i.Description)

		} else if i.ID == 8 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("%s\n", i.Description)
			fmt.Println("******************")

			f.WriteString("\nProje: " + i.Title)
			f.WriteString("\nProje Açıklaması: " + i.Description)

		} else if i.ID == 9 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("%s\n", i.Description)
			fmt.Println("******************")

			f.WriteString("\nProje: " + i.Title)
			f.WriteString("\nProje Açıklaması: " + i.Description)

		} else if i.ID == 10 {
			fmt.Printf("%s\n", i.Title)
			fmt.Printf("%s\n", i.Description)
			fmt.Println("******************")

			f.WriteString("\nProje: " + i.Title)
			f.WriteString("\nProje Açıklaması: " + i.Description)

		} else {
			fmt.Println("MAKSİMUM SEVİYEYE ULAŞTINIZ MAX=10!!")
		}
	}
}

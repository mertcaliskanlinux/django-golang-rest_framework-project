package main

import (
	"fmt"
	"go-portofillo/about"
	"go-portofillo/goodat"
	"go-portofillo/profile"
	"go-portofillo/projeler"
	"time"
)

func main() {
	go profile.ConvertJSON()
	fmt.Println("******************")
	time.Sleep(time.Second * 2)
	fmt.Println("ABOUT İNDİRİLİYOR...")
	go about.AboutConvertJson()
	fmt.Println("******************")
	time.Sleep(time.Second * 2)

	go goodat.GoodATConvertJson()
	fmt.Println("******************")
	time.Sleep(time.Second * 2)

	go projeler.ProfileConvertJson()
	fmt.Println("******************")
	time.Sleep(time.Second * 2)

	fmt.Println("WEB SİTE OKUMA TAMAMLANIYOR 3!")
	time.Sleep(time.Second * 1)
	fmt.Println("WEB SİTE OKUMA TAMAMLANIYOR 2!")
	time.Sleep(time.Second * 1)
	fmt.Println("WEB SİTE OKUMA TAMAMLANIYOR 1!")
	time.Sleep(time.Second * 1)
	defer fmt.Println("WEB SİTE OKUMA TAMAMLANDI..")
	fmt.Println("TXT DOSYASI OLUŞTURULDU...")

}

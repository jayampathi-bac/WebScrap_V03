package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"log"
)
func main() {
	var location string;
	var category string;

	fmt.Print("Enter Location : ")
	fmt.Scan(&location)
	fmt.Print("Enter Category : ")
	fmt.Scan(&category)

	var URL string = "https://ikman.lk/en/ads/"+ location +"/" +category ;

	findMainURL(URL)
}

func findMainURL(URL string)  {
	c1 := colly.NewCollector()
	//////////////////////////////////////////////////////////////////////////////////////
	c1.OnHTML(".normal--2QYVk >a[href]", func(e *colly.HTMLElement) {
		detailURL := "https://ikman.lk"+e.Attr("href")
		fmt.Println(detailURL)
		GrabDetails(detailURL)
	})
	//////////////////////////////////////////////////////////////////////////////////////
	c1.Visit(URL);
}
func GrabDetails(infoURL string){
	c2 := colly.NewCollector()
	////////////////////////////////////////////////////////////////////////////////////
	//db conc
	db, err := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/ikmandb")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	///////////////////////////////////////////////////////////////////////////////////
	c2.OnHTML(".main-section--34CB3", func(e *colly.HTMLElement) {
		title := e.ChildText(".title--3s1R8")
		subtitle := e.ChildText(".sub-title--37mkY")
		price := e.ChildText(".amount--3NTpl")
		contactName := e.ChildText(".contact-name--m97Sb")
		descr := e.ChildText(".description-section--oR57b > div > .description--1nRbz")

		fmt.Println("==============================")
		fmt.Println(title)
		fmt.Println(subtitle)
		fmt.Println(price)
		fmt.Println(contactName)
		fmt.Println(descr)
		fmt.Println("==============================")


		sql,err := db.Query("INSERT INTO cars(title, subtitle, price, contactName,descr) VALUES (?,?,?,?,?)",title, subtitle, price, contactName,descr)
		if err != nil {
			log.Fatal(err)
		}
		defer sql.Close()

	})
	///////////////////////////////////////////////////////////////////////////////////
	c2.Visit(infoURL)
}

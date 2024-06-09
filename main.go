package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)


func main(){
	
	// h2List:= []string{}
	collector:=colly.NewCollector(
		colly.AllowedDomains("github.com","x.com"),
		
	)

	collector.OnHTML("div",func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error occurred : ",err)
	})

	


	err:=collector.Visit("https://github.com/Numeez")
	if err !=nil{
		fmt.Println("Error  : ",err)
	}

	

}


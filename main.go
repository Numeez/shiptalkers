package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)


func main(){
	collector:=colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org","github.com","x.com"),
		
	)

	collector.OnHTML("body",func(h *colly.HTMLElement) {
		e:=h.ChildAttrs("h2","class")
		fmt.Println(e)
	})

	err:=collector.Visit("https://github.com/Numeez")
	if err !=nil{
		fmt.Println("Error  : ",err)
	}

	

}


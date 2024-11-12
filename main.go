package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Product struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Price string `json:"price"`
	Img   string `json:"img"`
}

var products []Product

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("gatry.com"),
	)

	c.OnHTML("section.promotions.row article", func(h *colly.HTMLElement) {
		product := Product{
			Name:  h.ChildText("div.description h3"),
			Price: h.ChildText("p.price"),
			Img:   h.ChildAttr("div.image a", "src"),
		}

		products = append(products, product)
	})

	//TODO: Get another way to adding search itens for collect the information
	//TODO: Saving information on database
	c.Visit("https://gatry.com/?q=galaxy+buds+2+pro")

	content, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("products.json", content, 0644)
}

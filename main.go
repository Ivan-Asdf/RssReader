package main

import (
	"log"

	"github.com/Ivan-Asdf/RssReader/internal/utils"
	"github.com/Ivan-Asdf/RssReader/pkg/reader"
)

func init() {
	log.SetFlags(0)
}

func main() {
	urls := utils.GetInput()

	results, errors := reader.Parse(urls)
	if len(errors) != 0 {
		log.Println(errors)
	}
	utils.PrintJson(results)
}

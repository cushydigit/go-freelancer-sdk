package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	"github.com/joho/godotenv"
)

var (
	apiAccessToken string
	client         *freelancer.Client
)

func main() {
	Init()
	// QuickExample()
	// ListBudgets()
	// ListCurrencies()
	// ListCategories()
	ListCountries()
	//ListTimezones()

}

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("the .env file not found: %v", err)
	}

	apiAccessToken = os.Getenv("FREELANCER_ACCESS_TOKEN")
	if apiAccessToken == "" {
		log.Fatal("environment variable not set correctly")
	}

	//create instance for freelancer client
	client = freelancer.NewClient(
		apiAccessToken,
		freelancer.WithDebug(true),
	)
}
func QuickExample() {

	// create client with access token
	client = freelancer.NewClient(apiAccessToken)
	opts := freelancer.SearchActiveProjectsOptions{
		FullDescription: freelancer.Bool(true),
		Limit:           freelancer.Int(10),
		Offset:          freelancer.Int(5),
		Query:           freelancer.String("golang"),
	}

	res, err := client.Services.Projects.SearchActive(context.Background(), &opts)
	// set parameters
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for index, p := range res.Result.Projects {
		fmt.Println(index, p)
	}
}

func ListTimezones() {
	res, err := client.Services.Common.ListTimezones(context.Background(), nil)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for index, t := range res.Result.Timezones {
		log.Println(index, t)
	}
}

func ListCountries() {
	res, err := client.Services.Common.ListCountries(context.Background(), nil)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for index, c := range res.Result.Countries {
		fmt.Println(index, c)
	}
}

func ListCurrencies() {
	res, err := client.Services.Projects.Extras.Currencies.List(context.Background(), nil)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for index, c := range res.Result.Currencies {
		fmt.Println(index, c)
	}
}

func ListBudgets() {
	res, err := client.Services.Projects.Extras.Budgets.List(context.Background(), nil)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for index, b := range res.Result.Budgets {
		fmt.Println(index, b)
	}
}

func ListCategories() {
	res, err := client.Services.Projects.Extras.Categories.List(context.Background(), nil)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for index, c := range res.Result.Categories {
		fmt.Println(index, c)
	}
}

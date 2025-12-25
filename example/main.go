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
	log.Println(apiAccessToken)
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
	client = freelancer.NewClient(apiAccessToken) // create client with access token
	opts := freelancer.SearchActiveProjectsOptions{
		FullDescription: freelancer.Bool(true),
		Limit:           freelancer.Int(10),
		Offset:          freelancer.Int(5),
	}
	res, err := client.Services.Projects.SearchActive(context.Background(), &opts)
	// set parameters
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for index, p := range res.Result.Projects {
		fmt.Printf("Project-%03d\tID-%d\tAt: %d", index, p.ID, p.TimeSubmitted)
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
	s := client.Services.Projects.Extras.Currencies.List()
	res, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	for index, c := range res.Result.Currencies {
		fmt.Printf("Currency-%03d\tID-%d\tName: %s\tCode: %s\tCountry:%s\tSign:%s\n", index, c.ID, c.Name, c.Code, c.Country, c.Sign)
	}
}

func ListBudgets() {
	s := client.Services.Projects.Extras.Budgets.List()
	res, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	for index, b := range res.Result.Budgets {
		fmt.Printf("Budget-%03d\tName-%s\tMin: %f\tMax: %f\tCurrencyID: %d\n", index, b.Name, b.Minimum, b.Maximum, b.CurrencyID)
	}
}

func ListCategories() {
	s := client.Services.Projects.Extras.Categories.List()
	res, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	for index, c := range res.Result.Categories {
		fmt.Printf("Category-%03d\tID-%d\tName: %s\n", index, c.ID, c.Name)
	}
}

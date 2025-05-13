package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	freelancer "github.com/shahinrahimi/go-freelancer-sdk/freelancer/v1"
)

var (
	apiAccessToken string
	client         *freelancer.Client
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("the .env file not found: %v", err)
	}

	apiAccessToken = os.Getenv("FREELANCER_ACCESS_TOKEN")
	if apiAccessToken == "" {
		log.Fatal("environment varaiable not set correctly")
	}

	//create instance for freelancer client
	client = freelancer.NewClient(apiAccessToken)

}

func ListActiveProjects() {
	// create active project service
	s := client.NewListActiveProjectsService()
	// fetch result
	resp, err := s.Do(context.Background())
	// hanlde error
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	// use easily typed response
	for index, p := range resp.Result.Projects {
		timeSubmited := time.Unix(p.TimeSubmitted, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("Project-%03d\tID-%d\tAt: %s\tTitle: %s\n", index, p.ID, timeSubmited, p.Title)
	}
}

func ListActiveProjectsWithMoreOptions() {
	// create active project service
	s := client.NewListActiveProjectsService()
	// exlude projects base on the query
	// query with space seprated
	s.SetQuery("go python")
	// exclude projects with type
	s.SetProjectTypes([]freelancer.ProjectType{freelancer.ProjectTypeFixed})
	// include full Description
	s.SetFullDescription(true)
	// fetch result
	resp, err := s.Do(context.Background())

	// hanlde error
	if err != nil {
		log.Printf("error: %v", err)
	}

	// use easily typed response
	for i, p := range resp.Result.Projects {
		timeSubmited := time.Unix(p.TimeSubmitted, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("Project-%03d\tID-%d\tAt: %s\tSeoUrl: %s\n", i, p.ID, timeSubmited, p.SeoURL)
	}

}

func ListBudgets() {
	s := client.NewListBudgetsService()

	resp, err := s.Do(context.Background())

	// handle erro
	if err != nil {
		log.Printf("error: %v", err)
	}

	for i, b := range resp.Result.Budgets {
		fmt.Printf("Budget-%03d\tName: %s\tProjectType:%s CurrencyID:%d Minimum:%f Maximum:%f\n", i, b.Name, b.ProjectType, b.CurrencyID, b.Maximum, b.Minimum)
	}
}

func ListCurrencies() {
	s := client.NewListCurrenciesService()
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	for i, c := range resp.Result.Currencies {
		fmt.Printf("Currency-%03d\tID-%d\tName: %s\tCode: %s\tCountry:%s\tSign:%s\n", i, c.ID, c.Name, c.Code, c.Country, c.Sign)
	}
}

func ListTimeZones() {
	s := client.NewListTimezonesService()
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	for i, t := range resp.Result.Timezones {
		fmt.Printf("Timezone-%03d\t Timezone: %v\t ID: %d\tCountry: %s\tOffset: %f\n", i, t.Timezone, t.ID, t.Country, t.Offset)
	}
}

func ListSelfDevices() {
	s := client.NewListSelfDevicesService()
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	for i, d := range resp.Result.Devices {
		lastLogin := time.Unix(d.LastLogin, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("Device-%02d\tCountry: %s\tCity: %s\tPlatform: %s\tUserAgent: %s\tLastlogin: %s", i, d.Country, d.City, d.Platform, d.UserAgent, lastLogin)
	}
}

func main() {
	Init()
	ListSelfDevices()
	// ListTimeZones()
	// ListCurrencies()
	// ListBudgets()
	// ListActiveProjects()
	// ListActiveProjectsWithMoreOptions()
}

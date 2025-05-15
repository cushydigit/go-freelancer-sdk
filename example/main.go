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
		log.Printf("error: %v", err)
	}

	// use easily typed response
	for index, p := range resp.Result.Projects {
		timeSubmited := time.Unix(p.TimeSubmitted, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("Project-%03d\tID-%d\tAt: %s\tTitle: %s\n", index, p.ID, timeSubmited, p.Title)
	}
}

func ListActiveLimitProjects() {
	s := client.NewListActiveProjectsService()
	s.SetLimit(10)
	s.SetOffset(10)
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	for i, p := range resp.Result.Projects {
		timeSubmited := time.Unix(p.TimeSubmitted, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("Project-%03d\tID-%d\tAt: %s\tSeoUrl: %s\n", i, p.ID, timeSubmited, p.SeoURL)
	}
	s.SetLimit(10)
	s.SetOffset(15)
	resp, err = s.Do(context.Background())
	if err != nil {
		log.Printf("erro: %v", err)
	}
	for i, p := range resp.Result.Projects {
		timeSubmited := time.Unix(p.TimeSubmitted, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("Project-%03d\tID-%d\tAt: %s\tSeoUrl: %s\n", i, p.ID, timeSubmited, p.SeoURL)
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

func ListCategories() {
	s := client.NewListCategoriesService()
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for i, c := range resp.Result.Categories {
		fmt.Printf("Category-%02d \tID:%3d\tCategoryName: %s\n", i, c.ID, c.Name)
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

func GetSelfInfo() {
	s := client.NewGetSelfInfoService()
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	info := resp.Result
	fmt.Printf("UserID: %d, Diplayname: %s, PublicName: %s, Role: %s\n", info.ID, info.DisplayName, info.PublicName, info.Role)

	s2 := client.NewGetUserService()
	resp2, err := s2.Do(context.Background(), info.ID)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	user := resp2.Result
	fmt.Printf("UserID: %d, Diplayname: %s, PublicName: %s, Role: %s\n", user.ID, user.DisplayName, user.PublicName, user.Role)

	s3 := client.NewListUsersService()
	s3.SetUsers([]int64{info.ID})
	resp3, err := s3.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for _, u := range resp3.Result.Users {
		fmt.Printf("UserID: %d, Diplayname: %s, PublicName: %s, Role: %s\n", u.ID, u.DisplayName, u.PublicName, u.Role)

	}
}

func ListFreelancerService() {
	s := client.NewListFreelancersService()
	s.SetOnlineOnly(true).SetReviewCountMax(300).SetReputation(true)
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for _, u := range resp.Result.Users {
		fmt.Printf("UserID: %8d\tRole: %s\tReviewCount: %3d\tDisplayName: %s\n", u.ID, u.Role, u.Reputation.EntireHistory.Reviews, u.DisplayName)
	}
}

func ListSelfJobs() {
	s := client.NewGetSelfInfoService()
	s.SetJobs(true)
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for _, j := range resp.Result.Jobs {
		fmt.Printf("ID: %d\tName: %s\tCategoryName: %s\n", j.ID, *j.Name, j.Category.Name)
	}

}

func AddSelfJobs() {
	s := client.NewAddSelfJobsService()
	reqBody := freelancer.JobsRequestBody{
		Jobs: []int32{1, 2},
	}
	resp, err := s.Do(context.Background(), reqBody)
	if err != nil {
		log.Printf("error %v", err)
		return
	}
	fmt.Printf("ruquested to add a jobs with status %s\n", resp.Status)
}

func main() {
	Init()
	// ListSelfJobs()
	// AddSelfJobs()
	// ListSelfJobs()
	ListFreelancerService()
	// ListCategories()
	// GetSelfInfo()
	// ListActiveLimitProjects()
	// ListSelfDevices()
	// ListTimeZones()
	// ListCurrencies()
	// ListBudgets()
	// ListActiveProjects()
	// ListActiveProjectsWithMoreOptions()
}

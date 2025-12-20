package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

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
	ListBudgets()
	ListCurrencies()
	ListCategories()
	ListTimezones()

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
	client = freelancer.NewClient(apiAccessToken)
	client.Debug = true
}

func QuickExample() {
	// use environment variable or set accessToken manually
	client = freelancer.NewClient(apiAccessToken)
	// create service for fetching active projects
	s := client.Services.Projects.SearchActive()
	// set parameters
	// fetch full description
	s.SetFullDescription(true)
	// limit the number of results
	s.SetLimit(10)
	// filter by project type
	s.SetProjectTypes([]freelancer.ProjectType{freelancer.ProjectFixed})
	// execute request
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	// use helper function for parsing response
	var res freelancer.SearchActiveProjectsResponse
	b, _ := json.Marshal(resp)
	_ = json.Unmarshal(b, &res)
	// print results
	for index, p := range res.Result.Projects {
		timeSubmitted := time.Unix(p.TimeSubmitted, 0).Format("2006-01-02 15:04:05")
		if len(p.Title) > 10 {
			p.Title = p.Title[0:10] + " *****"
		}
		fmt.Printf("Project-%03d\tID-%d\tAt: %s\tTitle: %s\n", index, p.ID, timeSubmitted, p.Title)
	}
}

func ListTimezones() {
	s := client.Services.Common.ListTimezones()
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	res := freelancer.ListTimezonesResponse{}
	b, _ := json.Marshal(resp)
	_ = json.Unmarshal(b, &res)
	for index, t := range res.Result.Timezones {
		fmt.Printf("Timezone-%03d\tID-%d\tCountry: %s\tTimezones: %s\n", index, t.ID, t.Country, t.Timezone)
	}
}

func ListCurrencies() {
	s := client.Services.Projects.Extras.Currencies.List()
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	res := freelancer.ListCurrenciesResponse{}
	b, _ := json.Marshal(resp)
	_ = json.Unmarshal(b, &res)
	for index, c := range res.Result.Currencies {
		fmt.Printf("Currency-%03d\tID-%d\tName: %s\tCode: %s\tCountry:%s\tSign:%s\n", index, c.ID, c.Name, c.Code, c.Country, c.Sign)
	}
}

func ListBudgets() {
	s := client.Services.Projects.Extras.Budgets.List()
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	res := freelancer.ListBudgetsResponse{}
	b, _ := json.Marshal(resp)
	_ = json.Unmarshal(b, &res)
	for index, b := range res.Result.Budgets {
		fmt.Printf("Budget-%03d\tName-%s\tMin: %f\tMax: %f\tCurrencyID: %d\n", index, b.Name, b.Minimum, b.Maximum, b.CurrencyID)
	}
}

func ListCategories() {
	s := client.Services.Projects.Extras.Categories.List()
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	res := freelancer.ListCategoriesResponse{}
	b, _ := json.Marshal(resp)
	_ = json.Unmarshal(b, &res)
	for index, c := range res.Result.Categories {
		fmt.Printf("Category-%03d\tID-%d\tName: %s\n", index, c.ID, c.Name)
	}
}

// func ListSelfDevices() {
// 	s := client.NewListUserLoginDevicesService()
// 	resp, err := s.DO(context.Background())
// 	if err != nil {
// 		log.Printf("error: %v", err)
// 	}
// 	for i, d := range resp.Result.Devices {
// 		lastLogin := time.Unix(d.LastLogin, 0).Format("2006-01-02 15:04:05")
// 		fmt.Printf("Device-%02d\tCountry: %s\tCity: %s\tPlatform: %s\tUserAgent: %s\tLastLogin: %s", i, d.Country, d.City, d.Platform, d.UserAgent, lastLogin)
// 	}
// }

// func GetSelfInfo() {
// 	s := client.NewGetSelfInfoService()
// 	resp, err := s.DO(context.Background())
// 	if err != nil {
// 		log.Printf("error: %v", err)
// 		return
// 	}
// 	info := resp.Result
// 	fmt.Printf("UserID: %d, DisplayName: %s, PublicName: %s, Role: %s\n", info.ID, info.DisplayName, info.PublicName, info.Role)

// 	s2 := client.NewGetUserService()
// 	resp2, err := s2.Do(context.Background(), info.ID)
// 	if err != nil {
// 		log.Printf("error: %v", err)
// 		return
// 	}
// 	user := resp2.Result
// 	fmt.Printf("UserID: %d, DisplayName: %s, PublicName: %s, Role: %s\n", user.ID, user.DisplayName, user.PublicName, user.Role)

// 	s3 := client.NewListUsersService()
// 	s3.SetUsers([]int64{info.ID})
// 	resp3, err := s3.Do(context.Background())
// 	if err != nil {
// 		log.Printf("error: %v", err)
// 		return
// 	}
// 	for _, u := range resp3.Result.Users {
// 		fmt.Printf("UserID: %d, DisplayName: %s, PublicName: %s, Role: %s\n", u.ID, u.DisplayName, u.PublicName, u.Role)

// 	}
// }

// func ListFreelancerService() {
// 	s := client.NewSearchFreelancersService()
// 	s.SetOnlineOnly(true).SetReviewCountMax(300).SetReputation(true)
// 	resp, err := s.Do(context.Background())
// 	if err != nil {
// 		log.Printf("error: %v", err)
// 		return
// 	}
// 	usersIDs := []int64{}
// 	for _, u := range resp.Result.Users {
// 		fmt.Printf("UserID: %8d\tRole: %s\tReviewCount: %3d\tDisplayName: %s\n", u.ID, u.Role, u.Reputation.EntireHistory.Reviews, u.DisplayName)
// 		usersIDs = append(usersIDs, u.ID)
// 	}

// 	// fetch reputation
// 	// s2 := client.NewListUsersReputationsService()
// 	// s2.SetUsers(usersIDs)
// 	// resp2, err := s2.DO(context.Background())
// 	// if err != nil {
// 	// 	log.Printf("error: %v", err)
// 	// 	return
// 	// }

// 	// fmt.Printf("%s", string(resp2.Result))
// 	// for id, r := range resp2.Result {
// 	// 	fmt.Printf("UserID: %8s\tReviewCount: %d\n", id, r.EntireHistory.Reviews)
// 	// }
// 	// fmt.Printf("Result %v", string(resp2.Result))
// 	// fetch users portfolios
// 	s3 := client.NewListUsersPortfoliosService()
// 	s3.SetUsers([]int64{usersIDs[3]})
// 	resp3, err := s3.DO(context.Background())
// 	if err != nil {
// 		log.Printf("error : %v", err)
// 		return
// 	}
// 	fmt.Printf("result: \n%s", string(resp3.Result))
// }

// func ListSelfJobs() {
// 	s := client.NewGetSelfInfoService()
// 	s.SetJobs(true)
// 	resp, err := s.DO(context.Background())
// 	if err != nil {
// 		log.Printf("error: %v", err)
// 		return
// 	}
// 	for _, j := range resp.Result.Jobs {
// 		fmt.Printf("ID: %d\tName: %s\tCategoryName: %s\n", j.ID, *j.Name, j.Category.Name)
// 	}

// }

// func AddSelfJobs() {
// 	s := client.NewAddSelfJobsService()
// 	reqBody := freelancer.JobsRequestBody{
// 		Jobs: []int32{1, 2},
// 	}
// 	resp, err := s.DO(context.Background(), reqBody)
// 	if err != nil {
// 		log.Printf("error %v", err)
// 		return
// 	}
// 	fmt.Printf("requested to add a jobs with status %s\n", resp.Status)
//}

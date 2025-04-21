package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	freelancer "github.com/shahinrahimi/go-freelancer-sdk/freelancer/v1"
)


func quickExmaple() {
	// create freelancer clinet
	c := freelancer.NewClient("YOUR-ACCESS-TOKEN")

	// create active project service 
	s := c.NewListActiveProjectsService()

	// get results
	res, err := s.Do(context.Background())

	// basic error handler
	if err != nil {
		log.Printf("error: %v", err)
	} 

	// loop through projects
	for index, p := range res.Result.Projects {
		log.Printf("Project-%d\tTitle: ", index, p.Title)
	}
}


func quickExample2() {
	// create freelancer clinet
	c := freelancer.NewClient("YOUR-ACCESS-TOKEN")

	// create active project service 
	s := c.NewListActiveProjectsService()

	// list projects base on query
	s.SetQuery("Go Python")

	// filter projects base on type
	s.SetProjectTypes([]freelancer.ProjectType{freelancer.ProjectTypeFixed})

	// include projects full description
	s.SetFullDescription(true)

	// get results
	res, err := s.Do(context.Background())

	// basic error handler
	if err != nil {
		log.Printf("error: %v", err)
	} 

	// loop through projects
	for index, p := range res.Result.Projects {
		log.Printf("Project-%d\ttype: \nDesctiption: %s", index, p.Title, p.Description)
	}

}


func main() {

	// load .env file 
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// get freelancer token
	token := os.Getenv("FREELANCER_ACCESS_TOKEN")
	if token == "" {
		panic("FREELANCER_ACCESS_TOKEN is not set")
	}
	// create freelancer client
	c := freelancer.NewClient(token)

	// set BaseUrl main or sandbox freelancer.con api
	c.SetBaseUrl(freelancer.BaseAPIMainURL)
	
	// list active projects
	service := c.NewListActiveProjectsService()

	pas.SetQuery("python json")
	pas.SetProjectTypes([]freelancer.ProjectType{freelancer.ProjectTypeFixed})
	pas.SetProjectUpgrades([]freelancer.ProjectUpgradeType{freelancer.ProjectUpgradeTypeSealed})
	pas.SetFullDescription(true)
	pas.SetUserCountryDetails(true)
	pas.SetUserLocationDetails(true)
	res, err := pas.Do(context.Background())
	if err != nil {
		if handledError, ok := err.(*freelancer.APIError2); ok {
			log.Printf("Handled error: %s", handledError)
			return
		}
		log.Printf("Unhandled Error: %s", err)
	}
	ownersID := make([]int, 0)
	projectsID := make([]int, 0)
	for pr, p := range res.Result.Projects {
		log.Printf("Project #%d: %s", pr, p.Title)
		log.Printf("Project #%d: %s", pr, p.SeoURL)
		//log.Printf("Project #%d: %s", pr, p.Description)
		log.Printf("Project #%d: %s", pr, p.PreviewDesc)
		log.Printf("Project #%d: %s", pr, p.Location.City)
		ownersID = append(ownersID, p.OwnerID)
		projectsID = append(projectsID, p.ID)
	}

	us := c.NewListUsersService()
	us.SetUsers(ownersID)
	us.SetAvatar(true)
	res2, err := us.Do(context.Background())
	if err != nil {
		if handledError, ok := err.(*freelancer.APIError2); ok {
			log.Printf("Handled error: %s", handledError)
			return
		}
		log.Printf("Unhandled Error: %s", err)
		return
	}
	for pr, p := range res2.Result.Users {
		log.Printf("User #%s: %s", pr, p.Username)
		log.Printf("User #%s: %s", pr, p.Avatar)
	}

	res3, err := c.NewListBudgetsService().Do(context.Background())
	if err != nil {
		if handledError, ok := err.(*freelancer.APIError2); ok {
			log.Printf("Handled error: %s", handledError)
			return
		}
		log.Printf("Unhandled Error: %s", err)
		return
	}
	for pr, p := range res3.Result.Budgets {
		log.Printf("Budget #%d with minimum: %f and maximum: %f , name: %s and project type: %s", pr, p.Minimum, p.Maximum, p.Name, p.ProjectType)
	}

	res4, err := c.NewListCountriesService().Do(context.Background())
	if err != nil {
		if handledError, ok := err.(*freelancer.APIError2); ok {
			log.Printf("Handled error: %s", handledError)
			return
		}
		log.Printf("Unhandled Error: %s", err)
		return
	}
	for pr, p := range res4.Result.Countries {
		log.Printf("Country #%d: %s, %s", pr, p.Name, p.Code)
	}

	res5, err := c.NewListCategoriesService().Do(context.Background())
	if err != nil {
		if handledError, ok := err.(*freelancer.APIError2); ok {
			log.Printf("Handled error: %s", handledError)
			return
		}
		log.Printf("Unhandled Error: %s", err)
		return
	}
	for pr, p := range res5.Result.Categories {
		log.Printf("Category #%d: %s", pr, p.Name)
	}

	ps := c.NewListProjectsService()
	ps.SetProjects(projectsID)
	ps.SetLocationDetails(true)
	ps.SetUserLocationDetails(true)
	res6, err := ps.Do(context.Background())
	if err != nil {
		if handledError, ok := err.(*freelancer.APIError2); ok {
			log.Printf("Handled error: %s", handledError)
			return
		}
		log.Printf("Unhandled Error: %s", err)
		return
	}
	for pr, p := range res6.Result.Projects {
		log.Printf("Project #%d: %s", pr, p.Title)
		log.Printf("Project #%d: %s", pr, p.SeoURL)
		//log.Printf("Project #%d: %s", pr, p.Description)
		log.Printf("Project #%d: %s", pr, p.Location.City)
	}

	res7, err := c.NewListCurrenciesService().Do(context.Background())
	if err != nil {
		if handledError, ok := err.(*freelancer.APIError2); ok {
			log.Printf("Handled error: %s", handledError)
			return
		}
		log.Printf("Unhandled Error: %s", err)
		return
	}
	for pr, p := range res7.Result.Currencies {
		log.Printf("Currency #%d: %s", pr, p.Name)
	}

}

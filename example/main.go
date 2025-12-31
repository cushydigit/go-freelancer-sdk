package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/cushydigit/go-freelancer-sdk/freelancer"
	"github.com/joho/godotenv"
	"golang.org/x/net/proxy"
)

var (
	apiAccessToken string
	client         *freelancer.Client
)

func main() {
	// Init()
	InitWithProxy()
	QuickExample()
	// ListBudgets()
	// ListCurrencies()
	// ListCategories()
	// ListCountries()
	//ListTimezones()

}

func Init() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	apiAccessToken = os.Getenv("FREELANCER_ACCESS_TOKEN")
	if apiAccessToken == "" {
		log.Fatal("'FREELANCER_ACCESS_TOKEN' environment  variable is not set or emtpy")
	}

	//create instance for freelancer client
	client = freelancer.NewClient(
		apiAccessToken,
		freelancer.WithDebug(true),
	)
}

func InitWithProxy() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}
	proxyAddr := os.Getenv("PROXY_ADDR")
	if proxyAddr == "" {
		log.Fatal("'PROXY_ADDR' environment  variable is not set or empty")
	}

	apiAccessToken = os.Getenv("FREELANCER_ACCESS_TOKEN")
	if apiAccessToken == "" {
		log.Fatal("'FREELANCER_ACCESS_TOKEN' environment  variable is not set or emtpy")
	}

	dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
	if err != nil {
		log.Fatalf("failed to create SOCKS5 dialer: %v", err)
	}

	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	client = freelancer.NewClient(
		apiAccessToken,
		freelancer.WithDebug(true),
		freelancer.WithHttpClient(httpClient),
	)

}
func QuickExample() {

	// create client with access token
	opts := freelancer.SearchActiveProjectsOptions{
		FullDescription: freelancer.Bool(true),
		Limit:           freelancer.Int(10),
		Offset:          freelancer.Int(5),
		Query:           freelancer.String("golang python"),
	}

	res, err := client.Services.Projects.SearchActive(context.Background(), &opts)
	// set parameters
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for index, p := range res.Result.Projects {
		fmt.Println(index, p.GetFullUrl())
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

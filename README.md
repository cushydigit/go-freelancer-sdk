<h1 align="center">🛠️ Unofficial Freelancer.com SDK (Go) </h1>
<p align="center">
    <a href="https://pkg.go.dev/github.com/cushydigit/go-freelancer-sdk">
        <img src="https://pkg.go.dev/badge/github.com/cushydigit/go-freelancer-sdk.svg" alt="Go Reference">
    </a>
    <a href="https://github.com/cushydigit/go-freealncer-sdk/LICENSE">
        <img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License MIT">
    </a>
    <a href="go.mod">
        <img src="https://img.shields.io/github/go-mod/go-version/cushydigit/go-freelancer-sdk" alt="Go Version">
    </a>
    <a href="https://github.com/yourusername/freelancer-go-sdk/actions">
        <img src="https://img.shields.io/github/actions/workflow/status/cushydigit/go-freelancer-sdk/test.yml?branch=main" alt="Build Status">
    </a>
    <img src="https://img.shields.io/badge/status-unofficial-orange" alt="Unofficial Status">
  </a>
</p>


# 📦 Freelancer.com Go SDK
A Go (Golang) SDK for interacting with the [Freelancer.com API](https://developers.freelancer.com/).

This library provides a simple, typed client for accessing Freelancer.com services like projects, bids, users, currencies, and more.

---

## ✨ Features

- Easy authentication using API Key
- Typed requests and responses
- Covers major Freelancer.com endpoints
- Clean, idiomatic Go design
- MIT Licensed — free to use and modify

---

## 📚 Installation

```bash
go get github.com/cushydigit/freelancer-go-sdk
```

## 🚀 Quick Example
fetch for search project

```Go
func QuickExample() {

	// use environment variable or set accessToken manually
	if err := godotenv.Load(); err != nil {
		log.Fatalf("the .env file not found: %v", err)
	}
	apiAccessToken = os.Getenv("FREELANCER_ACCESS_TOKEN")
	if apiAccessToken == "" {
		log.Fatal("environment variable not set correctly")
	}

	client = freelancer.NewClient(apiAccessToken)
	// create service for fetching active projects
	s := client.NewProjectsSearchActiveService()
	// set parameters
	s.SetFullDescription(true) // fetch full description
	s.SetLimit(10)i // limit the number or results
	s.SetProjectTypes([]freelancer.ProjectType{freelancer.ProjectFixed}) // filter by project type
	// make request
	resp, err := s.Do(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	// use helper function for parsing response
	var res helper.SearchActiveProjectsResponse
	b, _ := json.Marshal(resp)
	_ = json.Unmarshal(b, &res)
	// print results
	for index, p := range res.Result.Projects {
		timeSubmitted := time.Unix(p.TimeSubmitted, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("Project-%03d\tID-%d\tAt: %s\tTitle: %s\n", index, p.ID, timeSubmitted, p.Title)
	}
}
```
output:
```bash

Project-000     ID-40011543     At: 2025-11-25 16:24:37 Title: Figma-to-HTML Responsive Designer For Website Update -- 3
Project-001     ID-40011529     At: 2025-11-25 16:20:21 Title: Gothic Irish Novel Cover Design
Project-002     ID-40011493     At: 2025-11-25 16:19:57 Title: Vibrant website Visual redesign based on our mock up version. 
Project-003     ID-40011526     At: 2025-11-25 16:19:23 Title: Comprehensive Digital Marketing 
Project-004     ID-40011480     At: 2025-11-25 16:18:44 Title: Photographer's Portfolio Website Development
Project-005     ID-40011364     At: 2025-11-25 16:18:04 Title: DC-DC Boost Converter with MCU
Project-006     ID-40011525     At: 2025-11-25 16:17:56 Title: Application Security Pen-Test Needed
Project-007     ID-40011523     At: 2025-11-25 16:17:00 Title: Digital Marketer Needed for SSH Launch
Project-008     ID-40010912     At: 2025-11-25 16:16:26 Title: English-Arabic Document Translation
Project-009     ID-40011520     At: 2025-11-25 16:15:15 Title: Wedding Family Video Edit
```
```
func example() {
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

```
```
func example2() {
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
```
## 🛠️ Usage

## 📂 Project Structure

## 📄 License

## 🤝 Contributing

## 📣 Contact

Feel free to reach out if you have question or suggestions:
- 📧 [sh.rahimi.dev@gmail.com](mailto:sh.rahimi.dev@gmail.com)
- 💼 [LinkedIn](https://www.linkedin.com/in/shahin-rahimi-828447254/)
- 🧑‍💻 [Freelancer](https://www.freelancer.com/u/shahinrahimi)
- 🕊️ [X (formerly Twitter)](https://x.com/cushydigit)






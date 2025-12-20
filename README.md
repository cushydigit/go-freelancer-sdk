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

# 🚀 Freelancer.com Go SDK

A Go (Golang) SDK for interacting with the [Freelancer.com API](https://developers.freelancer.com/).

This library provides a simple, typed client for accessing Freelancer.com services like projects, bids, users, currencies, and more.

---

## 💡 Why I built this SDK?

As a freelancer, I spend a lot of time on the platform. I wanted to build automated tools (like bidding bots and project monitors) to streamline my workflow, but I realized there was no comprehensive Go client available.

I spent a significant amount of time hand-coding these service wrappers to handle the platform's API nuances. This project is built out of necessity to give Go developers the same power that Python developers have on the platform.

## ✨ Features

- **Authentication:** Easy authentication using API Key
- **Endpoints:** Covers major Freelancer.com endpoints (Users, Projects, Common)
- **Design**: Clean, idiomatic Go design
- **Context Support:** All methods support `context.Context` for timeouts and cancellation.
- **License:** MIT Licensed — free to use and modify

---

## 📦 Installation

```bash
go get github.com/cushydigit/freelancer-go-sdk
```

## 🛠️ Usage

### Authentication

You will need an OAuth2 access token from Freelancer.com. You can generate one in the [Freelancer Developer Portal](https://accounts.freelancer.com/).

### Quick Examples

fetch active projects

```Go
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


```

output:

```bash
Project-000     ID-40011579     At: 2025-11-25 16:40:39 Title: Drive Foot *****
Project-001     ID-40011578     At: 2025-11-25 16:40:23 Title: Mumbai Ins *****
Project-002     ID-40011577     At: 2025-11-25 16:40:04 Title: Modern Wix *****
Project-003     ID-40011576     At: 2025-11-25 16:39:58 Title: Classic Br *****
Project-004     ID-40011565     At: 2025-11-25 16:39:44 Title: Social Med *****
Project-005     ID-40011572     At: 2025-11-25 16:39:01 Title: Health Clu *****
Project-006     ID-40011570     At: 2025-11-25 16:38:59 Title: Cross-Plat *****
Project-007     ID-40011571     At: 2025-11-25 16:38:01 Title: TDS Reconc *****
Project-008     ID-40011567     At: 2025-11-25 16:35:48 Title: Responsive *****
Project-009     ID-40011566     At: 2025-11-25 16:35:39 Title: JOB POSTIN *****
```

fetch timezones

```GO
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
```

fetch currencies

```GO
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
```

fetch budgets

```Go
func ListBudgets() {
 s := client.Services.Projects.Extras.Budgets.List()
 resp, err := s.Do(context.Background())
 if err != nil {
  log.Printf("error: %v", err)
 }
 res := helper.ListBudgetsResponse{}
 b, _ := json.Marshal(resp)
 _ = json.Unmarshal(b, &res)
 for index, b := range res.Result.Budgets {
  fmt.Printf("Budget-%03d\tName-%s\tMin: %f\tMax: %f\tCurrencyID: %d\n", index, b.Name, b.Minimum, b.Maximum, b.CurrencyID)
 }
}
```

fetch categories

```Go
func ListCategories() {
 s := client.NewCategoriesListService()
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
```

## 📂 Project Structure

This SDK follows a modular service design. All core logic is located in `freelancer`.

- **`client.go`**: The main entry point and configuration.
- **`models.go` / `enums.go`**: Shared data structures and constants.
- **`service_*.go`**: Individual files for each API endpoint action.

## 📚 Documentation

Full documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/cushydigit/go-freelancer-sdk).

For details on the underlying API endpoints and parameters, refer to the official [Freelancer.com API Documentation](https://developers.freelancer.com/).

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development

We use a `Makefile` to automate common tasks:

```bash
make test    # Run all unit tests
make build   # Compile the project
make fmt     # Format code
```

## 🗺️ Roadmap

Current version covers **Projects**, **Users**, and **Common** services.

### Stability & Quality

- [ ] **Static Analysis:**
- [ ] **Unit Testing:**
- [ ] **Use Case:**

### Upcoming Features

- [ ] **Messaging:** Threads and direct message handling.
- [ ] **Contests:** Browsing and participating in contests.

## ⚠️ Disclaimer

This is an unofficial library and is not affiliated with, endorsed by, or associated with Freelancer.com.
Please ensure you comply with the [Freelancer API Terms](https://www.freelancer.com/about/terms) and Conditions when using this software.

## 📣 Contact

Feel free to reach out if you have question or suggestions:

- 📧 [sh.rahimi.dev@gmail.com](mailto:sh.rahimi.dev@gmail.com)
- 💼 [LinkedIn](https://www.linkedin.com/in/shahin-rahimi-828447254/)
- 🧑‍💻 [Freelancer](https://www.freelancer.com/u/shahinrahimi)
- 🕊️ [X](https://x.com/cushydigit)

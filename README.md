<!-- markdownlint-disable MD033 -->
<h1 align="center">️Unofficial Freelancer.com SDK (Go)</h1>

<p align="center">
  <a href="https://pkg.go.dev/github.com/cushydigit/go-freelancer-sdk">
    <img src="https://pkg.go.dev/badge/github.com/cushydigit/go-freelancer-sdk.svg" alt="Go Reference">
  </a>
  <a href="https://github.com/cushydigit/go-freelancer-sdk/LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License MIT">
  </a>
  <a href="go.mod">
    <img src="https://img.shields.io/github/go-mod/go-version/cushydigit/go-freelancer-sdk" alt="Go Version">
  </a>
  <a href="https://github.com/cushydigit/go-freelancer-sdk/actions">
    <img src="https://img.shields.io/github/actions/workflow/status/cushydigit/go-freelancer-sdk/test.yml?branch=main" alt="Build Status">
  </a>
  <img src="https://img.shields.io/badge/status-unofficial-orange" alt="Unofficial Status">
</p>
<!-- markdownlint-enable MD033 -->

	| Syntax | Description |
| ----------- | ----------- |
| Header | Title |
| Paragraph | Text |

# Freelancer.com Go SDK

A Go (Golang) SDK for interacting with the [Freelancer.com API](https://developers.freelancer.com/).

This library provides a simple, typed client for accessing Freelancer.com services like projects, bids, users, currencies, and more.

---

## Why I built this SDK?

As a freelancer, I spend a lot of time on the platform. I wanted to build automated tools (like bidding bots and project monitors) to streamline my workflow, but I realized there was no comprehensive Go client available.

I spent a significant amount of time hand-coding these service wrappers to handle the platform's API nuances. This project is built out of necessity to give Go developers the same power that Python developers have on the platform.

## Features

- **Authentication:** Easy authentication using API Key
- **Endpoints:** Covers major Freelancer.com endpoints (Users, Projects, Common)
- **Design**: Clean, idiomatic Go design
- **Context Support:** All methods support `context.Context` for timeouts and cancellation.
- **License:** MIT Licensed — free to use and modify

---

## Installation

```bash
go get github.com/cushydigit/freelancer-go-sdk
```

## Usage

### Authentication

You will need an OAuth2 access token from Freelancer.com. You can generate one in the [Freelancer Developer Portal](https://accounts.freelancer.com/).

### Quick Examples

fetch active projects

```Go
import (
 "log"
 "github.com/cushydigit/go-freelancer-sdk/freelancer"
)
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
  log.Println(index, p)
 }
```

fetch timezones

```GO
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
```

fetch countries

```GO
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

```

fetch budgets

```Go

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
```

fetch categories

```Go
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
```

## Project Structure

This SDK follows a modular service design. All core logic is located in `freelancer`.

```bash
  ├── client.go                           
  ├── endpoints.go
  ├── enums.go
  ├── errors.go
  ├── ratelimiter.go
  ├── responses.go
  ├── service_common.go
  ├── service_projects_bids.go
  ├── service_projects_collaborations.go
  ├── service_projects_extra.go
  ├── service_projects_jobs.go
  ├── service_projects_milestones.go
  ├── service_projects_reviews.go
  ├── service_projects_services.go
  ├── service_projects.go
  ├── service_users_extras.go
  ├── service_users_profiles.go
  ├── service_users_self.go
  ├── service_users.go
  ├── services.go
  ├── types.go
  └── utils.go
```

- **`client.go`**: It holds core logic
- **`endpoints.go`**: Contains all your constants for URLs
- **`types.go`**: Shared data structures
- **`responses`**: The wrappers for API replies
- **`enums.go`**: Custom types and constants for statuses, roles, and types
- **`services.go`**: The entry point for all services. It initializes all services when the client is created
- **`service_*.go`**: Resource-oriented service implementations. Each file encapsulates logic for a specific API domain

## Documentation

Full documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/cushydigit/go-freelancer-sdk).

For details on the underlying API endpoints and parameters, refer to the official [Freelancer.com API Documentation](https://developers.freelancer.com/).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development

We use a `Makefile` to automate common tasks:

```bash
make test    # Run all unit tests
make build   # Compile the project
```

## Roadmap

Current version covers **Projects**, **Users**, and **Common** services.

### Stability & Quality

- [x] **Static Analysis:**
- [x] **Unit Testing:** (7% coverage)
- [ ] **Use Case:**

### Upcoming Features

- [ ] **Messaging:** Threads and direct message handling.
- [ ] **Contests:** Browsing and participating in contests.

## Disclaimer

This is an unofficial library and is not affiliated with, endorsed by, or associated with Freelancer.com.
Please ensure you comply with the [Freelancer API Terms](https://www.freelancer.com/about/terms) and Conditions when using this software.

## Contact

Feel free to reach out if you have question or suggestions:

- 📧 [sh.rahimi.dev@gmail.com](mailto:sh.rahimi.dev@gmail.com)
- 💼 [LinkedIn](https://www.linkedin.com/in/shahin-rahimi-828447254/)
- 🧑‍💻 [Freelancer](https://www.freelancer.com/u/shahinrahimi)
- 🕊️ [X](https://x.com/cushydigit)

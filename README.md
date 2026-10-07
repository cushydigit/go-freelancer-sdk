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

to add this SDK to your project:

```bash
go get github.com/cushydigit/freelancer-go-sdk@v1.4.0
```

**Important Notes:**

- this package imports as github.com/cushydigit/go-freelancer-sdk/freelancer in you code
- You'll need a freelancer.com OAuth access token to authenticate request
- The SDK require GO 1.24 or later (specified)

**Minimal Import Example:**

```Go
package main

import (
    "github.com/cushydigit/go-freelancer-sdk/freelancer"
)

func main() {
    // SDK initialization - token should be provided securely
    client := freelancer.NewClient(yourAccessToken)
}
```

**Environment Variable:**
The examples assume these environment variable:

| Variable | Description | Example |
| :------- | :---------- | :------ |
| FREELANCER_ACCESS_TOKEN | OAuth2 access token | frl_xxxxxxxxxxxx |
| (Optional) PROXY_ADDR | SOCKS proxy address | socks5://localhost:1080 |

### Quick Start

Get your Freelancer.com access token from the [Freelancer Developer Portal](https://accounts.freelancer.com/), then run:

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/cushydigit/go-freelancer-sdk/freelancer"
    rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

func main() {
    // Step 1: Create client with your access token
    client := freelancer.NewClient(os.Getenv("FREELANCER_ACCESS_TOKEN"))
    
    // Step 2: Configure search options
    opts := rr.SearchActiveProjectsOptions{
        Limit:     rr.Int(5),           // Number of results per page
        Query:     rr.String("golang"), // Search keyword
        FullDescription: rr.Bool(true), // Include full project description
    }
    
    // Step 3: Fetch projects with context for timeout control
    ctx := context.Background()
    res, meta, err := client.Services.Projects.SearchActive(ctx, &opts)
    if err != nil {
        log.Fatalf("Search failed: %v", err)
    }
    
    // Step 4: Process results
    fmt.Printf("Found %d active projects\n", len(res.Result.Projects))
    for _, p := range res.Result.Projects {
        fmt.Printf("- Project %d: %s (URL: %s)\n", 
            p.ID, p.Title, p.GetFullUrl())
    }
    
    // Step 5: Access rate limit info (v1.4.0 feature)
    if meta.RateLimit.Remaining > 0 {
        fmt.Printf("Rate limit remaining: %d requests\n", meta.RateLimit.Remaining)
    }
}
```

### Authentication

This SDK authenticates with Freelancer.com using OAuth2 access tokens. The token is sent via the `freelancer-oauth-v1` header on every request.

**Creating a Client:**

```go
import (
    "github.com/cushydigit/go-freelancer-sdk/freelancer"
)

// Basic client creation
client := freelancer.NewClient("your_access_token_here")

// With custom HTTP client (for timeouts, proxies, etc.)
httpClient := &http.Client{
    Timeout: 30 * time.Second,
}
client = freelancer.NewClient(
    "your_access_token",
    freelancer.WithHttpClient(httpClient),
)

// Use sandbox environment for testing
client = freelancer.NewClient(
    "your_access_token",
    freelancer.WithSandBox(), // Uses api-sandbox.freelancer.com
)
```

**Important Security Notes:**

- **Never hard-code tokens in source code** - use environment variables
- **Rotate credential regularly** - if your token expired or compromised, revoke and generate a new one.
- **Use the sandbox for development** - `WithSandBox()` uses the test API endpoint
- **Set appropriate timeouts** - The default 30-second timeout may not be suitable for all use cases

## Error Handling

All API requests return three values: `(result, meta, error)`.

### Basic Error Checking

```go
res, meta, err := client.Services.Projects.SearchActive(ctx, opts)
if err != nil {
    // Handle various error types
    
    // Check for API-specific errors
    if apiErr, isAPIError := freelancer.IsAPIError(err); isAPIError {
        fmt.Printf("API Error: Code=%s, Message=%s\n", 
            apiErr.LegacyErrorCode, apiErr.Message)
        
        // Access detailed error info
        if apiErr.InnerError.Detail != "" {
            fmt.Println(apiErr.InnerError.Detail)
        }
    } else {
        // Network or other errors
        fmt.Printf("Request failed: %v", err)
    }
} else {
    // Process successful response
}
```

### API Error Structure

| Field | Description |
| :---- | :---------- |
| `StatusCode` | HTTP status code (e.g, 429 for rate limited) |
| `Status` | API status text (success, error, etc.) |
| `Message` | User-facing error message |
| `RequestID` | Freelancer's request ID fro support tickets |
| `InnerError.Code` | Specified API error code |
| `InnerError.Details` | Detailed explanation of the error |
| `RawPayload` | Raw response bytes (for custom handling) |
| `Meta` | Associated metadata including rate limit info |

### Rate Limit Errors

THe SDK no longer enforces client-side rate limits (v1.4.0 change)

```Go
// Check if you've been rate limited
if meta.RateLimit.Remaining == 0 {
    fmt.Printf("Rate limit reached. Current window: %v\n", 
        meta.RateLimit.Limits)
    
    // Calculate wait time based on API response headers
}
```

**Rate Limit Headers Explained:**

- `Remaining`: Requests remaining in current window
- `Limits`: Array of quote windows
- `RawRemaining`: Raw header value as string
- `RawLimit`: Raw limit info from headers

### Context Cancellation Errors

```Go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

res, meta, err := client.Services.Projects.SearchActive(ctx, opts)
if err != nil {
    if errors.Is(err, ctx.Err()) {
        fmt.Println("Request timed out or was cancelled")
    } else {
        // Other errors
        log.Printf("API error: %v", err)
    }
}
```

## Project Structure

This SDK follows a modular service design. All core logic is located in `freelancer`.

- **`client.go`**: It holds core logic
- **`types.go`**: Shared data structures
- **`responses`**: The wrappers for API replies
- **`enums.go`**: Custom types and constants for statuses, roles, and types
- **`resources.go`**: The entry point for all services.
- **`resources_*.go`**: Each file encapsulates logic for a specific API domain

## Resources (Services)

This SDK wraps the Freelancer API through typed resource services. Every public method accepts `context.Context` for cancellation and timeout control, returns `(*RawResponse, *ResponseMeta, error)` (or a typed response where the API defines one), and delegates to a shared internal `execute` pipeline handling authentication, base-URL resolution, and JSON round-tripping.

| Resource (Service) | Endpoint |
| :----------------- | :------- |
| `Projects` | `/projects/0.1/projects` |
| `Collaborations` | `/projects/0.1/projects/collaboration` |
| `Services` | `/projects/0.1/services` |
| `Reviews` | `/projects/0.1/reviews` |
| `Bids` | `/projects/0.1/bids` |
| `Milestones` | `/projects/0.1/milestones` |
| `ExpertGuarantees` | `/projects/0.1/expert_guarantees/` |
| `Users` | `/users/0.1/users/` |
| `Profiles` | `/users/0.1/profiles` |
| `Self` | various endpoints |
| `Common` | various endpoints most static resource |

### Projects — `freelancer/resources_projects.go`

Projects is the largest service: full project lifecycle (create, update, close, end, delete), search, upgrade fees, and per-project views of bids, milestones, contract info, and freelancer invitations.

| Method | Endpoint | Description |
|---|---|---|
| `Create` | `POST /projects/0.1/projects` | Creates a new project |
| `SignNDA` | `PUT /projects/0.1/projects/{project_id}` | Signs a project NDA with identity details |
| `Upgrade` | `PUT /projects/0.1/projects/{project_id}` | Applies selected upgrades to a project |
| `Update` | `PUT /projects/0.1/projects/{project_id}` | Updates project description and skills |
| `Close` | `PUT /projects/0.1/projects/{project_id}` | Closes an open project to new bids |
| `End` | `PUT /projects/0.1/projects/{project_id}` | Ends an awarded/accepted project, cancelling the contract |
| `List` | `GET /projects/0.1/projects` | Returns the user's projects (newest first) |
| `Get` | `GET /projects/0.1/projects/{project_id}` | Returns a specific project (with user projection options) |
| `SearchActive` | `GET /projects/0.1/projects/active` | Searches active projects |
| `SearchAll` | `GET /projects/0.1/projects/all` | Searches all projects |
| `InviteFreelancer` | `POST /projects/0.1/projects/{project_id}/invite` | Invites specific freelancers to bid |
| `ListUpgradeFees` | `GET /projects/0.1/projects/fees` | Returns upgrade fees per currency and free-upgrade eligibility |
| `ListBids` | `GET /projects/0.1/projects/{project_id}/bids` | Returns bids for a project (expensive — avoid full user projections) |
| `GetBidInfo` | `GET /projects/0.1/projects/{project_id}/bids_info` | Returns information needed to post a bid |
| `ListMilestones` | `GET /projects/0.1/projects/{project_id}/milestones` | Returns milestones on a project |
| `ListMilestoneRequests` | `GET /projects/0.1/projects/{project_id}/milestone_requests` | Returns milestone requests on a project |
| `GetHourlyContractInfo` | `GET /projects/0.1/hourly_contract_info` | Fetches the hourly contract matching the query |
| `GetIPContractInfo` | `GET /projects/0.1/projects/{project_id}/ip_contract_info` | Fetches IP contract info for a project |
| `Delete` | `DELETE /projects/0.1/projects/{project_id}` | Deletes a project (pending/rejected states only) |

### Collaborations — `freelancer/resources_collaborations.go`

Collaborations manage shared access to projects and control the permissions collaborators hold (chat and bid award).

| Method | Endpoint | Description |
|---|---|---|
| `List` | `GET /projects/0.1/projects/{project_id}/collaborations` | Returns project collaborations for a project |
| `Create` | `POST /projects/0.1/projects/{project_id}/collaborations` | Creates a new project collaboration |
| `Revoke` | `PUT /projects/0.1/projects/{project_id}/collaborations/{collaboration_id}/actions` | Revokes a collaboration with permission overrides |
| `UpdatePermissions` | `PUT /projects/0.1/projects/{project_id}/collaborations/{collaboration_id}/actions` | Changes chat and bid award permissions on a collaboration |
| `ListAll` | `GET /projects/0.1/projects/collaborations` | Returns all collaborations for the authenticated user |

### Bids — `freelancer/resources_bids.go`

Bids manage the full bid lifecycle on a project: creation, award/acceptance actions, time tracking, edit requests, and ratings.

| Method | Endpoint | Description |
|---|---|---|
| `List` | `GET /projects/0.1/bids` | Returns bids matching the specified criteria |
| `Get` | `GET /projects/0.1/bids/{bid_id}` | Returns a specific bid |
| `Create` | `POST /projects/0.1/bids` | Places a bid on a project |
| `Update` | `PUT /projects/0.1/bids/{bid_id}` | Updates description, amount, or milestone percentage of a bid |
| `Accept` | `PUT /projects/0.1/bids/{bid_id}` | Bid owner accepts the award |
| `Deny` | `PUT /projects/0.1/bids/{bid_id}` | Bid owner declines the award |
| `Retract` | `PUT /projects/0.1/bids/{bid_id}` | Bid owner retracts a bid before it is awarded |
| `Highlight` | `PUT /projects/0.1/bids/{bid_id}` | Bid owner highlights a bid |
| `Sponsor` | `PUT /projects/0.1/bids/{bid_id}` | Bid owner sponsors a bid |
| `Award` | `PUT /projects/0.1/bids/{bid_id}` | Project owner awards a bid to a freelancer |
| `Revoke` | `PUT /projects/0.1/bids/{bid_id}` | Project owner revokes an awarded bid |
| `Shortlist` | `PUT /projects/0.1/bids/{bid_id}` | Project owner adds a bid to the shortlist |
| `Unshortlist` | `PUT /projects/0.1/bids/{bid_id}` | Project owner removes a bid from the shortlist |
| `Hide` | `PUT /projects/0.1/bids/{bid_id}` | Project owner hides a bid |
| `Unhide` | `PUT /projects/0.1/bids/{bid_id}` | Project owner unhides a bid |
| `RequestLocationSharing` | `PUT /projects/0.1/bids/{bid_id}` | Project owner requests location sharing from the freelancer |
| `GetTimeTracking` | `GET /projects/0.1/bids/{bid_id}/time_tracking` | Returns aggregate time tracking data for a bid |
| `CreateTimeTracking` | `POST /projects/0.1/bids/{bid_id}/time_tracking` | Creates a time tracking session for a bid |
| `ListEditRequests` | `GET /projects/0.1/bids/{bid_id}/edit_requests` | Returns edit requests for a bid |
| `CreateEditRequest` | `POST /projects/0.1/bids/edit_requests` | Creates a bid edit request (accepted/awarded bids only) |
| `AcceptEditRequest` | `PUT /projects/0.1/bids/{bid_id}/edit_requests/{edit_request_id}` | Employer accepts the proposed amount and period |
| `DeclineEditRequest` | `PUT /projects/0.1/bids/{bid_id}/edit_requests/{edit_request_id}` | Employer declines the proposed amount and period |
| `GetRating` | `GET /projects/0.1/bids/{bid_id}/bid_ratings` | Fetches the rating of a single bid |
| `ListRatings` | `GET /projects/0.1/bid_ratings` | Fetches ratings for multiple bids |
| `CreateRating` | `POST /projects/0.1/bids/{bid_id}/bid_ratings` | Rates a bid |
| `UpdateRating` | `PUT /projects/0.1/bids/{bid_id}/bid_ratings/{bid_rating_id}` | Updates an existing bid rating |

### Common — `freelancer/resources_common.go`

Common exposes platform-wide reference data and job discovery: countries, timezones, currencies, categories, budgets, jobs, and job bundles.

| Method | Endpoint | Description |
|---|---|---|
| `ListCountries` | `GET /common/0.1/countries` | Lists countries |
| `ListTimezones` | `GET /common/0.1/timezones` | Lists timezones |
| `ListCurrencies` | `GET /projects/0.1/currencies` | Lists currencies |
| `ListCategories` | `GET /projects/0.1/categories` | Lists categories (optionally with jobs per category) |
| `ListBudgets` | `GET /projects/0.1/budgets` | Lists budgets by currency |
| `ListJobs` | `GET /projects/0.1/jobs` | Lists jobs |
| `SearchJobs` | `GET /projects/0.1/jobs/search` | Searches jobs (substring match across all parameters) |
| `ListJobBundles` | `GET /projects/0.1/job_bundles` | Lists job bundles |
| `ListJobBundleCategories` | `GET /projects/0.1/job_bundle_categories` | Lists job bundle categories |

### Expert Guarantees — `freelancer/resources_expertguarantees.go`

ExpertGuarantees manages expert guarantee listings and their release flow between the guarantee creator and the project owner.

| Method | Endpoint | Description |
|---|---|---|
| `List` | `GET /projects/0.1/expert_guarantees` | Returns expert guarantees |
| `Release` | `PUT /projects/0.1/expert_guarantees/{expert_guarantee_id}` | Project owner releases an expert guarantee |
| `RequestRelease` | `PUT /projects/0.1/expert_guarantees/{expert_guarantee_id}` | Guarantee creator requests a release |

### Milestones — `freelancer/resources_milestones.go`

Milestones covers both milestone CRUD with release/cancel actions and the separate milestone request lifecycle (request, accept, reject, delete).

| Method | Endpoint | Description |
|---|---|---|
| `List` | `GET /projects/0.1/milestones` | Returns milestones (excludes un-awarded prepaid) |
| `Get` | `GET /projects/0.1/milestones/{milestone_id}` | Returns a specific milestone |
| `Create` | `POST /projects/0.1/milestones` | Creates a milestone |
| `Release` | `PUT /projects/0.1/milestones/{milestone_id}` | Project owner releases a milestone payment |
| `Update` | `PUT /projects/0.1/milestones/{milestone_id}` | Project owner updates the milestone description |
| `RequestCancel` | `PUT /projects/0.1/milestones/{milestone_id}` | Project owner requests milestone cancellation |
| `RequestRelease` | `PUT /projects/0.1/milestones/{milestone_id}` | Freelancer requests milestone release |
| `Cancel` | `PUT /projects/0.1/milestones/{milestone_id}` | Freelancer cancels a milestone |
| `RejectCancel` | `PUT /projects/0.1/milestones/{milestone_id}` | Freelancer rejects a cancellation request |
| `ListRequests` | `GET /projects/0.1/milestone_requests` | Returns milestone requests |
| `GetRequest` | `GET /projects/0.1/milestone_requests/{milestone_request_id}` | Returns a specific milestone request |
| `CreateRequest` | `POST /projects/0.1/milestone_requests` | Creates a milestone request |
| `AcceptRequest` | `PUT /projects/0.1/milestone_requests/{milestone_request_id}` | Project owner accepts a milestone request |
| `RejectRequest` | `PUT /projects/0.1/milestone_requests/{milestone_request_id}` | Project owner rejects a milestone request |
| `DeleteRequest` | `PUT /projects/0.1/milestone_requests/{milestone_request_id}` | Bid owner deletes a milestone request |

### Profiles — `freelancer/resources_profiles.go`

Profiles provides CRUD for the authenticated user's public profile.

| Method | Endpoint | Description |
|---|---|---|
| `Create` | `POST /users/0.1/profiles` | Creates a new profile (returns the created profile) |
| `Get` | `GET /users/0.1/profiles` | Gets the user's profile(s) |
| `Update` | `PUT /users/0.1/profiles` | Updates the user's profile |

### Reviews — `freelancer/resources_reviews.go`

Reviews handles posting reviews (per role) and featuring/unfeaturing them by review type (project or contest).

| Method | Endpoint | Description |
|---|---|---|
| `List` | `GET /projects/0.1/reviews` | Returns project reviews |
| `CreateForFreelancer` | `POST /projects/0.1/reviews` | Posts a review as a freelancer |
| `CreateForEmployer` | `POST /projects/0.1/reviews` | Posts a review as an employer |
| `FeatureProject` | `PUT /projects/0.1/reviews/{review_id}` | Features a project review |
| `UnfeatureProject` | `PUT /projects/0.1/reviews/{review_id}` | Unfeatures a project review |
| `FeatureContest` | `PUT /projects/0.1/reviews/{review_id}` | Features a contest review |
| `UnfeatureContest` | `PUT /projects/0.1/reviews/{review_id}` | Unfeatures a contest review |

### Self — `freelancer/resources_self.go`

Self covers the authenticated user's own account state: identity, devices, job preferences, pools, and their own projects/contests.

| Method | Endpoint | Description |
|---|---|---|
| `Get` | `GET /users/0.1/self` | Returns information about the current user |
| `ListDevices` | `GET /users/0.1/self/devices` | Returns recently logged-in devices |
| `AddJobs` | `POST /users/0.1/self/jobs` | Adds jobs to the user's job list |
| `UpdateJobs` | `PUT /users/0.1/self/jobs` | Replaces the user's job list |
| `DeleteJobs` | `DELETE /users/0.1/self/jobs` | Removes jobs from the user's job list |
| `ListPools` | `GET /users/0.1/pools` | Returns pools belonging to the current user |
| `ListProjects` | `GET /projects/0.1/self` | Returns projects/contests the user created or participated in |

### Services — `freelancer/resources_services.go`

Services covers the platform's orderable services (list, search, and ordering).

| Method | Endpoint | Description |
|---|---|---|
| `Order` | `POST /projects/0.1/services/{service_type}/{service_id}/order` | Orders a service |
| `List` | `GET /projects/0.1/services` | Returns services |
| `SearchActive` | `GET /projects/0.1/services/active` | Returns active services |

### Users — `freelancer/resources_users.go`

Users covers the public user directory: users, freelancer search, reputations, enterprises, portfolios, and violation reports.

| Method | Endpoint | Description |
|---|---|---|
| `List` | `GET /users/0.1/users` | Returns a list of users |
| `Get` | `GET /users/0.1/users/{user_id}` | Returns a specific user |
| `SearchFreelancer` | `GET /users/0.1/users/directory` | Returns a paginated list of eligible freelancers |
| `ListReputations` | `GET /users/0.1/reputations` | Returns reputations for a list of users |
| `ListEnterprises` | `GET /users/0.1/enterprises` | Returns a list of enterprises |
| `ListPortfolios` | `GET /users/0.1/portfolios` | Returns portfolios for a list of users |
| `CreateViolationReport` | `POST /users/0.1/violation_reports` | Creates a user violation report |

## Use-Cases

### List Countries

```Go
opts := rr.ListCountriesOptions{
    ExtraDetails: rr.Bool(true), // Include extra details
}

res, _, err := c.Services.Common.ListCountries(ctx, &opts)
// nil slices has length of zero
if err == nil && len(res.Result.Countries) > 0 {
    for _, c := range res.Result.Countries {
        fmt.Printf("Name: %s, Code: %s PhoneCode: %.0f\n", c.Name, c.Code, c.PhoneCode)
    }
    fmt.Printf("Fetched %d countries\n", len(res.Result.Countries))
}
```

### Listing Timezones

```Go
opts := rr.ListTimezonesOptions{
    TimezoneNames: []string{
        "America/New_York",
        "Europe/London",
        "Asia/Tokyo",
    },
}

res, _, err := c.Services.Common.ListTimezones(ctx, &opts)
// nil slices has length of zero
if len(res.Result.Timezones) > 0 && err == nil {
    for _, tz := range res.Result.Timezones {
        fmt.Printf("%d: country=%s (UTC+%.1f)\n", tz.ID, tz.Country, tz.Offset)
    }
    fmt.Printf("Fetched %d timezone(s)\n", len(res.Result.Timezones))
}
```

### Listing Currencies

```Go
res, _, err := c.Services.Projects.Currencies.List(ctx, nil)
// nil slices has length of zero
if len(res.Result.Currencies) > 0 && err == nil {
    for _, cur := range res.Result.Currencies {
        // Each currency record supports exchange rates for calculations
        fmt.Printf("%s: %.3f USD rate, Country=%s\n", cur.Code, cur.ExchangeRate, cur.Country)
    }
    fmt.Printf("Fetched %d countries\n", len(res.Result.Currencies))
}
```

### Searching Active Projects

```go
from := time.Now().Add(-24 * time.Hour) // 

// Basic search with filters
opts := rr.SearchActiveProjectsOptions{
    Query:  rr.String("Go developer"),
    Limit:  rr.Int(20), // Results per page
    Offset: rr.Int(0),

    // Filtering options
    FromTime:    &from, // Filter projects within 24 hours
    UserDetails: rr.Bool(true), // Include user info
    // Sort
    SortField:   rr.Enum(rr.SortFieldsTimeUpdated),
    ReverseSort: rr.Bool(false), // Newest first
}
// Execute search
res, _, err := c.Services.Projects.SearchActive(ctx, &opts)
if err == nil && len(res.Result.Projects) > 0 {
    for _, p := range res.Result.Projects {
        budgetString := fmt.Sprintf(
            "[%s%1.f - %s%1.f]",
            p.Currency.Sign,
            p.Budget.Minimum,
            p.Currency.Sign,
            p.Budget.Minimum,
        )
        fmt.Printf("\n-%d: %s %s\n", p.ID, budgetString, p.Title)
        // access unix time easily
        time1, time2, time3 := p.SubmitDateAt(), p.UpdatedAt(), p.SubmittedAt()
        if time1 != nil && time2 != nil && time3 != nil {
            fmt.Printf(
                "SubmitDate: %s\tTimeUpdated: %s\tTimeSubmitted: %s\n",
                time1.Format("Jan 2, 2006 at 15:04"),
                time2.Format("Jan 2, 2006 at 15:04"),
                time3.Format("Jan 2, 2006 at 15:04"),
            )
        }
        if p.SubmitDateAt() != nil && p.UpdatedAt() != nil && p.SubmittedAt() != nil {

        }
    }
    fmt.Printf("Showing %d of %d total projects\n", len(res.Result.Projects), res.Result.TotalCount)
}
```

### Searching with Geographic Bounds

Use lat/lng coordinates for location-based filtering:

```Go
opts := rr.SearchActiveProjectsOptions{
    Query:     rr.String("website"),
    
    // Geographic bounding box
    Latitude:          rr.Float64(40.7128),
    Longitude:         rr.Float64(-74.0060),  // New York
    TopRightLatitude:  rr.Float64(41.0),
    TopRightLongitude: rr.Float64(-73.5),
    BottomLeftLatitude: rr.Float64(40.4),
    BottomLeftLongitude: rr.Float64(-74.5),
}

res, _, err := client.Services.Projects.SearchActive(ctx, &opts)
```

### Creating a Project

```Go
body := reqres.CreateProjectBody{
    Title:       "Go Developer Needed",
    Description: "Need an experienced Go developer...",

    Budget: rr.Budget{
        Minimum:    100.0,
        Maximum:    250.0,
        CurrencyID: 1,
    },

    Type: rr.Enum(rr.ProjectBudgetFixed),
}

res, _, err := c.Services.Projects.Create(ctx, body)
if err != nil {
    fmt.Println(err)
    return
}
```

### Working with Bids on a Project

```Go
// List bids for a project
bidsOpts := rr.ListBidsOptions{
    Limit:       rr.Int(10),
    Reputation:  rr.Bool(true),   // Include reputation data
    UserDetails: rr.Bool(false),  // Don't load full user info yet
}

res, _, err := client.Services.Projects.ListBids(ctx, projectId, &bidsOpts)
```

### Milestones and Milestone Request

```Go
// List milestones
mileOpts := rr.ListMilestonesOptions{
    Projects: []int64{projectId},
    Limit:    rr.Int(25),
}
res, _, err := client.Services.Projects.Milestones.List(ctx, &mileOpts)

// Create a milestone request
body := reqres.CreateMilestoneRequestBody{
    ProjectID:   projectId,
    BidID:       bidId,          // Required: the bid to award this milestone on
    Amount:      2500,           // Milestone amount
}
res, _, err := client.Services.Projects.MilestoneRequests.Create(ctx, body)

// Award (release) a milestone request
action := rr.MilestoneActionRequestRelease()
apiBody := reqres.ActionMilestoneRequestBody{
    Action: rr.Enum(action),
}
res, _, err := client.Services.Projects.MilestoneRequests.Action(ctx, requestId, &apiBody)
```

### Directory Search - Finding Freelancers

```Go
// Find freelancers matching criteria
searchOpts := rr.SearchFreelancerOptions{
    Query:     rr.String("golang"),
    
    // Price and skill filters
    HourlyRateMin: rr.Int(30),
    ReviewCountMin:  rr.Int(5),        // At least 5 reviews
    
    // Pagination
    Limit:  rr.Int(20),
    Offset: rr.Int(0),
    
    // Include detailed data
    Reputation:        rr.Bool(true),
    ProfileDescription: rr.Bool(true),
}

res, meta, err := client.Services.Users.SearchFreelancer(ctx, &searchOpts)
if err != nil {
    log.Printf("Error: %v", err)
    return nil
}

// Process results
freelancers := res.Result.Users
for _, u := range freelancers {
    fmt.Printf("- %s: $%.2f/hr, Reviews: %d\n", 
        u.DisplayName, u.HourlyRate, u.Reputation.Entrity)
}

fmt.Printf("Found %d freelancers out of %d total\n", 
    len(freelancers), res.Result.TotalCount)
```

### Getting User Details

```Go
// Direct user lookup by ID
userData, _, err := client.Services.Users.Get(ctx, userId)
if err != nil {
    log.Printf("Error: %v", err)
    return
}
```

### Working with User Profile

```Go
// Create a profile for yourself or another user
createOpts := rr.CreateProfileBody{
    Tagline:     "Experienced Go developer",
    HourlyRate:  75,
    Description: "Full-stack software engineer specializing in Go and Rust.",
}

res, _, err := client.Services.Users.Profiles.Create(ctx, createOpts)
if err != nil {
    log.Printf("Error: %v", err)
    return
}

// Update an existing profile
updateOpts := rr.UpdateProfileBody{
    HourlyRate:  85,
    Description: "Senior Go engineer with 8 years experience...",
}

res, _, err = client.Services.Users.Profiles.Update(ctx, updateOpts)
```

### Device Management

```Go
// List devices used by current user (authenticated endpoint)
devices, _, err := client.Services.UsersSelfDevices.List(ctx)
if err != nil {
    log.Printf("Error: %v", err)
}

for _, device := range devices.Result.Devices {
    fmt.Printf("- %s in %s (%d)\n", 
        device.City, device.Country, device.LastLogin)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

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
- [x] **Unit Testing:** (77.5% coverage)
- [ ] **Use Case:**

### Upcoming Features

- [ ] **Messaging:** Threads and direct message handling.
- [ ] **Contests:** Browsing and participating in contests.

## Changelog

For a detailed list of changes, please see the [CHANGELOG.md](CHANGELOG.md).

## Disclaimer

This is an unofficial library and is not affiliated with, endorsed by, or associated with Freelancer.com.
Please ensure you comply with the [Freelancer API Terms](https://www.freelancer.com/about/terms) and Conditions when using this software.

For details on the underlying API endpoints and parameters, refer to the official [Freelancer.com API Documentation](https://developers.freelancer.com/).

## Contact

Feel free to reach out if you have question or suggestions:

- 📧 [sh.rahimi.dev@gmail.com](mailto:sh.rahimi.dev@gmail.com)
- 💼 [LinkedIn](https://www.linkedin.com/in/shahin-rahimi-828447254/)
- 🧑‍💻 [Freelancer](https://www.freelancer.com/u/shahinrahimi)
- 🕊️ [X](https://x.com/cushydigit)

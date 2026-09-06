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

### Installation

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
| Variable | Description | Example value |
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

### Error Handling

All API requests return three values: `(result, meta, error)`.

#### Basic Error Checking

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

#### API Error Structure

| Field | Description |
| StatusCode | HTTP status code (e.g, 429 for rate limited) |
| Status | API status text (OK, ERROR, etc.) |
| Message | User-facing error message |
| RequestID | Freelancer's request ID fro support tickets |
| InnerError.Code | Specified API error code |
| InnerError.Details | Detailed explanation of the error |
| RawPayload | Raw response bytes (for custom handling) |
| Meta | Associated metadata including rate limit info |

#### Rate Limit Errors

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

#### Context Cancellation Errors

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
- **`services.go`**: The entry point for all services.
- **`service_*.go`**: Each file encapsulates logic for a specific API domain

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
- [x] **Unit Testing:** (32% coverage)
- [ ] **Use Case:**

### Upcoming Features

- [ ] **Messaging:** Threads and direct message handling.
- [ ] **Contests:** Browsing and participating in contests.

## Changelog

For a detailed list of changes, please see the [CHANGELOG.md](CHANGELOG.md).

## Disclaimer

This is an unofficial library and is not affiliated with, endorsed by, or associated with Freelancer.com.
Please ensure you comply with the [Freelancer API Terms](https://www.freelancer.com/about/terms) and Conditions when using this software.

## Contact

Feel free to reach out if you have question or suggestions:

- 📧 [sh.rahimi.dev@gmail.com](mailto:sh.rahimi.dev@gmail.com)
- 💼 [LinkedIn](https://www.linkedin.com/in/shahin-rahimi-828447254/)
- 🧑‍💻 [Freelancer](https://www.freelancer.com/u/shahinrahimi)
- 🕊️ [X](https://x.com/cushydigit)

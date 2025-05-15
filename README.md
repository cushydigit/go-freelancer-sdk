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
- 📧 sh.rahimi.dev@gmail.com
- [![LinkedIn](https://upload.wikimedia.org/wikipedia/commons/0/01/LinkedIn_Logo_2013.svg)](https://www.linkedin.com/in/shahinrahimi)
- [![Twitter](https://upload.wikimedia.org/wikipedia/commons/6/60/Twitter_Logo_2021.svg)](https://twitter.com/cushydigit)
- [![GitHub](https://upload.wikimedia.org/wikipedia/commons/9/91/Octicons-mark-github.svg)](https://github.com/cushydigit)






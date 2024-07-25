# URLShortner
A simple URL Shortner made in GO and GORM.

## Prerequisites
- Go

``` 
brew install go
```

### Installation

```
git clone https://github.com/ashwinexe/urlshortner
```
### Build the project

```
go build
```

### Execute the project

```
go run main.go
```

### How to use
1. To shorten a URL, make a POST request to your localhost, (default `8080`) with response URL. You can use curl or API tool like postman
```
curl -X POST -d "url=https://google.com" https://localhost:<port>/shorten
```
2. Response url will be in JSON object
```
{"shortened_url: "https://localhost:<port>/url"}
```
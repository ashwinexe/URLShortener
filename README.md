# URLShortener
A simple URL Shortener made in GO and GORM.

Try it out in ![https://hub.docker.com/r/ashwinexe/urlshortener](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white): 

```
docker pull ashwinexe/urlshortener

docker run -p 8080:8080 ashwinexe/urlshortener
```

## How to shorten a URL

> Note: If the URLs are not being stored properly or top-domains are not being updated, you may have an existing db file, delete it and let the program create a fresh one.

### Using command line

To shorten a link make a **POST** reqeuest to `/shorten`
```
curl -X POST -d "url=http://youtube.com" http://localhost:8080/shorten
```

To check top 3 domains shortened make a **GET** request to `/topdomains`
```
curl http://localhost:8080/topdomains
```

### Using API testing tools
#### [Hoppscotch](https://hoppscotch.io/) or [Postman](https://www.postman.com/)
1. Create a new REST API tab.
2. Set it request to `GET`
3. Paste the URL `http://localhost:8080/shorten`
4. Set **Body** of request to **Content type:** *application/x-www-form-urlencoded*
5. Put the *parameter* as `url` and *value* as the original link you want to shorten. (example: `http://github.com/ashwinexe`)

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

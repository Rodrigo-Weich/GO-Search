<p align="center">
  <img src="https://raw.githubusercontent.com/Rodrigo-Weich/GO-Search/main/content/logo.png">
</p>

<p align="center">
  <a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Golang-1.15-blue.svg">
  </a>
  <a href="https://github.com/Rodrigo-Weich/GO-Search">
    <img src="https://img.shields.io/badge/Release-1.X-red.svg">
  </a>
    <a href="https://opensource.org">
    <img src="https://img.shields.io/badge/Open%20Source-%E2%9D%A4-brightgreen.svg">
  </a>
</p>

<p align="center">
  GO Search is a script that takes a parameter from the user. This parameter is used for a Google search, which returns all links from the search results, thus allowing the user to access them through the terminal, without the need to use a search interface.
</p>

## How to Install
```bash
# Install dependencies (if necessary)
$ sudo apt install golang git -y

# Get this repository
$ git clone git@github.com:Rodrigo-Weich/GO-Search.git

# Go into the repository
$ cd GO-Search/

# Run
$ go run main.go
```

## How to Use

When started, the system displays to the user to inform an argument for research, after informed, the system performs a query on google and returns the user the URLs of the results found, that way, the user can only open the URLs without having the need opening a search interface.

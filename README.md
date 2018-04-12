# Iselia

## Description

Iselia is a service for tchalo written in go.

## SLO and SLI

- Availability: 99%
- Average response time: < 10ms

## Onboarding and Development Guide

### Prerequisite

- Git
- Go 1.9 or later
- dep
- govendor

### Setup

- Install Git

  See [Git Installation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

- Install Go (Golang)

  See [Golang Installation](https://golang.org/doc/install)

- Install dep

  See [dep Installation](https://github.com/golang/dep)

- Install govendor

  See [govendor Installation](https://github.com/kardianos/govendor)
  
- Clone this repo in `$GOPATH/src/github.com/tchalo`

  If you have not set your GOPATH, set it using [this](https://golang.org/doc/code.html#GOPATH) guide.
  If you don't have directory `src`, `github.com`, or `tchalo` in your GOPATH, please make them.

  ```sh
  git clone https://github.com/tchalo/iselia.git
  ```

- Go to iselia directory and install dependency
  ```sh
  cd $GOPATH/src/github.com/tchalo/iselia
  make dep
  ```

- Copy env.sample and if necessary, modify the env value(s)

  ```sh
  cp env.sample .env
  ```

- Build binary file

    ```
    govendor build -o app/web/web app/web/main.go
    ```

- Run iselia

    ```
    cd app/web
    ./web
    ```

- Check whether it is ran correctly. It should return ```ok``` message

    ```
    curl -X GET "http://localhost:3000/healthz"
    ```

### Development

- Make new branch with descriptive name about the change(s) and checkout to the new branch

  ```sh
  git checkout -b branch-name
  ```

- Make your change(s) and make the test(s)

- Beautify and review the codes by running this command (note: **this is a must!**)

  ```sh
  make pretty
  ```
  
- Run test in your local

  ```sh
  make test
  ```

  If there are any errors after executing the command, please fix them first

- Save dependencies

  ```sh
  govendor add +external
  govendor fetch -v +outside
  ```

- Commit and push your change to upstream repository

  ```sh
  git commit -m "a meaningful commit message"
  git push origin branch-name
  ```

- Open Pull Request in Repository

- Pull request should be merged only if review phase is passed

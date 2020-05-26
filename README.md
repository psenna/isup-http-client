# isup-http-client

[![Go Report Card](https://goreportcard.com/badge/github.com/psenna/isup-http-client)](https://goreportcard.com/report/github.com/psenna/isup-http-client)
[![Build Status](https://travis-ci.org/psenna/isup-http-client.svg?branch=master)](https://travis-ci.org/psenna/isup-http-client)

Http client from isup

# testing 

go test ./... -coverprofile=cover.out

go tool cover -func=cover.out

go tool cover -html=cover.out -o cover.html

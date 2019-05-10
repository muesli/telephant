Telephant!
==========

A lightweight but modern Social Media client, written in Go & QML.

![telephant logo](/assets/telephant.png)

## Features

- [x] Live feed via Mastodon's Streaming API
- [x] Multi pane support
- [x] Linux/macOS/Windows (Android & iOS should be working, but aren't tested yet)
- [x] Media previews
- [x] Shortened URL resolving
- [ ] System notifications
- [ ] Multiple accounts (work-in-progress)
- [ ] Support for more networks

## Installation

Make sure you have a working Go environment (Go 1.8 or higher is required).
See the [install instructions](http://golang.org/doc/install.html).

### Dependencies

Before you can build Telephant you need to install the [Go/Qt bindings](https://github.com/therecipe/qt/wiki/Installation#regular-installation).

    go get -u -v github.com/therecipe/qt/cmd/...

### Building Telephant

    mkdir -p $(go env GOPATH)/src/github.com/muesli
    cd $(go env GOPATH)/src/github.com/muesli
    git clone https://github.com/muesli/telephant.git

    cd telephant
    go get -u -v
    $(go env GOPATH)/bin/qtdeploy build desktop .

### Run it

    ./deploy/linux/telephant

![telephant Screenshot](/assets/screenshot.png)

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/telephant)
[![Build Status](https://travis-ci.org/muesli/telephant.svg?branch=master)](https://travis-ci.org/muesli/telephant)
[![Go ReportCard](http://goreportcard.com/badge/muesli/telephant)](http://goreportcard.com/report/muesli/telephant)

Telephant!
======

A lightweight but modern Social Media client, written in Go & QML.

![telephant Screenshot](/assets/screenshot.png)

## Features

- [x] Live feed via Mastodon's Streaming API
- [x] Live feed via Twitter's Streaming API
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

### Building Telephant!

    git clone https://github.com/muesli/telephant.git
    qtdeploy build desktop telephant/

### Run it

    ./telephant/deploy/linux/telephant

![telephant logo](/assets/telephant.png)

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/telephant)
[![Build Status](https://travis-ci.org/muesli/telephant.svg?branch=master)](https://travis-ci.org/muesli/telephant)
[![Go ReportCard](http://goreportcard.com/badge/muesli/telephant)](http://goreportcard.com/report/muesli/telephant)

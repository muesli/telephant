Chirp!
======

A lightweight but modern Twitter client, written in Go & QML.

## Features

- [x] Live feed via Twitter's Streaming API
- [x] Multi pane support
- [ ] Multiple Twitter accounts (work-in-progress)
- [ ] Media previews
- [ ] System notifications

## Installation

Make sure you have a working Go environment. See the [install instructions](http://golang.org/doc/install.html).

### Dependencies

Before you can build Chirp you need to install the [Go/Qt bindings](https://github.com/therecipe/qt/wiki/Installation#regular-installation).

### Building Chirp!

    git clone https://github.com/muesli/chirp.git
    qtdeploy build desktop chirp/

### Run it

    ./chirp/deploy/linux/chirp

## Development

API docs can be found [here](http://godoc.org/github.com/muesli/chirp).

[![Build Status](https://secure.travis-ci.org/muesli/chirp.png)](http://travis-ci.org/muesli/chirp)
[![Go ReportCard](http://goreportcard.com/badge/muesli/chirp)](http://goreportcard.com/report/muesli/chirp)

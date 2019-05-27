Telephant!
==========

A lightweight but modern Mastodon client, written in Go & QML.

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

### Packages & Installers

- Arch Linux: [telephant-git](https://aur.archlinux.org/packages/telephant-git/)
- [Linux Static 64bit](https://github.com/muesli/telephant/releases/download/v0.1/telephant_0.1pre_Linux_64bit)
- [Windows 64bit](https://github.com/muesli/telephant/releases/download/v0.1/telephant_0.1pre_Windows_64bit.exe)

### From Source

Make sure you have a working Go environment (Go 1.9 or higher is required).
See the [install instructions](http://golang.org/doc/install.html).

You will also need Qt5 and its development headers installed.

#### Dependencies

Before you can build Telephant you need to install the [Go/Qt bindings](https://github.com/therecipe/qt/wiki/Installation#regular-installation).

#### Qt5 dependencies (Ubuntu example)

    apt-get --no-install-recommends install build-essential libglib2.0-dev libglu1-mesa-dev libpulse-dev
    apt-get --no-install-recommends install libqt*5-dev qt*5-dev qt*5-doc-html qml-module-qtquick*

    $(go env GOPATH)/bin/qtsetup -test=false

#### Building Telephant

    export QT_PKG_CONFIG=true
    go get -u -v -tags=no_env github.com/therecipe/qt/cmd/...
    go get -d -u -v github.com/muesli/telephant
    cd $(go env GOPATH)/src/github.com/muesli/telephant
    $(go env GOPATH)/bin/qtdeploy build desktop .

#### Within a Docker container

Follow the build instructions above, but instead of the last command, run:

    $(go env GOPATH)/bin/qtdeploy -docker build linux

#### Run it

    ./deploy/linux/telephant

![telephant Screenshot](/assets/screenshot.png)

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/telephant)
[![Build Status](https://travis-ci.org/muesli/telephant.svg?branch=master)](https://travis-ci.org/muesli/telephant)
[![Go ReportCard](http://goreportcard.com/badge/muesli/telephant)](http://goreportcard.com/report/muesli/telephant)

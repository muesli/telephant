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
- [x] System notifications
- [ ] Multiple accounts (work-in-progress)
- [ ] Support for more networks

## Installation

### Packages & Binaries

- Arch Linux: [telephant-git](https://aur.archlinux.org/packages/telephant-git/)
- [Ubuntu 64bit Binary](https://github.com/muesli/telephant/releases/download/v0.1-rc2/telephant_0.1rc2_Ubuntu_64bit)
- [Linux Static 64bit Binary](https://github.com/muesli/telephant/releases/download/v0.1-rc2/telephant_0.1rc2_Linux_64bit)
- [Windows 64bit Binary](https://github.com/muesli/telephant/releases/download/v0.1-rc2/telephant_0.1rc2_Windows_64bit.exe)

#### Ubuntu

You need to install the following dependencies to run the Ubuntu binary:

```bash
apt install libqt5gui5 libqt5qml5 libqt5quickcontrols2-5 libqt5multimedia5-plugins \
            qml-module-qtquick2 qml-module-qtmultimedia qml-module-qtquick-layouts \
            qml-module-qtquick-controls qml-module-qtquick-controls2 \
            qml-module-qtquick-window2 qml-module-qtgraphicaleffects \
            qml-module-qtquick-dialogs qml-module-qt-labs-folderlistmodel \
            qml-module-qt-labs-settings
```

### From Source

Make sure you have a working Go environment (Go 1.9 or higher is required).
See the [install instructions](http://golang.org/doc/install.html).

You will also need Qt5 and its development headers installed.

#### Dependencies

Before you can build Telephant you need to install the [Go/Qt bindings](https://github.com/therecipe/qt/wiki/Installation#regular-installation).

#### Qt5 dependencies (Ubuntu example)

    apt-get --no-install-recommends install build-essential libglib2.0-dev libglu1-mesa-dev libpulse-dev
    apt-get --no-install-recommends install libqt*5-dev qt*5-dev qt*5-doc-html qml-module-qtquick*
    apt-get install qml-module-qtmultimedia qml-module-qt-labs-folderlistmodel qml-module-qt-labs-settings

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

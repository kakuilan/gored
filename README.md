[![Build Status](https://travis-ci.com/bitontop/gored.svg?branch=master)](https://travis-ci.com/bitontop/gored)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/bitontop/gored/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/bitontop/gored?status.svg)](https://godoc.org/github.com/bitontop/gored)
[![Coverage Status](http://codecov.io/github/bitontop/gored/coverage.svg?branch=master)](http://codecov.io/github/bitontop/gored?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/bitontop/gored)](https://goreportcard.com/report/github.com/bitontop/gored)

A Realtime-Exchange-Data SDK is supporting multiple exchanges written in Golang.

**Please note that this SDK is under heavily development and is not ready for production!**

## Community

Join our telegram to discuss all things related to GoRed! [GoRed Telegram](https://t.me/goredtalk)

## Exchange Support Table

|   Exchange   | Public API | Private API | goCAD |
|--------------|------------|-------------|-------|
| BiBox| Yes  | Yes  | NA  |
| BigONE | Yes | Yes  | NA|
| Binance| Yes  | Yes  | NA  |
| BitFinex | Yes | Yes  | NA|
| BitForex | Yes | No  | NA|
| BitMAX | Yes  | Yes  | NA  |
| BitMEX | Yes | Yes  | NA |
| BitStamp | Yes  | No  | NA  |
| Bittrex | Yes | Yes  | NA |
| BitZ | Yes | Yes  | NA |
| CoinEX | Yes | Yes  | NA |
| DargonEX | Yes | No  | NA |
| GateIO | Yes | Yes  | NA |
| HitBTC | Yes | Yes  | NA |
| Huobi PRO | Yes | Yes  | NA |
| Huobi OTC | Yes | No  | NA |
| IDEX | Yes | No  | NA |
| KuCoin | Yes | Yes  | NA |
| Liquid | Yes | Yes  | NA |
| MXC | Yes | No  | NA |
| OKEX | Yes | Yes  | NA |
| OTCBTC | Yes | Yes  | NA |
| Stex | Yes | Yes  | NA |
| TOKOK | Yes | Yes  | NA |

We are aiming to support the top 100 highest volume exchanges based off the [CoinMarketCap exchange data](https://coinmarketcap.com/exchanges/volume/24-hour/).

** NA means not applicable as the Exchange does not support the feature.

## Current Features

+ Unified all symbols / pairs into Bitontop standard.
+ Support for all Exchange fiat and digital currencies, with the ability to individually toggle them on/off.

## Planned Features

Planned features can be found on our [community Trello page](https://trello.com/gored).

## Contribution

Please feel free to submit any pull requests or suggest any desired features to be added.

When submitting a PR, please abide by our coding guidelines:

+ Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
+ Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
+ Code must adhere to our [coding style](https://github.com/bitontop/gored/blob/master/.github/CONTRIBUTING.md).
+ Pull requests need to be based on and opened against the `master` branch.

## Compiling instructions

Download and install Go from [Go Downloads](https://golang.org/dl/) for your
platform.

### Linux/OSX

gored is built using [Go Modules](https://github.com/golang/go/wiki/Modules) and requires Go 1.11 or above
Using Go Modules you now clone this repository **outside** your GOPATH

```bash
git clone https://github.com/bitontop/gored.git
cd gored
go build
mkdir ~/.gored

```

### Windows

```bash
git clone https://github.com/bitontop/gored.git
cd gored
go build

```

+ Run the `gored` binary file inside your GOPATH bin folder.

## Donations

<img src="" hspace="70">

If this framework helped you in any way, or you would like to support the developers working on it, please donate Bitcoin to:

***Bitcoin Address***
3BMEXqK2hMJx2N2j6B4b5MsJMXQ9kAC3t8

***AIB Address***
ALVNJC94McFjN7enZsY2XaR6fY3GWDY3Wh


## Binaries

Binaries will be published once the codebase reaches a stable condition.

## Contributor List

### A very special thank you to all who have contributed to this program:

|User|Github|Contribution Amount|
|--|--|--|
| chunlee1991 | https://github.com/chunlee1991 
| tony0408 | https://github.com/tony0408 
| 9cat | https://github.com/9cat 
| botemple | https://github.com/botemple 
| iobond | https://github.com/iobond 


## Who used GoRED
[CoinDataPage](https://www.coindatapage.com)

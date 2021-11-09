# GoAsterix

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Go Reference](https://pkg.go.dev/badge/github.com/mokhtarimokhtar/goasterix.svg)](https://pkg.go.dev/github.com/mokhtarimokhtar/goasterix)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

[![Build-Test Status](https://github.com/mokhtarimokhtar/goasterix/actions/workflows/ci.yml/badge.svg)](https://github.com/mokhtarimokhtar/goasterix/actions?workflow=test)
[![Coverage](https://codecov.io/gh/mokhtarimokhtar/goasterix/branch/main/graphs/badge.svg?branch=main)](https://codecov.io/gh/mokhtarimokhtar/goasterix)
[![Go Report Card](https://goreportcard.com/badge/github.com/mokhtarimokhtar/goasterix)](https://goreportcard.com/report/github.com/mokhtarimokhtar/goasterix)
[![GitHub issues](https://img.shields.io/github/issues/mokhtarimokhtar/goasterix)](https://github.com/mokhtarimokhtar/goasterix/issues)

This library provides an ASTERIX Frame(binary data) decoding/parsing(json,xml) capabilities for Go.

## ASTERIX

ASTERIX (All Purpose Structured EUROCONTROL Surveillance Information Exchange) is an application/presentation protocol
responsible for data definition and data assembly developed to support the transmission and exchange of surveillance
related data. Its purpose is to allow a meaningful transfer of information between two application entities using a
mutually agreed representation of the data to be exchanged.

More about ASTERIX: [Eurocontrol asterix](https://www.eurocontrol.int/asterix)

### Documentation

* [API Reference](https://pkg.go.dev/github.com/mokhtarimokhtar/goasterix)
* [Hex to String example](https://github.com/mokhtarimokhtar/goasterix/tree/main/examples/hextostring)
* [Decode binary file example](https://github.com/mokhtarimokhtar/goasterix/tree/main/examples/readfile)
* [Parsing Json example](https://github.com/mokhtarimokhtar/goasterix/tree/main/examples/readfiletojson)

## Installation

    go get github.com/mokhtarimokhtar/goasterix

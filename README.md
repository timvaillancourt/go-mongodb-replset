# go-mongodb-replset
[![](https://godoc.org/github.com/timvaillancourt/go-mongodb-replset?status.svg)](http://godoc.org/github.com/timvaillancourt/go-mongodb-replset)
[![Build Status](https://travis-ci.org/timvaillancourt/go-mongodb-replset.svg?branch=master)](https://travis-ci.org/timvaillancourt/go-mongodb-replset)
[![Go Report Card](https://goreportcard.com/badge/github.com/timvaillancourt/go-mongodb-replset)](https://goreportcard.com/report/github.com/timvaillancourt/go-mongodb-replset)
[![codecov](https://codecov.io/gh/timvaillancourt/go-mongodb-replset/branch/master/graph/badge.svg)](https://codecov.io/gh/timvaillancourt/go-mongodb-replset)

A package of golang structs for reading/modifying MongoDB replset config and state. The structs are to unmarshal the output of the ['replSetGetConfig'](https://docs.mongodb.com/manual/reference/command/replSetGetConfig/) and ['replSetGetStatus'](https://docs.mongodb.com/manual/reference/command/replSetGetStatus/) server commands

## Docs
- [github.com/timvaillancourt/go-mongodb-replset/config](https://godoc.org/github.com/timvaillancourt/go-mongodb-replset/config)
- [github.com/timvaillancourt/go-mongodb-replset/status](https://godoc.org/github.com/timvaillancourt/go-mongodb-replset/status)

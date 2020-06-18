# Hasgo2

This is an experimental version of [hasgo](https://www.github.com/DylanMeeus/hasgo) which includes
actual generics instead of generating them.

# How to get started

## Install Go from source

You'll need to compile Go from source, you can follow these instructions:
https://golang.org/doc/install/source. But, instead of checking out the stable branch make sure you
check out `dev.go2go`

## Setup your paths

In general it's pretty much the same as for stable go, but you can set the $GO2PATH as documented
here: https://go.googlesource.com/go/+/refs/heads/dev.go2go/README.go2go.md

## Get hacking

If everything went well - you should have a version of Go with generics up and running. You will
need to use the `go2go tool` to run 'Go2' code however. 

For example, to run the unit tests, run: `cd functions && go tool go2go test`

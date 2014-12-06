bank-ocr
========

## Usage

The program reads input from stdin, and reports results to stdout:

    ./bank-ocr < input.txt > report.txt


## Install

Compile and download binaries via http://gobuild.io/github.com/lnmx/bank-ocr

or

Install Go 1.3 from https://golang.org/dl/ and http://dl.golang.org/doc/install

    # create an empty workspace directory
    mkdir example

    cd example

    # set workspace path
    set GOPATH=`pwd`

    # get source
    go get github.com/lnmx/bank-ocr

    cd src/github.com/lnmx/bank-ocr

    # run unit tests
    go test -v ./...

    # build executable as ./bank-ocr
    go build



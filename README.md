# go on IBMi

This repository is based on all work done by the community in order to make the Go programming language available on IBMi.

This repo contains a fork of Go 1.24 with modifications merged from [https://github.com/JasonTashtego/go]. 
It's a wonderful job done by the community and here it's to have all existing information about all changes that need to be done to have go compiling on IBMi. I tested only on IBMi V7R5 with go 1.22.8 and go 1.24. It compiles fine go itself and hello world program.



## DISCLAIMER

It's an experimental project and it's not intended to be used in production. It's just a proof of concept to show that it's possible to compile go on IBMi.


## Congratulations

Congratulations to everyone who contributed to making go avalaible on IBMi! Very impressive job from all people involved in this open issue: [[golang#45017](https://github.com/golang/go/issues/45017)].

Here it's based on Jason Tashtego's repository that is the most updated and has all the changes that are needed to compile go on IBMi.

For more information about the reason of each change, please check the opened issue and Jason Tashtego's repository.

## List of changes

list of files impacted by the changes:

    - src/crypto/x509/cert_pool.go (updated)
    - src/crypto/x509/parser.go (updated)
    - src/os/removeall_at.go (updated)
    - src/os/removeall_noat.go (updated)
    - src/runtime/internal/goexperiment/exp_iseriesaix_on.go (added)
    - src/runtime/internal/goexperiment/exp_iseriesaix_off.go (added)
    - src/runtime/internal/goexperiment/flags.go (updated)
    - src/runtime/tagptr_64bit.go (updated)
    - src/runtime/malloc.go (updated)
     
    GCC-12 is needed to be able to compile with CGO_ENABLED=1 without issues and GOGC=off is mandatory to avoid crashes during garbage collector on IBMi.
    

## steps to compile go on IBMi

 The first step is to compile the go on linux or windows. Here I tested only on linux.
 update .go file : 

    #export GOEXPERIMENT=iseriesaix
    #export GOOS=aix
    #export GOARCH=ppc64
    export GOOS=linux
    export GOARCH=amd64
    export CC=gcc
    export CGO_ENABLED=1
    export GOMAXPROCS=1
    export GOROOT_BOOTSTRAP=~/go1.22.8
    export GOTMPDIR=~/tmp
    export GOROOT=~/go1.22.8
    export PATH=~/go1.22.8/bin:$PATH

    source .go

You need to have a go compiler on your Linux in this example it was 1.24. You download the go version you want to use for your IBMi.
Go into src and execute :
    
    ./all.bash

it generates the go binary 1.24 on linux.

Second step is to generate the boostrap for your IBMi for that updates the .go : 

    export GOEXPERIMENT=iseriesaix
    export GOOS=aix
    export GOARCH=ppc64
    #export GOOS=linux
    #export GOARCH=amd64
    export CC=gcc
    export CGO_ENABLED=0
    export GOMAXPROCS=1
    export GOROOT_BOOTSTRAP=~/go1.24
    export GOTMPDIR=~/tmp
    export GOROOT=~/go1.24
    export PATH=~/go1.24/bin:$PATH
    #issue with go1.24 and garbage collector 
    export GOGC=off

    source .go


Go into src and execute :

    ./bootstrap.bash

It will generate the go-aix-ppc64-boostrap.btz file that will be sent to the IBMi.

Third step is on IBMi side. Unzip the go-aix-ppc64-bootstrap.btz using :

    bzip2 -d go-aix-ppc64-bootstrap.btz
    tar -xvf go-aix-ppc64-bootstrap.tar
    cp -r go-aix-ppc64-bootstrap ~/go1.24
    
define the .go like : 

    export GOEXPERIMENT=iseriesaix
    export GOOS=aix
    export GOARCH=ppc64
    export CC=/QOpenSys/pkgs/bin/gcc-10
    export CGO_ENABLED=1
    export GOMAXPROCS=1
    export GOROOT_BOOTSTRAP=~/go1.24
    export GOTMPDIR=~/tmp
    export GOROOT=~/go1.24
    export PATH=~/go1.24/bin:$PATH
    export GOGC=off
    export GOPROXY=direct

    source .go

Go to go-aix-ppc64-bootstrap/src and execute :

    ./all.bash

And if all is correct you will see something like : 

    ##### Test execution environment.
    # GOARCH: ppc64
    # CPU: POWER10
    # GOOS: aix
    # OS Version: OS400 5 7 00780005CAB1

## current status

The go 1.24 needs to deactivate the garbage collector by setting GOGC=off.
The issue was :

    SIGSEGV: segmenation violation
    routine.getGCMaskOnDemand (0x100a59b80) .../src/runtime/type.go:108

After that at least go1.24 can compile the hello world example. 

Another issue is due to sigset_t struct already defined in os400 in /usr/include/sys/time.h (similar issue described here : https://community.ibm.com/community/user/power/discussion/gcc-struct-sigset-t-conflicts-with-aix-systimeh). Gcc-12 solves the issue. 

And go 1.24 is working fine on IBMi! 

The issues with go dependencies are due to proxy because some dependencies like github.com work fine but dependencies with redirection failed with timeout.
I'll test GOPROXY=direct to see if it works better.


# The Go Programming Language

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)
*Gopher image by [Renee French][rf], licensed under [Creative Commons 4.0 Attribution license][cc4-by].*

Our canonical Git repository is located at https://go.googlesource.com/go.
There is a mirror of the repository at https://github.com/golang/go.

Unless otherwise noted, the Go source files are distributed under the
BSD-style license found in the LICENSE file.

### Download and Install

#### Binary Distributions

Official binary distributions are available at https://go.dev/dl/.

After downloading a binary release, visit https://go.dev/doc/install
for installation instructions.

#### Install From Source

If a binary distribution is not available for your combination of
operating system and architecture, visit
https://go.dev/doc/install/source
for source installation instructions.

### Contributing

Go is the work of thousands of contributors. We appreciate your help!

To contribute, please read the contribution guidelines at https://go.dev/doc/contribute.

Note that the Go project uses the issue tracker for bug reports and
proposals only. See https://go.dev/wiki/Questions for a list of
places to ask questions about the Go language.

[rf]: https://reneefrench.blogspot.com/
[cc4-by]: https://creativecommons.org/licenses/by/4.0/

#export GOEXPERIMENT=iseriesaix
#export GOOS=aix
#export GOARCH=ppc64
export GOOS=linux
export GOARCH=amd64
export CC=gcc
export CGO_ENABLED=1
export GOMAXPROCS=1
export GOROOT_BOOTSTRAP=~/go1.22.9
export GOTMPDIR=~/tmp
export GOROOT=~/go1.22.9
export PATH=~/go1.22.9/bin:$PATH


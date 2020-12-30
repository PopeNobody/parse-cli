#!/usr/bin/env bash
this_dir=$(pwd)
export GOPATH=$HOME/gopath
export PATH=$HOME/gopath/bin:$PATH
rsync -az ./ $HOME/gopath/src/github.com/back4app/parse-cli/
cd $HOME/gopath/src/github.com/back4app/parse-cli
go get -t -v ./
echo Version $(parse-cli version) installed locally
rm -rf $HOME/gopath/bin/b4a-homolog
ln -s $HOME/gopath/bin/parse-cli $HOME/gopath/bin/b4a-homolog
cd $this_dir

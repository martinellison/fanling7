#!/usr/bin/env bash

export BASE=$(git rev-parse --show-toplevel)
if [[ "$BASE" == "" ]]
then
    echo "need to be in the git repository"
    exit 0
fi
cd $BASE
export PATH=$PATH:$BASE
MACHINE=`uname -n`
echo "machine is $MACHINE, setting machine-specific options"
case $MACHINE in
    edward) export GOX="$HOME/gocode";;
    pinkypi) export GOX="$HOME/golang";;
    *) export GOX=/work/golang;;
esac
echo "go code in $GOX"
export GOPATH="$GOX:$BASE"
echo "GOPATH is now " $GOPATH
export GODIR="/usr/local/go"
if [[ -d $GODIR ]]
then
    export PATH=$GODIR/bin:$PATH
fi

echo "pulling from git..."
git pull

$HOME/git/other2/beautify.sh
$BASE/build.sh

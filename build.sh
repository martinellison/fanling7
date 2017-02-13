#!/usr/bin/env bash

reset
echo "preparing to build..."
export BASE=$(git rev-parse --show-toplevel)
if [[ "$BASE" == "" ]]
then
    echo "need to be in the git repository"
    exit 0
fi
cd $BASE

echo "GOPATH is now " $GOPATH
MAINPACKAGES="fanling7"
CPPPACS="ui"
SUBPACS="global store engine itemset filepersist pagestore $CPPPACS"

echo "building c++..."
for PACKAGE in $CPPPACS
do
    astyle --style=attach --unpad-paren --align-pointer=type --remove-brackets $BASE/src/$PACKAGE/*.cpp $BASE/src/$PACKAGE/*.h
    rm -f $BASE/src/$PACKAGE/*.orig
    if [[ ! -d $BASE/temp/$PACKAGE/  ]]
    then
        mkdir -p $BASE/temp/$PACKAGE/
    fi
    swig -c++ -go  -cgo -outdir $BASE/temp/$PACKAGE/ -intgosize 64 $BASE/src/$PACKAGE/*.swigcxx
    rm -f $BASE/src/ui/ui_wrap.*
done

export CGO_CXXFLAGS="$(wx-config --cppflags) -std=c++14"
export CGO_LDFLAGS="$(wx-config --libs std,stc)"

echo "formatting..."
for PACKAGE in $MAINPACKAGES $SUBPACS
do
    gofmt -s -w -e $BASE/src/$PACKAGE
    FMTRES=$?
    if [[ $FMTRES != 0 ]]
    then
        echo "Format result is:" $FMTRES
        exit 1
    fi
    goimports -w $BASE/src/$PACKAGE
done
echo "testing..."
TESTPACKAGES="itemset filepersist pagestore"
mkdir -p cover
for PACKAGE in $TESTPACKAGES
do
    TESTBIN=bin/test$PACKAGE
    rm -f $TESTBIN
    go test -c -cover -o $TESTBIN $PACKAGE
    COVERPROF=cover/$PACKAGE.cover
    if [[ -f $TESTBIN ]]
    then
        echo '========= running test' $PACKAGE '========='
        TESTRUNRRES=99
        $TESTBIN -test.coverprofile=$COVERPROF
        TESTRUNRRES=$?
        if [[ $TESTRUNRRES != 0 ]]
        then
            echo "test failed"
            exit 1
        fi
        go tool cover -func=$COVERPROF | grep -v '100.0%'
    fi
done
echo "building..."
for MAINPACKAGE in $MAINPACKAGES
do
    MAIN=$BASE/bin/$MAINPACKAGE
    echo "building main:" $MAINPACKAGE "to:" $MAIN
    if [[ -f $MAIN ]]
    then
        rm $MAIN
    fi
    go build -o $MAIN $MAINPACKAGE
    BUILDRES=$?
    if [[ $BUILDRES != 0 ]]
    then
        echo $MAIN " build result is:" $BUILDRES
        exit 1
    fi
done
echo "vetting..."
for PACKAGE in $MAINPACKAGES $SUBPACS
do
    go tool vet -all $BASE/src/$PACKAGE
    VETRES=$?
    if [[ $VETRES != 0 ]]
    then
        echo "vet for $PACKAGE failed with status $VETRES"
        exit 1
    fi
done
if [[ -d cclog ]]
then
    for LOG in $(ls $BASE/log)
    do
        truncate --size 0 $LOG
    done
else
    mkdir -p $BASE/log
fi
if [[ ! -x $MAIN ]]
then
    echo $MAIN "is not executable"
    exit 1
fi
$BASE/../docgo/docgo --outdir doc $BASE/src/*/*.go

echo "build complete,"

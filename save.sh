#!/usr/bin/env bash

MSG="$1"
if [[ "$MSG" == "" ]]
then
    echo "need a commit message"
    exit 0
fi
export BASE=$(git rev-parse --show-toplevel)
if [[ "$BASE" == "" ]]
then
    echo "need to be in the git repository"
    exit 0
fi
cd $BASE
gofmt -w -s -l src
rm -f *~
echo `date +"%F %T"` "$MSG" >> $BASE/save-log.txt
git commit -m "$MSG" .
git push
echo 'saved'

#!/usr/bin/env bash

export BASE=$(git rev-parse --show-toplevel)
if [[ "$BASE" == "" ]]
then
    echo "need to be in the git repository"
    exit 0
fi
cd $BASE
echo "beautifying JS"
for JS in `find exper|grep '\.js$'|grep -v 'old/'` mkdoc/*.js
do
    echo 'beautifying' $JS
    js-beautify -t -d -w 80 "$JS" >"$JS.out"
    mv "$JS.out" "$JS"
done
echo "beautifying CSS"
for CSS in `find exper|grep '\.css$'|grep -v 'old/'`
do
    echo 'beautifying' $CSS
    css-beautify -t -d -w 80 "$CSS" >"$CSS.out"
    mv "$CSS.out" "$CSS"
done
echo "beautifying HTML"
for HTML in `find exper|grep '\.html$'|grep -v 'old/'` *.html mkdoc/*.html
do
    echo 'beautifying' $HTML
    html-beautify -t -d -w 80 "$HTML" >"$HTML.out"
    mv "$HTML.out" "$HTML"
done
BB=~/git/beautify_bash/beautify_bash.py
if [[ -f $BB ]]
then
    for F in *.sh scripts/*.sh exper/*.sh
    do
        if [[ -f $F ]]
        then
            echo 'reformatting bash script' $F
            $BB $F
        fi
    done
fi

echo "beautified"

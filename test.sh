#!/bin/bash

if [ -z $1 ]
then
  echo "Usage $0 [fill ratio]"
  exit 1
fi

make > /dev/null
FILL_RATIO=$1
SOURCE_FILE='testdata/normalized/1-100.txt'
UNFILLED_FILE="testdata/unfilled/1-100-$FILL_RATIO.txt"
GUESSED_FILE="testdata/guessed/1-100-$FILL_RATIO.txt"
COMPARISON_FILE="testdata/compared/1-100-$FILL_RATIO.txt"
./unfill -t -s 100 -f $FILL_RATIO < $SOURCE_FILE > $UNFILLED_FILE &&
./fill -d -s 100 -t < $UNFILLED_FILE > $GUESSED_FILE 2> test.fill.debug &&
./compare -d -t -s 100 $SOURCE_FILE < $GUESSED_FILE > $COMPARISON_FILE &&
./compare -t -td 1000 -s 100 $SOURCE_FILE < $GUESSED_FILE > test.compare.hugediffs

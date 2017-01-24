#!/bin/bash

if [ -z $1 ]
then
  echo "Usage $0 [fill ratio]"
  exit 1
fi

LOG_DIR='log'

if [ ! -d $LOG_DIR ] ; then
  mkdir $LOG_DIR
fi

FILL_RATIO=$1
SOURCE_FILE='testdata/normalized/1-100.txt'
UNFILLED_FILE="testdata/unfilled/1-100-$FILL_RATIO.txt"
GUESSED_FILE="testdata/guessed/1-100-$FILL_RATIO.txt"
COMPARISON_FILE="testdata/compared/1-100-$FILL_RATIO.txt"
make > /dev/null &&
./unfill -t -s 100 -f $FILL_RATIO < $SOURCE_FILE > $UNFILLED_FILE &&
./fill -d -s 100 -t < $UNFILLED_FILE > $GUESSED_FILE 2> $LOG_DIR/fill.log &&
./compare -t -d -td 1000 -s 100 $SOURCE_FILE < $GUESSED_FILE > /dev/null 2> $LOG_DIR/compare.log &&
tail -4 $LOG_DIR/compare.log
./badness -s=100 < $GUESSED_FILE

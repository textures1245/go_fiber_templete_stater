#!/bin/bash
echo "-= Production =-"
export AWS_PROFILE=payso
WORK_DIR=.
DEST_DIR=s3://s3-payso-gondor/
EXEC_FILE=payso-simple-noti
cd $WORK_DIR
echo "-= Build =-"
env GOOS=linux GOARCH=amd64 go build -o $WORK_DIR/$EXEC_FILE
env GOOS=linux GOARCH=arm64 go build -o $WORK_DIR/$EXEC_FILE-arm

echo "-= Deploy PRD =-"
aws s3 cp $WORK_DIR/$EXEC_FILE $DEST_DIR/$EXEC_FILE
aws s3 cp $WORK_DIR/$EXEC_FILE-arm $DEST_DIR/$EXEC_FILE-arm

echo "-= Deploy Finished =-"
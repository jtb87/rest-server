#!/bin/sh
echo "starting build & deploy"
go build -o todo-app

scp todo-app ec2:todo-app
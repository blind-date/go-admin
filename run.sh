#!/bin/bash
echo "... rm ..."
rm -rf go-admin
sleep 1

echo "... build ..."
go build main.go
sleep 1

echo "... killall ..."
killall -9 go-admin
sleep 1

echo "... chmod ..."
chmod -R 777 go-admin
sleep 1

echo "... run ..."
nohup ./go-admin &
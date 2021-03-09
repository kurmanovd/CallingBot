#!/bin/sh
cd ../
make clean
docker-compose.exe down
docker-compose.exe up --build -d
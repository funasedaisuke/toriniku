#!/bin/bash

# until mysqladmin ping -h mysql --silent; do
#   echo 'waiting for mysqld to be connectable...'
#   sleep 2
# done

echo "toriniku_app is starting...!"
exec go run main.go 
#!/bin/sh

mysql -uroot -proot --host="first_mysql" --execute="CREATE DATABASE IF NOT EXISTS first; \
    GRANT ALL PRIVILEGES ON first.* TO root@localhost;"
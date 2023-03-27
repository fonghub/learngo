#! /bin/bash

git pull origin master

read -p "please input commit comments " first

git add .
git commit -m $first

git push origin master

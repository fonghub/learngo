#! /bin/bash

read -p "please input commit comments " first

git add .
git commit -m $first

git push origin master

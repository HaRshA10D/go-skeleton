#! /usr/bin/env bash

rm -rf -f temp
mkdir temp
cd temp

if ! command -v git &> /dev/null
then
    echo "Git could not be found: install and configure git"
    exit
fi

if [ -z "$1" ]
  then
    echo "No argument supplied: provide project name: no spaces"
    echo "Exiting"
    exit
fi

rm -rf -f "$1"
rm -rf -f go-skeleton-temp

echo "Pulling skeleton repo"

PWD=$(pwd)

git clone https://github.com/HaRshA10D/go-skeleton.git "$PWD/$1"

cd "$1"

rm -rf .git

rm bootstrap.sh
sed -i -e "s/skeleton-name/$1/" go.mod

git init

git add .
git commit -m "project setup"

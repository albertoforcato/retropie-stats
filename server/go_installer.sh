#!/bin/bash

RED=$'\e[1;31m'
GREEN=$'\e[1;32m'
YELLOW=$'\e[1;33m'
BLUE=$'\e[1;34m'
ENDCOLOR=$'\e[0m'

check_installation () {
    if ! [ -x "$(command -v $1)" ];
    then
        printf "${RED}✗ $1 not installed${ENDCOLOR}\n"
        exit 1
    else
        printf "${GREEN}✔ $1 installed${ENDCOLOR}\n"
    fi
}

printf "\n${YELLOW}Install GOLANG${ENDCOLOR}\n"
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install -y golang-go
check_installation go

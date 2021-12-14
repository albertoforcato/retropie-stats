#!/bin/bash

export GOLANG=go1.17.5.linux-armv6l.tar.gz
wget https://dl.google.com/go/$GOLANG
sudo mv go /usr/local
sudo sudo tar -xvf $GOLANG
rm $GOLANG
unset GOLANG
echo GOROOT="/usr/local/go" >>~/.profile
# shellcheck disable=SC2016
echo PATH='$GOROOT'/bin:'$PATH' >>~/.profile

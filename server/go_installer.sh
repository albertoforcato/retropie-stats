export GOLANG=go1.17.5.linux-armv6l.tar.gz
wget https://go.dev/dl/$GOLANG
sudo tar -C /usr/local -xzf $GOLANG
rm $GOLANG
unset GOLANG

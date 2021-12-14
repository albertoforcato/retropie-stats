export GOLANG=go1.17.5.linux-armv6l.tar.gz
wget https://go.dev/dl/$GOLANG
sudo tar -C /usr/local -xzf $GOLANG
rm $GOLANG
unset GOLANG
echo PATH="$PATH:/usr/local/go/bin" >>~/.profile
echo GOPATH="$HOME/go" >>~/.profile
cd ~ && source ~/.profile

FROM golang:1.17

RUN adduser --disabled-password --gecos "" go

ARG UNAME=go
ARG UID=1000
ARG GID=1000

RUN if grep -q $GID /etc/group; then GROUP_NAME=$(grep :$GID: /etc/group | sed -En "s/(.*):x:$GID:/\1/p") && if [ "$GROUP_NAME" != "$UNAME" ]; then delgroup $GROUP_NAME; fi; fi;
RUN if grep -q $UNAME /etc/group; then groupmod -g $GID $UNAME; else groupadd -g $GID $UNAME; fi;
RUN usermod -u $UID $UNAME

RUN su go -s /bin/sh -c "mkdir ~/go"

ENV GOPATH=""
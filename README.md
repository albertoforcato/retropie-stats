# RetroPie Stats

This golang program is used to gather statistics about the RetroPie game usage. Record the game you played, and the time
you played it.

## Server side

Change directory to the server side of the program.

```shell
cd server
```

### Installation

#### **[OPTIONAL]** Install GO

```shell
make install-go
```

#### Add command to retropie system

Append to `runcommand-onstart.sh` and `runcommand-onend.sh` retropie-stats command in order to run on game startup and
shutdown.

```shell
make init
```

#### Install server and commands cli

Finally, install the server and the cli commands

```shell
make install
```
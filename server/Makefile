init:
	cat runcommand_hooks/runcommand-onstart.sh >> /opt/retropie/configs/all/runcommand-onstart.sh
	cat runcommand_hooks/runcommand-onend.sh >> /opt/retropie/configs/all/runcommand-onend.sh
install-go:
	sudo chmod +x go_installer.sh && ./go_installer.sh
build-commands:
	go build ./cmd/commands/...
build-server:
	go build ./cmd/server/...
run:
	go run cmd/server/main.go;
mod-tidy:
	go mod tidy
install:
	make build-commands
	chmod +x commands && cp commands ~/bin/retropie-stats
	make build-server
	chmod +x server && cp server ~/bin/retropie-stats-start-server
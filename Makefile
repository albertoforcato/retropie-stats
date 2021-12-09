CUR_UID := $(shell id -u)
CUR_GID := $(shell id -g)

DOCKER_DIR := .dev/docker
DOCKERFILE := docker-compose.yml
DOCKER_SERVICES := sqlite3

docker_compose := cd ${DOCKER_DIR} && docker-compose -f ${DOCKERFILE}

run_as_go_user := run --rm --use-aliases --user go
go_server := ${docker_compose} ${run_as_go_user} go
go_cli := cd ${DOCKER_DIR} && CMD_NAME=cli_go SERVICE_ALIAS="" docker-compose -f ${DOCKERFILE} ${run_as_go_user} go

init-project:
	cd .dev/docker && mkdir -p data
	make dist-files build-docker-services
	mkdir -p ~/go/pkg
dist-files:
	[ -f ${DOCKER_DIR}/.env ] || cp ${DOCKER_DIR}/.env.dist ${DOCKER_DIR}/.env \
	&& [ -f ${DOCKER_DIR}/docker-compose.yml ] || cp ${DOCKER_DIR}/docker-compose.yml.dist ${DOCKER_DIR}/docker-compose.yml \
	&& cd -
build-docker-services:
	${docker_compose} build --build-arg UID=${CUR_UID} --build-arg GID=${CUR_GID} ${DOCKER_SERVICES_BUILD}
start-docker-services:
	${docker_compose} up -d ${DOCKER_SERVICES}
stop-docker-services:
	${docker_compose} stop ${DOCKER_SERVICES}
restart-docker-services:
	${docker_compose} restart ${DOCKER_SERVICES}
purge-docker-services:
	${docker_compose} down
run:
	${go_server} run server/cmd/main.go
mod-tidy:
	${go_cli} mod tidy
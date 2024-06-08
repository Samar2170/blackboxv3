IP := $(shell hostname -I | awk '{print $$1}')
PATTERN:= HOSTNAME=*
UPLOADSDIRPATTERN:= UPLOADSDIR=*
BASEDIR := $(shell pwd)
HOMEDIR := $(shell echo $$HOME)
UPLOADSDIR := $(HOMEDIR)/BlackboxDrive/


clear-existing-hostname:
	@grep -vE "^$(PATTERN)" < .env > .env.tmp
	@mv .env.tmp .env

clear-existing-uploadsdir:
	@grep -vE "^$(UPLOADSDIRPATTERN)" < .env > .env.tmp
	@mv .env.tmp .env

make-uploads-dir:
	@mkdir -p $(UPLOADSDIR)

set-hostname:
	@echo "HOSTNAME=$(IP)" >> .env
	@echo "Hostname set to $(IP)"

set-uploads-dir:
	@echo "UPLOADSDIR=$(UPLOADSDIR)" >> .env
	@echo "Uploads directory set to $(UPLOADSDIR)"

start-server:
	@cd $(BASEDIR)
	@echo "Starting server..."
	@go run .

proto-gen: 
	@echo "Going to api directory... $(BASEDIR)/api"
	@cd $(BASEDIR)/api
	@echo "in directory $(shell pwd)"
	@echo "Generating proto files..."
	@protoc --go_out=./apigen --go_opt=paths=source_relative  --go-grpc_out=./apigen --go-grpc_opt=paths=source_relative *.proto

.PHONY:
run:
	@make clear-existing-hostname
	@make set-hostname
	@make clear-existing-uploadsdir
	@make set-uploads-dir
	@make make-uploads-dir
	@make start-server
	
.PHONY:
proto:
	@make proto-gen
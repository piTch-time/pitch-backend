.PHONY: run swagger build docker up down clean

GO ?= GO111MODULE=on go
APP_NAME = pitch
BIN_DIR = ./bin
BUILD_DIR = ./application/cmd
BUILD_FILE = $(addprefix $(BUILD_DIR)/, main.go)

# local run
run:
	make swagger
	$(GO) run $(BUILD_FILE)

# generate swagger
swagger:
	echo "Update swagger to /docs"
	swag init  -g ./application/cmd/main.go

# build binary
build:
	$(GO) build -ldflags="-s -w" -o $(BIN_DIR)/$(APP_NAME) $(BUILD_FILE)

docker:
	make dbuild && make drun

# docker build
dbuild:
	docker build \
		-t $(APP_NAME):latest \
		-f Dockerfile --no-cache .

# docker local run
drun:
	docker run --rm -p 8080:8080 --name pitch pitch

# docker compose up
up:
	docker compose up -d --build --remove-orphans
	docker compose logs -f

# docker compose down
down:
	docker compose down --rmi local

# rm  binary		
clean:
	echo "remove bin exe"
	rm -f $(addprefix $(BIN_DIR)/, $(APP_NAME))
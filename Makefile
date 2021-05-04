APP_NAME = service
BIN = ./bin/$(APP_NAME)
APIARY = ./files/apiary

.PHONY: build

build:
	@GOOS=linux GOARCH=amd64 go build -o $(BIN) ./cmd

apiary:
	@rm -f ./apiary.apib
	@awk '{print}' $(APIARY)/*.apib > $(APIARY)/index.apib
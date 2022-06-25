# https://github.com/brunoshiroma/go-gin-poc/blob/main/Dockerfile
# Build Step
FROM golang:alpine AS build

RUN apk update
RUN apk add make
RUN apk add git

WORKDIR /go/github.com/piTch-time/pitch-backend
COPY . .
# RUN export PATH=$(go env GOPATH)/bin:$PATH
# RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go mod tidy
RUN GO111MODULE=on go build -ldflags="-s -w" -o pitch ./application/cmd/main.go
# RUN swag init -g ./application/cmd/main.go --output=./docs

# Final Step
FROM alpine as runtime


# Base packages
RUN apk update
RUN apk upgrade
RUN apk add --no-cache bash
RUN apk --no-cache add curl
RUN apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
RUN rm -rf /var/cache/apk/*

WORKDIR /home
# Copy binary from build step
COPY --from=build /go/github.com/piTch-time/pitch-backend/pitch pitch
# Copy config files to runtime
COPY --from=build /go/github.com/piTch-time/pitch-backend/infrastructure infrastructure
# Define timezone
ENV TZ=Asia/Seoul

# sandbox phase will ignore this command, so this docker file cmd care for prod phase.
CMD [ "/home/pitch", "-phase=prod"]
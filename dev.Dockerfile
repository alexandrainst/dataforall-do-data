FROM golang:1.24

RUN apt update

# These commands are for installing nodejs
RUN apt install -y curl
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash -
RUN apt install -y nodejs

RUN mkdir /app

# Install reflex used as the watcher/runner for go (PocketBase)
RUN go install github.com/cespare/reflex@latest


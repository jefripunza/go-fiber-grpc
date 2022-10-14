# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
MAINTAINER Jefri Herdi Triyanto, jefri.triyanto@goapotik.com

WORKDIR /app
COPY . .

# 🌊 Install Dependencies
RUN go mod download

# ⚒️ Build
RUN go build -o ./run

# 🚫 Remove Project Files & Folders
RUN ls -a &&\
    rm -rf .vscode &&\
    rm -rf docs &&\
    rm -rf proto &&\
    rm -rf src &&\
    rm .gitignore &&\
    rm Dockerfile &&\
    rm go.mod &&\
    rm go.sum &&\
    rm main.go &&\
    rm package.json || true &&\
    rm README.md

# 💯 Configuration Environment
RUN sed -i 's/localhost/172.17.0.1/g' .env

# 🚀 Finish !!
CMD [ "./run" ]
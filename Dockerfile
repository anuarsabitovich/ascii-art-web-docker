FROM golang:1.19-alpine3.17

LABEL name = "ASCII-ART-WEB-DOCKERIZE"
LABEL authors = "anuarsabitovich, Bakytzhan16"
LABEL release date = "10.01.23"

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]








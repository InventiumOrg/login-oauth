FROM golang:1.21

ARG SERVICE_NAME=default

RUN mkdir /${SERVICE_NAME}

WORKDIR /${SERVICE_NAME}

COPY . .

RUN go mod download 

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/inventium .

CMD ["/usr/local/bin/inventium"]
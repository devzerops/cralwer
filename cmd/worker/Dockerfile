FROM golang:1.23.4

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/worker

RUN go build -o /worker

EXPOSE 8081

CMD [ "/worker" ]\
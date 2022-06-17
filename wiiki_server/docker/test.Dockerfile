FROM golang:1.18.3-buster AS tester

WORKDIR /develop

RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y postgresql-client
RUN psql --version
RUN apt-get install -y openssl
RUN apt-get install -y tzdata 
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
RUN apt-get install -y ca-certificates

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make goose_build

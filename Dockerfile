FROM golang:1.17

RUN apt-get -qq -y update
RUN apt-get -qq -y upgrade

RUN apt-get install -qq -y make

WORKDIR /work
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go get -v golang.org/x/lint/golint
COPY . .

RUN make

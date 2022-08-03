FROM golang:1.18

RUN apt-get -qq -y update
RUN apt-get -qq -y upgrade

RUN apt-get install -qq -y make python

WORKDIR /work
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install -v golang.org/x/lint/golint@latest
COPY . .

RUN make clean verbose
RUN cat examples.txt

FROM golang:1.19-alpine3.16

RUN apk add make python3 bash

WORKDIR /work
RUN go install honnef.co/go/tools/cmd/staticcheck@2022.1
RUN go install -v golang.org/x/lint/golint@latest
COPY . .
RUN chmod +x l1c
RUN ln -s /usr/bin/python3 /usr/bin/python
RUN make clean verbose

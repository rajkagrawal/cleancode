FROM golang:1.14.4-alpine3.12 AS builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN make install

FROM scratch
WORKDIR /
COPY --from=builder /cleancode .


ENTRYPOINT ["/cleancode"]


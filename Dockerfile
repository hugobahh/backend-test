FROM golang:1.22.2-alpine3.19 AS builder

WORKDIR $GOPATH/src/
COPY . .

RUN go mod download
RUN wget https://truststore.pki.rds.amazonaws.com/global/global-bundle.pem -O rds-combined-ca-bundle.pem
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o build-backend-test ./cmd/backend-test

FROM gcr.io/distroless/static-debian11
ARG version
ENV VERSION=${version}

WORKDIR /

COPY --from=builder --chown=nonroot:nonroot /go/src/backend-test/build-backend-test ./build-backend-test

USER nonroot
EXPOSE 5003

CMD ["./build-backend-test"]
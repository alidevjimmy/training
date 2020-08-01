FROM golang as builder

ENV GO111MODULE=on
WORKDIR /training

COPY go.mod .
COPY go.sum .
	
RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

#ENTRYPOINT ["./training"]
FROM scratch

COPY --from=builder /training/ /app/

ENTRYPOINT ["/app/training"]

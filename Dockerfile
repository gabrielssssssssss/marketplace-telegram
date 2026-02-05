FROM golang:1.25

WORKDIR /app

ENV GODEBUG=netdns=go+1
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-marketplace-telegram ./cmd/marketplace-telegram/main.go

EXPOSE 5432:5432

CMD ["/docker-marketplace-telegram"]
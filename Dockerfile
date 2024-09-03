FROM golang:alpine3.18 AS buildStage
LABEL author="github.com/rydhshlkhn"

ENV GO111MODULE=on CGO_ENABLED=1
WORKDIR /mirae
COPY ./go.mod go.mod
COPY ./go.sum go.sum
COPY ./cmd cmd
COPY ./config config
COPY ./delivery delivery
COPY ./domain domain
COPY ./infra infra
COPY ./mocks mocks
COPY ./repository repository
COPY ./usecase usecase

RUN apk add upx util-linux-dev build-base
RUN go mod download
RUN go test ./delivery/resthandler ./usecase ./repository -v
ENV GO111MODULE=on CGO_ENABLED=0
RUN go build -a -gcflags=all="-l -B" -ldflags "-s -w" -o ./mirae-app cmd/main.go
RUN upx -9 --lzma /mirae/mirae-app
RUN echo 'DATABASE_URL=postgres://mirae-user:mirae-password@tcp(db:5432)/mirae?sslmode=disable&TimeZone=Asia/Jakarta' > /mirae/temp.conn

FROM gcr.io/distroless/static:latest AS runtimeStage
WORKDIR /app
COPY --from=buildStage /mirae/mirae-app .
COPY --from=buildStage /mirae/temp.conn .env
ENTRYPOINT ["/app/mirae-app"]
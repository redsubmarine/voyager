FROM golang:1.22.1-alpine3.19 AS builder

WORKDIR /server
COPY . .

RUN go build -o /server/server

FROM scratch
COPY --from=builder /server/server /server
COPY --from=builder /server/app.env /app.env
ENTRYPOINT [ "./server" ]

FROM golang:1.22.3 as builder

WORKDIR /app
COPY .  .
RUN go env -w GOPROXY=https://goproxy.cn,direct && env make build

FROM alpine:3.17
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
WORKDIR /app
COPY --from=builder /app/dist/ikubeops-cli .
COPY config /app/config
CMD ["./ikubeops-cli", "start"]

# docker buildx build --platform linux/amd64  -t harbor.ikubeops.local/public/ingress-manager:latest .
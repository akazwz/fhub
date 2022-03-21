FROM golang:1.18 AS devlopement
MAINTAINER akazwz
WORKDIR /home/fhub
ADD . /home/fhub
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build -o app -buildvcs=false --ldflags "-extldflags -static"

FROM alpine:latest AS production
WORKDIR /root/
COPY --from=devlopement /home/fhub/app .
EXPOSE 8080:8080
ENV GIN_MODE=release
ENTRYPOINT ["./app"]
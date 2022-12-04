FROM golang:1.19.2-alpine
RUN apk update &&  apk add git
ENV TZ /usr/share/zoneinfo/Asia/Tokyo
ENV ROOT=/golangStudy
WORKDIR ${ROOT}
ENV GO111MODULE=on
COPY . .
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]

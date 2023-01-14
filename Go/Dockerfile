FROM golang:1.16-alpine

WORKDIR /gameOfLife

COPY *.go ./

EXPOSE 8090

CMD ["go", "run", "grid.go" ,"main.go"]
FROM golang

WORKDIR /app

COPY . /app

# installing dependency
RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "main.go"]

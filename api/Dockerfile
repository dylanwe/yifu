FROM golang:1.22.3-alpine3.19
WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 8080
CMD ["go", "run", "."]


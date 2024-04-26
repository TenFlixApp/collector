FROM golang:1.22

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY data/ ./data/
COPY helpers/ ./helpers/
COPY routes/ ./routes/
COPY *.go ./
RUN go build -o ./export.out

EXPOSE 8089/tcp
CMD ["/app/export.out"]

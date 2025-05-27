FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o math

CMD [ "./math" ]
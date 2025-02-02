FROM golang

WORKDIR /app

COPY . .

RUN go mod tidy 

EXPOSE 8080

CMD ["go","run", "."]
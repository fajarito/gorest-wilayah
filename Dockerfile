FROM --platform=linux/amd64 golang:1.20-alpine3.17
WORKDIR /app
COPY . .
RUN go build -o gorest-pk
EXPOSE 8080
CMD ./gorest-pk
FROM --platform=linux/amd6 golang:1.19-alpine3.16
WORKDIR /app
COPY . .
RUN go build -o gorest-pk
EXPOSE 8080
CMD ./gorest-pk
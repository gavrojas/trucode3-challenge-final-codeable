FROM golang:1.21.13-alpine

RUN mkdir app

WORKDIR /app

COPY trucode_data_management_api/go.mod trucode_data_management_api/go.sum ./

RUN go mod download

ARG PORT=:443 \
    GIN_MODE=release

ENV PORT=${PORT} \
    GIN_MODE=${GIN_MODE}

COPY trucode_data_management_api/ .
COPY trucode_data_management_tools/fullchain.pem /etc/ssl/certs/ 
COPY trucode_data_management_tools/privkey.pem /etc/ssl/private/

RUN go build -o main .

CMD ["./main"]
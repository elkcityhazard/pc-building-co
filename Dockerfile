FROM golang:1.26.1-alpine

RUN apk add --no-cache nodejs npm

WORKDIR /var/www

ARG BIN_NAME="bin/pc_building_co"
ENV BIN_NAME=${BIN_NAME}
ENV MARIADB_HOST=""
ENV MARIADB_PORT=""
ENV MARIADB_USER=""
ENV MARIADB_PASSWORD=""
COPY package.json package-lock.json* ./
RUN npm install

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

# Build assets and binary
RUN npm run css:build \
    && npm run webpack:build \
    && go build -o "${BIN_NAME}" cmd/web/*.go

RUN chmod +x entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/var/www/entrypoint.sh"]

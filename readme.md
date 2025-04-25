## EXAMPLE MAKEFILE

```
WEBSITE_NAME="P.C. Building Company"
WEBSITE_URL="http://localhost"
BINARY_NAME=bin/pc_building_co
SMTP_HOST=localhost
SMTP_PORT=1025
SMTP_USERNAME=username
SMTP_PASSWORD=password
SMTP_FROM_ADDR=from@test.com
SMTP_TO_ADDR=to@test.com

run:
	if [ -f "${BINARY_NAME}" ]; then \
		rm ${BINARY_NAME}; \
	fi; \
	go build \
	-o ${BINARY_NAME} cmd/web/*.go && \
	bin/pc_building_co \
	-website_name=${WEBSITE_NAME} \
	-website_url=${WEBSITE_URL} \
	-smtp_host=${SMTP_HOST} \
	-smtp_port=${SMTP_PORT} \
	-smtp_username=${SMTP_USERNAME} \
	-smtp_password=${SMTP_PASSWORD} \
	-smtp_toaddr=${SMTP_TO_ADDR} \
	-smtp_fromaddr=${SMTP_FROM_ADDR}

```


## EXAMPLE dockerfile

```
FROM  golang:1.24.1-alpine
RUN apk add --no-cache nodejs npm
ENV BIN_NAME ${BIN_NAME}

COPY entrypoint.sh /entrypoint.sh

WORKDIR /var/www

COPY . .

RUN mkdir -p bin/ node_modules/

EXPOSE 8080

RUN npm install \
&& npm run css:build \
&& npm run webpack:build \
&& go build \
-o ${BIN_NAME} \
cmd/web/*.go

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
```

## Example Docker Compose

```
services:
  mailhog:
    image: mailhog/mailhog
    restart: always
    logging:
      driver: none
    ports:
      - 1025:1025
      - 8025:8025

```

## Example entrypoint.sh

```
#!/bin/sh
BIN_NAME=bin\/mybinname
WEBSITE_NAME="Awesome Website Name"
SMTP_HOST="smtp.mysmtpprovider.com"
SMTP_PORT=465
SMTP_USERNAME=smtpuser@smtphost.com
SMTP_PASSWORD=myverysecretpassword
SMTP_TOADDR=whereiwantservertosendemailto@example.com
SMTP_FROMADDR=fromaddressaliasofserver@exampe.com


exec ${BIN_NAME} \
    -website_name="${WEBSITE_NAME}" \
    -website_url="${WEBSITE_URL}" \
    -smtp_host="${SMTP_HOST}" \
    -smtp_port="${SMTP_PORT}" \
    -smtp_username="${SMTP_USERNAME}" \
    -smtp_password="${SMTP_PASSWORD}" \
    -smtp_toaddr="${SMTP_TOADDR}" \
    -smtp_fromaddr="${SMTP_FROMADDR}"


```

## Build

```
docker buildx build --platform linux/arm64,linux/amd64 -t
dockerhubname/dockerhubproject:latest --push .
```

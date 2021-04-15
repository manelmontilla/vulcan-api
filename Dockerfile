# Copyright 2021 Adevinta

FROM golang:1.13.3-alpine3.10 as builder

WORKDIR /app

COPY . .

RUN cd cmd/vulcan-api && GOOS=linux GOARCH=amd64 go build -mod vendor . && cd -

FROM alpine:3.10

ENV FLYWAY_VERSION 6.1.4
WORKDIR /flyway

RUN apk add --no-cache --update openjdk8-jre bash gettext libc6-compat
RUN wget https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/${FLYWAY_VERSION}/flyway-commandline-${FLYWAY_VERSION}.tar.gz \
    && tar -xzf flyway-commandline-${FLYWAY_VERSION}.tar.gz && mv flyway-${FLYWAY_VERSION}/* . \
    && rm flyway-commandline-${FLYWAY_VERSION}.tar.gz \
    && find /flyway/drivers/ -type f -not -name 'postgres*' -delete \
    && ln -s /flyway/flyway /bin/flyway

ARG BUILD_RFC3339="1970-01-01T00:00:00Z"
ARG COMMIT="local"

ENV BUILD_RFC3339 "$BUILD_RFC3339"
ENV COMMIT "$COMMIT"

WORKDIR /app

COPY db/sql /app/sql/

RUN mkdir -p /app/output

COPY --from=builder /app/cmd/vulcan-api/vulcan-api .

COPY config.toml .
COPY run.sh .

CMD [ "./run.sh" ]
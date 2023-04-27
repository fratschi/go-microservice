###################################################################################
## Multistage docker build for creating a smallest possible docker container
###################################################################################

ARG GO_VERSION=1.20

## Stage 1
## Prepare dev environment for building services

FROM golang:${GO_VERSION}-alpine AS dev

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

ENV GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOFLAGS="-mod=vendor"


## Stage 2
## Downloading required modules and building go services in separate build environment

FROM dev as build

ENV USER=serviceuser
ENV UID=10001

RUN mkdir /compile && mkdir vendor
COPY . /compile
WORKDIR /compile

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

## Build
RUN (([ ! -d "./vendor" ] && go mod download && go mod vendor) || true) && go build -ldflags="-s -w" -mod vendor  ./cmd/main.go && chmod +x main

## remove all groups and users except serviceuser
RUN cat /etc/passwd | grep -e "serviceuser" > /pwd  && cat /etc/group | grep -e "serviceuser" > /grp

## Stage 3
## Assemble final services container from an empty scratch image

FROM scratch AS service

COPY --from=build /compile/main /service
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /pwd /etc/passwd
COPY --from=build /grp /etc/group

EXPOSE 8080

USER serviceuser:serviceuser

ENTRYPOINT ["/service"]
CMD ""
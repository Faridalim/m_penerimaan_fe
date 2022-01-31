FROM golang:1.17-alpine AS build-env
ADD . /src
RUN cd /src && go build -o goapp

FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
COPY --from=build-env /src/pages /app/pages/
ENTRYPOINT ./goapp
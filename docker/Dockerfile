#### front ####
FROM node:latest AS front

WORKDIR /front 

COPY cmd/webapp/front/src ./src
COPY cmd/webapp/front/*.json .

RUN ["npm", "run", "build"]

#### webapp ###
FROM golang:1.19.7 AS app

WORKDIR /app 

ENV GOARCH=amd64
# ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GO111MODULE=on

COPY cmd/webapp/*.go ./cmd/webapp/
COPY internal ./internal
COPY pkg ./pkg
COPY go.mod go.sum ./
COPY cmd/webapp/.env ./cmd/webapp/

### webapp ###
WORKDIR /app/cmd/webapp

RUN ["go", "build", "."]

FROM alpine:3.17.2

WORKDIR /app
COPY --from=front /front/dist ./front/dist
COPY --from=app /app/cmd/webapp/webapp ./webapp
COPY --from=app /app/cmd/webapp/.env ./.env

EXPOSE 3000
ENTRYPOINT [ "./webapp" ]

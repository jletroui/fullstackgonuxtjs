FROM node:20-alpine3.20 AS NODE_BUILD

WORKDIR /app
ENV NODE_ENV=production

COPY ./frontend/package*.json /app/
RUN npm ci

COPY ./frontend /app
RUN npm run build


FROM golang:1.23-alpine3.20 AS GO_BUILD

WORKDIR /app

COPY ./backend /app/backend
RUN cd /app/backend \
 && go get -d -v ./... \
 && go build -o backend .

COPY ./config /app/config


FROM alpine:3.20

WORKDIR /app
# Need to provide ENV=production or ENV=staging to execute

COPY --from=GO_BUILD /app/backend/backend /app/backend
COPY ./ops/db/migrations /app/ops/db/migrations
COPY --from=NODE_BUILD /app/.output/public /app/www
# NOT TRUE, TODO: Little note: getting the config file from BUILD ensures that the secrets are encrypted, because we ran the build
COPY --from=GO_BUILD /app/config/*.ejson /app/config/

RUN ls -la && env

EXPOSE 8080

ENTRYPOINT ["/app/backend"]

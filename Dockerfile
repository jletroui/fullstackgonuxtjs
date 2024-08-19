FROM golang:1.23-alpine3.20 AS BUILD

WORKDIR /app
ENV NODE_ENV=production

COPY ./backend /app/backend
RUN cd /app/backend \
 && go get -d -v ./... \
 && go build -o backend .

# COPY ./frontend /app/frontend/
# RUN rm /app/frontend/jest.config.ts
# RUN gradle :frontend:build -x :frontend:check

COPY ./config /app/config
COPY ./ops/db/migrations /app/ops/db/migrations

FROM alpine:3.20

WORKDIR /app
# Need to provide ENV=production or ENV=staging to execute

COPY --from=BUILD /app/backend/backend /app/backend
COPY --from=BUILD /app/ops/db/migrations /app/ops/db/migrations
#COPY --from=BUILD /app/frontend/out /app/www
# Little note: getting the config file from BUILD ensures that the secrets are encrypted, because we ran the build
COPY --from=BUILD /app/config/*.ejson /app/config/

RUN ls -la && env

EXPOSE 8080

ENTRYPOINT ["/app/backend"]

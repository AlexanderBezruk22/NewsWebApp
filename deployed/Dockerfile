FROM golang:1.24.1-alpine

ARG UID
ARG GID

RUN addgroup -g ${GID} newsapp && adduser -D -u ${UID} -G newsapp newsapp
USER newsapp

RUN go install github.com/air-verse/air@v1.61.7

WORKDIR /app

COPY --chown=newsapp:newsapp ../go.mod ../go.sum ./
RUN go mod download

COPY --chown=newsapp:newsapp .. /app
CMD ["air", "-c", ".air.toml"]
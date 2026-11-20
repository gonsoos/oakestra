FROM golang:alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /bin/autoscaler ./main.go

FROM gcr.io/distroless/static-debian12:latest@sha256:61b7ccecebc7c474a531717de80a94709d20547cdcdaf740c25876f2a8e38b44

COPY --from=build-stage /bin/autoscaler /autoscaler

ENV RESOURCE_ABSTRACTOR_URL=localhost
ENV RESOURCE_ABSTRACTOR_PORT=11011

ENV DEBUG=False

ENTRYPOINT ["/autoscaler"]
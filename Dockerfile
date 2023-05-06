FROM golang:1.20-buster AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 go build -o /bin/app
COPY app.env /bin/app.env

FROM gcr.io/distroless/static-debian11 AS build-release-stage

COPY --from=build /bin/app /bin
COPY --from=build /bin/app.env /

EXPOSE 9000-9999

ENTRYPOINT [ "/bin/app" ]

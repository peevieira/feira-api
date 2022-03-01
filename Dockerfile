FROM golang as base

WORKDIR /src

COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

FROM base AS iniciar
EXPOSE 8181
CMD ["src"]
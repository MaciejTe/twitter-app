FROM golang:1.16.5 as build

WORKDIR /app

# Copy golang dependency manifests
COPY go.mod .
COPY go.sum .

# Cache the downloaded dependency in the layer
RUN go mod download

# Copy the source code
COPY . .

ARG SKAFFOLD_GO_GCFLAGS
RUN CGO_ENABLED=0 go build -o /app/twitter .

FROM scratch as run
ENV GOTRACEBACK=single

COPY --from=build /app/twitter /bin/twitter
ENTRYPOINT ["/bin/twitter"]

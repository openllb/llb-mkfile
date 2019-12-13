# syntax = docker/dockerfile:experimental
FROM golang:1.12 AS golang
WORKDIR /go/src/mkfile/ 
COPY . /go/src/mkfile/ 
# RUN --mount=target=. \
#     --mount=type=cache,target=/root/.cache \
#     --mount=type=cache,target=/go/pkg/mod \
RUN GO111MODULE=on CGO_ENABLED=0 go build -o /frontend --ldflags "-s -w" ./cmd/mkfile-frontend

FROM scratch AS release
COPY --from=golang /frontend /bin/frontend
ENTRYPOINT ["/bin/frontend"]

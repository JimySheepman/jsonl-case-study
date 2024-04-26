# stage 1
FROM golang:1.22.1-alpine AS BUILDER

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

ARG revisionID
ARG commitID
ENV REVISION_ID ${revisionID}
ENV COMMIT_ID ${commitID}

RUN CGO_ENABLED=0 go build -ldflags "-s -w -X main.Revision=`date -u +%Y.%-m.%-d`.${REVISION_ID} -X main.Commit=${COMMIT_ID}" -o app ./cmd/job/.

# stage 2
FROM alpine

COPY --from=BUILDER /app/app /app
COPY --from=BUILDER /app/config.yml /config.yml

ENTRYPOINT ["/app"]
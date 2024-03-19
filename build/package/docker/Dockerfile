#####################################
### Step 1: Build the application ###
#####################################
FROM golang:alpine AS builder

# Install git
RUN apk update && apk add --no-cache git

# SECURITY - scratch image does not come with ca-certificates / timezone data
# Install ca-certificates to allow SSL services
RUN apk add --no-cache ca-certificates tzdata && update-ca-certificates
#
# SECURITY - NEVER RUN A PROCESS AS ROOT IN A CONTAINER
# Create user for application
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/questions/49955097/how-do-i-add-a-user-when-im-using-alpine-as-a-base-image/55757473#55757473
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# Set the working directory inside builder
WORKDIR $GOPATH/src/go-kit/app/
COPY . .
RUN go mod tidy
RUN GOOS=linux \
    GOARCH=arm64 \
    go build \
    -ldflags="-w -s -X main.VersionCommit=$(git rev-list --abbrev-commit --abbrev=8 -1 HEAD) -X main.VersionSemVer=$(git name-rev HEAD --name-only)" \
    -o /go/bin/app \
    examples/app/main.go

###################################
### Step 1: Build a small image ###
###################################

FROM scratch
# Import the user and group files from the builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Import the executable
COPY --from=builder /go/bin/app /go/bin/app

# Use an unprivileged user
USER appuser:appuser
ENTRYPOINT ["/go/bin/app", "version"]
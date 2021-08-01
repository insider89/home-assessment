############################
# STEP 1 build executable binary
############################
FROM golang:1.15-alpine AS builder
# Set working directory
WORKDIR /app
# Copy source code
COPY app/* ./
# Build small static executable
RUN go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /s3listobjects

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /s3listobjects /s3listobjects
# Require to fix problem x509: certificate signed by unknown authority
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Expose port 8080
EXPOSE 8080
# Run the s3listobject binary
ENTRYPOINT ["/s3listobjects"]
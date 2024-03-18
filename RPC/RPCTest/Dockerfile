# Dockerfile in utils/school_scraping directory

# Use the golang image as the base image
FROM golang:1.22 AS builder

# Set the working directory to /canopy/app
WORKDIR /app

# Copy the entire project directory into the image
COPY go.mod .
COPY go.sum .
ADD modules ./modules
ADD RPC ./RPC
ADD grpc ./grpc
ADD errors ./errors
ADD APP/Helper ./APP/Helper
ADD Configuration ./Configuration

# Change the working directory to /canopy/app/utils/school_scraping
WORKDIR /app/RPC/RPCTest

# Display the content of the school_scraping directory
RUN ls /app

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o RPCTest .

############################
# STEP 2 build a small image
############################
FROM alpine:latest

WORKDIR /app

COPY .env /canopy/app/.env

COPY --from=builder /usr/local/go/lib/time/ /usr/local/go/lib/time/

COPY --from=builder /app/RPC/RPCTest /app

RUN ls /app

CMD ["/app/RPCTest"]
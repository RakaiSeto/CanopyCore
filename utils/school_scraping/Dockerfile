# Dockerfile in utils/school_scraping directory

# Use the golang image as the base image
FROM golang:1.22.0-bookworm AS builder

# Set the working directory to /canopy/app
WORKDIR /canopy/app

# Copy the entire project directory into the image
COPY ../../../. .

# Change the working directory to /canopy/app/utils/school_scraping
WORKDIR /canopy/app/utils/school_scraping

# Display the content of the school_scraping directory
RUN ls /canopy/app

# Build the Go application
RUN go build -o school_scraping .

# Set the command to run the executable
CMD ["./school_scraping"]

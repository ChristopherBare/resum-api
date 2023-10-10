# Use the official Amazon Linux image as the base image
FROM amazonlinux:2

# Install required packages
RUN yum -y update && \
    yum -y install glibc-langpack-en && \
    amazon-linux-extras enable golang1.21 && \
    yum -y install golang

# Set the working directory to the root of your project
WORKDIR /app

# Copy your Go code into the container
COPY . /app

# Build the Go code inside the container
RUN go build -o main

# Command to run your application
CMD ["/app/main"]

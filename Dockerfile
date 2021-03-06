# The base go-image
FROM golang:1.15-alpine

# Add Maintainer Info
LABEL maintainer="Joshua Swiss <josh@joshuaswiss.dev>"

# Create a directory for the app
RUN mkdir /bookshelf

# Copy all files from the current directory to the app directory
COPY . /bookshelf

# Set working directory
WORKDIR /bookshelf
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

# Run command as described:
# go build will build an executable file named server in the current directory
# Run the server executable
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o server ." -command="./server"

EXPOSE 3000

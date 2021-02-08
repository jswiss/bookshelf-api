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

# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o server .

EXPOSE 3000

# Run the server executable
CMD ["/bookshelf/server"]

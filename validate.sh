#!/bin/zsh
   # Make sure we're in the project directory within our GOPATH
 #   cd "dynamoDBLoader"
      # Fetch all dependencies
    go get -t ./...
      # Ensure code passes all lint tests
    golangci-lint run
      # Check the Go code for common problems with 'go vet'
    go vet .
      # Run all tests included with our application
    go test .
    cd ..

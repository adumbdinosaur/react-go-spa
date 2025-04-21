#!/bin/bash

# Navigate to the server directory and start the Go server
echo "Starting the Go server..."
cd server || exit
go run cmd/text-app-server/main.go &
SERVER_PID=$!

# Navigate to the React app directory and start the development server
echo "Starting the React application..."
cd ../site/text-app || exit
yarn
yarn dev &
REACT_PID=$!

# Wait for both processes to finish
trap "kill $SERVER_PID $REACT_PID" SIGINT SIGTERM
wait $SERVER_PID $REACT_PID
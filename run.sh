#!/bin/bash

# Install frontend dependencies and start the frontend server
cd frontend
npm install
npm run dev &

# Install backend dependencies and start the backend server
cd ../backend
go mod tidy
go run main.go &

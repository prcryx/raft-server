# Raft-Server

This project is a Go backend application that serves as the server-side component for an exemplary social media platform.

## Features

- Authentication using Twilio and Jwt
- Logout Operation
- Post Feed
- Follow-Followers Mechanism

## Prerequisite

Before running the project, make sure you have the following:

- [Go](https://golang.org/) installed on your machine.

## How to run the app

1. Clone the repository:

   ```sh
   git clone https://github.com/prcryx/raft-server.git
   ```

2. Navigate to the project directory

   ```sh
   cd raft-server
   ```

3. Install dependencies:
   ```sh
   go mod download
   ```
4. Install Air for live reload:
   ```sh
   go install github.com/cosmtrek/air@latest
   ```
5. Install Google's wire for genererating dependency injection codes:
   ```sh
   go install github.com/google/wire/cmd/wire@latest
   ```

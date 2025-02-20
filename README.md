# JWT Authorization in Go

![Go](https://img.shields.io/badge/Language-Go-blue)

A simple JWT authorization system built with Go. Handles user registration, authentication, and protected routes. Testable via Postman without a frontend.

## Features

- **User Registration** with hashed passwords
- **JWT Token Generation**
- **Protected Endpoints**

## Installation

1. **Clone the repo**
   ```bash
   git clone https://github.com/ajarodiy/jwt-go-auth.git
   cd jwt-go-auth
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

## Usage

1. **Set Environment Variables** Create a .env file:
   ```ini
   JWT_KEY=your_secret_key
   PORT=8000
   ```

2. **Run the Server**
   ```bash
   go run main.go
   ```
   Server runs at `http://localhost:8000`.

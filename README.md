# Secure Login System

A simple full-stack authentication project built with React, Go, MySQL, JWT, and the Docker SDK for Go.

## Features

- Email and password login
- Bcrypt password hashing
- JWT-based authentication
- Backend-controlled role authorization
- MySQL user storage
- Admin-triggered Nginx Docker container
- Nginx exposed on `http://localhost:8081`
- React frontend with a simple dashboard

## Project Structure

```text
secure-login-system/
├── backend/
│   ├── cmd/server/       # Application entry point
│   ├── cmd/hashgen/      # Password-hash helper
│   ├── config/           # Environment configuration
│   ├── database/         # MySQL connection
│   ├── docker/           # Docker SDK services
│   ├── handlers/         # HTTP handlers
│   ├── middleware/       # JWT middleware
│   ├── routes/           # API routes
│   └── utils/            # Password and JWT helpers
├── database/
│   └── schema.sql
├── frontend/
│   └── src/
│       ├── pages/
│       ├── App.jsx
│       └── main.jsx
├── .gitignore
└── README.md
```

## Requirements

- Go
- Node.js and npm
- MySQL
- Docker Desktop

## Setup

### 1. Database

Run `database/schema.sql` in MySQL.

### 2. Backend environment

Copy `backend/.env.example` to `backend/.env` and set your local database password and JWT secret.

Do not commit `.env` to GitHub.

### 3. Backend

```bash
cd backend
go mod download
go run ./cmd/server
```

The API runs on `http://localhost:8080`.

### 4. Frontend

```bash
cd frontend
npm install
npm run dev
```

Open the Vite URL shown in the terminal, normally `http://localhost:5173`.

### 5. Docker

Make sure Docker Desktop is running. After an admin logs in, the Go Docker SDK starts the `nginx:latest` container if it is not already running.

Nginx is available at:

```text
http://localhost:8081
```

The service reuses the existing `secure-login-nginx` container instead of creating a new one on every dashboard request.

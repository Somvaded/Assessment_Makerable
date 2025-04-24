**API Documentation link (Postman link)**
https://documenter.getpostman.com/view/28217638/2sB2ixju19


# Patient Management System

This is a backend service for a patient management system designed for a clinic or hospital setup. The system includes authentication and role-based access control for receptionists and doctors to manage patient information securely.

## Features

- **User Authentication**
  - JWT-based login
  - Role-based access control (`receptionist`, `doctor`)

- **Receptionist Portal**
  - Add new patients
  - View patient by Aadhar ID
  - Update patient information
  - Delete patient records

- **Doctor Portal**
  - View all assigned patients
  - Update medical information for a patient

## Tech Stack

- **Golang** (Gin framework)
- **PostgreSQL**
- **Render** (for hosting)
- **JWT Authentication**
- **Clean Code Structure with MVC Principles**

## Project Structure

/cmd -> Main application entry point

/routes -> Route definitions and groupings

/handlers -> HTTP handler logic

/repositories -> Database operations

/models -> Data models

/middlewares -> Authentication & role check middleware

/config -> Configuration loading

/db -> DB connection logic

## Deployment (Render)

- **Build Command**: `go build -o main ./cmd`
- **Start Command**: `./main`
- **Root Directory**: `.`

## Getting Started

1. Clone the repository
2. Create a `.env` file with your `DBURL`, `JWTSecret` and `PORT` (remove PORT if hosting on Render) 
3. Run `go run ./cmd` or deploy to Render

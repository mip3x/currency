# Currency

This repository provides a simple currency information service with several endpoints. Below you can find details about each endpoint, how to run the application (locally or via Docker), and example requests/responses.

---

## Table of Contents

1. [Overview](#overview)
2. [Prerequisites](#prerequisites)
3. [Running Locally](#running-locally)
4. [Running with Docker](#running-with-docker)
5. [Configuration](#configuration)
6. [Endpoints](#endpoints)
   - [Health Check (`/health`)](#health-check-health)
   - [Service Info (`/info`)](#service-info-info)
   - [Currency Info (`/info/currency`)](#currency-info-infocurrency)
7. [Examples](#examples)
8. [License](#license)

---

## Overview

This service provides:
- A **health check** endpoint for monitoring.
- A **service info** endpoint to retrieve metadata (version, service name, author).
- A **currency info** endpoint to fetch currency rates based on a given date and optional currency filter.

---

## Prerequisites

- **Go 1.22** or higher (if building from source).
- **Docker** (if running in containers).
- **docker compose** (v2) if you want to run multiple services easily.

---

## Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/yourname/currency_app.git
   cd currency_app
   ```
2. Ensure you have `Go 1.22+` installed
3. Build and run:
   ```bash
   go mod tidy
   go build -o main main.go
   ./main
   ```
4. By default, the service listens on port 8080 (or reads from the PORT environment variable if configured). You can verify it’s running by opening http://localhost:8080/health in your browser or running:
   ```bash
   curl http://localhost:8080/health
   ```
You should see `OK` if everything is working properly

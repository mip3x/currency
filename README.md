# Currency

This repository provides a simple currency information service with several endpoints. Below you can find details about each endpoint, how to run the application (locally or via Docker), and example requests/responses

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

---

## Overview

This service provides:
- A **health check** endpoint for monitoring
- A **service info** endpoint to retrieve metadata (version, service name, author)
- A **currency info** endpoint to fetch currency rates based on a given date and optional currency filter

---

## Prerequisites

- **Go 1.22** or higher (if building from source)
- **Docker** (if running in containers)
- **docker compose** (v2) if you want to run multiple services easily

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

---

## Running with Docker

1. Run via `docker compose`:
   ```bash
   docker compose up -d
   ```
2. By default, `nginx` listens on port 8081 (or reads from the NGINX_PORT environment variable if configured). You can verify it’s running by opening http://localhost:8081/health in your browser or running:
   ```bash
   curl http://localhost:8081/health
   ```
You should see `OK` if everything is working properly

---

## Configuration

A `config.yml` file can be used for additional settings (e.g., version, service name, author). By default, it might look like this:
```yaml
version: v2
service: currency
author: AUTHOR
```

---

## Endpoints

### Health Check (`/health`)

- **Path**: `/health`
- **Method**: `GET`
- **Description**: Returns a simple string `OK` if the service is running
- **Response**:
  - `200 OK` with body: `OK`

#### Example

```bash
curl http://localhost:8080/health
```
`OK`

### Service info (`/info`)

- **Path**: `/info`
- **Method**: `GET`
- **Description**: Returns JSON with metadata about the service: version, service name, author (read from `config.yml`)
- **Response**:
  - `200 OK` with JSON:
    ```json
    {
     "version": "v2",
     "service": "currency",
     "author": "AUTHOR"
    }
    ```

#### Example

```bash
curl http://localhost:8080/info
```
```json
{
  "version": "v2",
  "service": "currency",
  "author": "mip3x"
}
```

### Currency info (`/info/currency`)

- **Path**: `/info/currency`
- **Method**: `GET`
- **Query Parameters**:
-    `date` (optional): `YYYY-MM-DD` format. If omitted, it may default to a current or fallback date
-    `currency` (optional): Currency code (e.g., `USD`, `EUR`). If provided, returns only the specified currency’s rate
- **Description**: Returns currency rates for a given date. If `currency` is specified, returns only that currency. Otherwise returns all available currencies
- **Response**:
  - `200 OK` with JSON:
    ```json
    {
       "data": {
       "USD": 75.7668,
       "EUR": 86.8894,
       "...": "..."
       },
    "service": "currency"
    }
    ```

#### Example

```bash
curl http://localhost:8080/info/currency?date=2022-01-17
```

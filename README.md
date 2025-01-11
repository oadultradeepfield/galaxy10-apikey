# Galaxy10 DECals Anomaly Detection API Key Generator

![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)
![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)
![MUI](https://img.shields.io/badge/MUI-%230081CB.svg?style=for-the-badge&logo=mui&logoColor=white)
![Vercel](https://img.shields.io/badge/vercel-%23000000.svg?style=for-the-badge&logo=vercel&logoColor=white)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Google Cloud](https://img.shields.io/badge/GoogleCloud-%234285F4.svg?style=for-the-badge&logo=google-cloud&logoColor=white)

This repository contains the web application for generating API keys to access the public API of my previous project, [Galaxy10 Anomaly Detection](https://github.com/oadultradeepfield/galaxy10-anomaly-detection). For detailed context about the project, please refer to the [Galaxy10 Anomaly Detection repository](https://github.com/oadultradeepfield/galaxy10-anomaly-detection).

## Prerequisites

Before proceeding, ensure you have the following installed:

- [Node.js](https://nodejs.org/en)
- [Bun](https://bun.sh/)
- [Docker](https://www.docker.com/)

## Getting Started

This repository is organized into two subdirectories:

- **Frontend**: Built with Vite and uses Bun as the package manager.
- **Backend**: Written in Go, containerized with Docker, and supports live reload using [Air](https://github.com/cosmtrek/air).

### Frontend Setup

To run the frontend locally:

1. Navigate to the `frontend` directory.
2. Install dependencies and start the development server:

   ```bash
   cd frontend
   bun install
   bun run dev
   ```

### Backend Setup

The backend is containerized with Docker. For local development, you can use the provided `Dockerfile.dev`, which includes configurations for live reloading.

1. Build the Docker image using the development Dockerfile:

   ```bash
   docker build -f backend/Dockerfile.dev -t galaxy10-apikey:latest .
   ```

2. Run the container with volume mounting and environment variable support:

   ```bash
   docker run -v $(pwd):/app --env-file .env -d --name galaxy10-apikey -p 8080:8080 galaxy10-apikey:latest
   ```

# 🎵 CollabGo - Collaborative Music Playlist API

CollabGo is a collaborative music playlist API built with Golang. It allows users to create, manage, and share playlists in real-time. The API integrates PostgreSQL (running in a Docker container) and supports live reloading using Air.

---

## Features
- **User Authentication & Authorization**
- **Playlist Management** (Create, update, delete, and share playlists)
- **Real-time Collaboration** (WebSocket-based updates)
- **Song Voting System** (Upvote/downvote songs in playlists)
- **PostgreSQL Database** (Managed with Docker Compose)
- **Hot Reloading** (Using Air for live code updates)
- **Caching** (Redis for faster playlist access)

---

## Prerequisites
- **Golang** (>= 1.18)
- **Docker & Docker Compose**
- **Air** (for hot reloading)

---

## Installation & Setup

### 1. Clone the Repository
```bash
git clone https://github.com/JerryLegend254/CollabGo.git
cd CollabGo
```

### 2. Install Air for Hot Reloading
```bash
go install github.com/cosmtrek/air@latest
```

### 3. Start the Database (PostgreSQL via Docker)
```bash
docker-compose up -d
```

### 4. Run Database Migrations (Make sure you have migrate installed)
```bash
make migrate-up
```

### 5. Start the API Server with Live Reloading
```bash
air
```

---

## Environment Variables (`.envrc`)

```env
ADDR=":8080"
DB_ADDR="postgres://admin:adminpassword@localhost/collabgo?sslmode=disable"
DB_MAX_OPEN_CONNS=30
DB_MAX_IDLE_CONNS=30
DB_MAX_IDLE_TIMEOUT="15m"
```

---

## Contact
For any issues or feature requests, please open an issue on GitHub.

---

Happy coding! 🎵
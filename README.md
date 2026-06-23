# Rhoam Together

A collaborative travel planning app where users can create trips, invite friends, manage group access, and share real-time suggestions for things to do.

## Features (MVP)

- User authentication (email/password signup and login)
- Create and manage travel trips with dates
- Invite friends to trips with role-based access (read-only, editor, admin)
- Add suggestions and plans for trip activities
- Real-time updates via WebSockets when suggestions are added
- Responsive web interface

## Tech Stack

- **Backend**: Go (Gorilla Mux, WebSockets)
- **Frontend**: React with React Router
- **Database**: PostgreSQL
- **Deployment**: Railway/Render (backend + DB), Vercel (frontend)

## Project Structure

```
rhoam-together/
├── backend/          # Go API server
├── frontend/         # React web application
├── docs/            # Documentation
├── configs/         # Configuration files
└── README.md
```

## Getting Started

### Prerequisites
- Go 1.26+
- Node.js 18+
- PostgreSQL 14+

### Backend Setup
```bash
cd backend
go mod init rhoam-together
# Install dependencies (coming next)
```

### Frontend Setup
```bash
cd frontend
npm create vite@latest . -- --template react
npm install
```

## Environment Variables

Create `.env` files in both `backend/` and `frontend/` directories (see Phase 2 and 5 docs for details).

## Development

This application uses make commands to start the development servers.

- `make start`: Start both backend and frontend dev servers
- `make dev`: Start both backend and frontend dev servers
- `make backend`: Start only backend (Go on http://localhost:8080)
- `make frontend`: Start only frontend (React on http://localhost:3000)
- `make install`: Install dependencies for both backend and frontend
- `make backend-install`: Install Go dependencies
- `make frontend-install`: Install Node dependencies
- `make stop`: Stop all development servers

## Deployment

See Phase 7 documentation for deployment to Railway/Render and Vercel.

## License

See LICENSE file

.PHONY: start backend frontend install help dev stop

help:
	@echo "Rhoam Together - Development Commands"
	@echo "======================================"
	@echo "make start       - Start both backend and frontend dev servers"
	@echo "make dev         - Alias for 'make start'"
	@echo "make backend     - Start only backend (Go on :8080)"
	@echo "make frontend    - Start only frontend (React on :3000)"
	@echo "make install     - Install dependencies for both backend and frontend"
	@echo "make backend-install - Install Go dependencies"
	@echo "make frontend-install - Install Node dependencies"
	@echo "make stop        - Stop all development servers (if running in background)"
	@echo ""
	@echo "Frontend will be available at: http://localhost:3000"
	@echo "Backend API will be available at: http://localhost:8080"

# Default target
.DEFAULT_GOAL := help

# Start both services in parallel
start:
	@echo "Starting Rhoam Together development environment..."
	@echo ""
	@echo "Backend will start on: http://localhost:8080"
	@echo "Frontend will start on: http://localhost:3000"
	@echo ""
	@echo "Press Ctrl+C to stop both servers"
	@echo ""
	@(cd backend && go run main.go) & \
	(cd frontend && npm run dev) & \
	wait

# Alias for start
dev: start

# Start backend only
backend:
	@echo "Starting backend server on port 8080..."
	cd backend && go run main.go

# Start frontend only
frontend:
	@echo "Starting frontend dev server on port 3000..."
	cd frontend && npm run dev

# Install all dependencies
install: backend-install frontend-install
	@echo "✓ All dependencies installed"

# Install backend dependencies
backend-install:
	@echo "Installing backend dependencies..."
	cd backend && go mod download
	cd backend && go mod tidy

# Install frontend dependencies
frontend-install:
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

# Stop all development servers (for use when running in background)
stop:
	@echo "Stopping development servers..."
	@pkill -f "go run main.go" 2>/dev/null || true
	@pkill -f "vite" 2>/dev/null || true
	@echo "✓ Development servers stopped"

# Development helper - install and start
setup: install
	@echo ""
	@echo "Setup complete! Run 'make start' to begin development."
	@echo ""

# Build for production
build: frontend-build backend-build
	@echo "✓ Build complete"

frontend-build:
	@echo "Building frontend..."
	cd frontend && npm run build

backend-build:
	@echo "Building backend..."
	cd backend && go build -o rhoam-together-api

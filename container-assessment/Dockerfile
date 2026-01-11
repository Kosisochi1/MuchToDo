# ============================
# Stage 1 — Build the Go app
# ============================
FROM golang:1.25-alpine AS builder



# Working directory inside the build container
WORKDIR /build

# Copy module files first for caching
COPY application-code/go.mod application-code/go.sum ./
RUN go mod download

# Copy the full Go source code
COPY ./application-code .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/api/main.go


# ============================
# Stage 2 — Minimal runtime
# ============================
FROM alpine:3.19

# Enable HTTPS for external calls
RUN apk add --no-cache curl ca-certificates
# Create non-root user
RUN addgroup -S appuser && adduser -S appuser -G appuser

WORKDIR /app



# Copy compiled binary from builder stage
COPY --from=builder /build/server .
#COPY .env .env


# Add venv to PATH
# ENV PATH="/app/.venv/bin:$PATH"


RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser


# Expose backend port (optional)
EXPOSE 8080


HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

# Run the Go application
CMD ["./server"]

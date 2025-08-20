# --- STAGE 1: Budowanie aplikacji ---
FROM golang:1.24 AS builder

# Ustaw katalog roboczy w kontenerze
WORKDIR /app

# Skopiuj pliki modów i pobierz zależności
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Skopiuj resztę kodu
COPY backend/ .

# Zbuduj binarkę z katalogu cmd/api
RUN go build -o main ./cmd/api

# --- STAGE 2: Minimalny obraz uruchomieniowy ---
FROM debian:12-slim

WORKDIR /app
COPY --from=builder /app/main .

# Otwórz port dla API
EXPOSE 8080

# Domyślne polecenie
CMD ["./main"]

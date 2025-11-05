# Build Step

FROM golang:1.24 AS build

RUN apt-get update && apt-get install -y --no-install-recommends \
    nodejs \
    npm \
    make \
    ca-certificates \
    libolm-dev \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY . .
RUN make core_build

# Runtime Step
FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    libolm3 \
    && rm -rf /var/lib/apt/lists/*
COPY --from=build /app/dist/core /app/core
WORKDIR /app
CMD ["/app/core"]
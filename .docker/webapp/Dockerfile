# Build Step
FROM arm64v8/golang:1.23 AS build

RUN apt-get update && apt-get install -y --no-install-recommends \
    nodejs \
    npm \
    make \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY . .
RUN make core_build

# Runtime Step
FROM arm64v8/debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*
COPY --from=build /app/dist/core /app/core
WORKDIR /app
CMD ["/app/core"]
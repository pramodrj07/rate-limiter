FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive
ENV GO_VERSION=1.22.3
ENV CGO_ENABLED=0

# Install base packages
RUN apt-get update && apt-get install -y \
    curl \
    git \
    build-essential \
    ca-certificates \
    sudo \
    && rm -rf /var/lib/apt/lists/*

# Install Go (multi-arch safe)
RUN ARCH=$(dpkg --print-architecture) && \
    curl -LO https://go.dev/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-${ARCH}.tar.gz && \
    rm go${GO_VERSION}.linux-${ARCH}.tar.gz

# Create vscode user
RUN useradd -ms /bin/bash vscode && \
    echo "vscode ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

# Setup Go environment
ENV GOPATH=/home/vscode/go
ENV PATH=/usr/local/go/bin:/home/vscode/go/bin:$PATH

# Create GOPATH directories BEFORE switching user
RUN mkdir -p /home/vscode/go/bin && \
    chown -R vscode:vscode /home/vscode

USER vscode
WORKDIR /workspace

# Install Go tools
RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install github.com/air-verse/air@latest

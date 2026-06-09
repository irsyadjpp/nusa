# Podman Setup Guide

This guide provides instructions for setting up and using Podman with the NUSA project as an alternative to Docker.

## What is Podman?

Podman is a daemonless container engine for developing, managing, and running OCI Containers on your Linux System. It's compatible with Docker images and Dockerfiles, making it a drop-in replacement for Docker in many scenarios.

## Installation

### Linux (Ubuntu/Debian)

```bash
# Update package index
sudo apt-get update

# Install Podman
sudo apt-get install -y podman

# Install podman-compose (Python-based docker-compose alternative)
pip3 install podman-compose
```

### Linux (Fedora/RHEL)

```bash
# Podman is usually pre-installed on Fedora/RHEL
sudo dnf install podman

# Install podman-compose
pip3 install podman-compose
```

### Linux (Arch Linux)

```bash
sudo pacman -S podman
pip3 install podman-compose
```

### macOS

```bash
# Install Homebrew if not already installed
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Install Podman
brew install podman

# Install podman-compose
pip3 install podman-compose

# Initialize Podman machine
podman machine init
podman machine start
```

### Windows

```bash
# Install Podman Desktop from https://podman.io/getting-started/installing
# Podman Desktop includes podman-compose support
```

## Verification

After installation, verify Podman is working:

```bash
# Check Podman version
podman --version

# Check podman-compose version
podman-compose --version

# Test with a simple container
podman run hello-world
```

## Podman vs Docker Compatibility

Podman is largely compatible with Docker:

- **Dockerfiles**: Podman can use Dockerfiles without modification
- **Images**: Podman can pull and run Docker images from Docker Hub
- **Docker Compose**: Use `podman-compose` as a drop-in replacement for `docker-compose`
- **Commands**: Most Docker commands have Podman equivalents (e.g., `podman run` instead of `docker run`)

## Key Differences

1. **Daemonless**: Podman doesn't require a background daemon process
2. **Rootless**: Podman can run containers without root privileges (recommended for security)
3. **Systemd Integration**: Podman integrates better with systemd for service management
4. **Pods**: Podman has native support for Kubernetes-like pods

## Common Issues

### Permission Denied Errors

If you encounter permission errors, you may need to configure user namespaces:

```bash
# Enable user namespaces (Linux)
echo "user.max_user_namespaces=15000" | sudo tee -a /etc/sysctl.conf
sudo sysctl -p
```

### Network Issues

Podman uses a different networking model. If you encounter network issues:

```bash
# Check Podman network configuration
podman network ls

# Create a custom network if needed
podman network create nusa_network
```

### Volume Permissions

For rootless Podman, ensure your user has proper permissions on volume directories:

```bash
# Fix volume permissions
sudo chown -R $USER:$USER ./data
```

## Next Steps

After installing Podman:

1. Configure your environment variables in `.env` file
2. Use `make podman-up` to start services with Podman
3. Use `make podman-logs` to view logs
4. Use `make podman-down` to stop services

For more detailed usage instructions, see [PODMAN_GUIDE.md](PODMAN_GUIDE.md).

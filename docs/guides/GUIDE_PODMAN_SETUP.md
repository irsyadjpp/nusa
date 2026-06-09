# Podman Usage Guide

This comprehensive guide covers using Podman with the NUSA project, including command comparisons, troubleshooting, and best practices.

## Table of Contents

- [Quick Start](#quick-start)
- [Command Comparison](#command-comparison)
- [Working with Containers](#working-with-containers)
- [Working with Pods](#working-with-pods)
- [Volumes and Networks](#volumes-and-networks)
- [Troubleshooting](#troubleshooting)
- [Migration Tips](#migration-tips)

## Quick Start

### Starting Services

```bash
# Using Podman Compose (recommended)
cd backend
make podman-up

# Or directly with podman-compose
podman-compose -f podman-compose.yml up -d
```

### Viewing Logs

```bash
# Using Makefile
make podman-logs

# Or directly
podman-compose logs -f
```

### Stopping Services

```bash
# Using Makefile
make podman-down

# Or directly
podman-compose down
```

## Command Comparison

### Basic Container Operations

| Docker Command | Podman Equivalent | Notes |
|---------------|-------------------|-------|
| `docker run` | `podman run` | Same syntax |
| `docker ps` | `podman ps` | Same syntax |
| `docker images` | `podman images` | Same syntax |
| `docker build` | `podman build` | Same syntax |
| `docker exec` | `podman exec` | Same syntax |
| `docker logs` | `podman logs` | Same syntax |
| `docker stop` | `podman stop` | Same syntax |
| `docker rm` | `podman rm` | Same syntax |
| `docker rmi` | `podman rmi` | Same syntax |

### Compose Operations

| Docker Command | Podman Equivalent | Notes |
|---------------|-------------------|-------|
| `docker compose up` | `podman-compose up` | Requires podman-compose |
| `docker compose down` | `podman-compose down` | Requires podman-compose |
| `docker compose logs` | `podman-compose logs` | Requires podman-compose |
| `docker compose ps` | `podman-compose ps` | Requires podman-compose |

### Volume Operations

| Docker Command | Podman Equivalent | Notes |
|---------------|-------------------|-------|
| `docker volume ls` | `podman volume ls` | Same syntax |
| `docker volume create` | `podman volume create` | Same syntax |
| `docker volume rm` | `podman volume rm` | Same syntax |

### Network Operations

| Docker Command | Podman Equivalent | Notes |
|---------------|-------------------|-------|
| `docker network ls` | `podman network ls` | Same syntax |
| `docker network create` | `podman network create` | Same syntax |
| `docker network rm` | `podman network rm` | Same syntax |

## Working with Containers

### Running Containers

```bash
# Run a container
podman run -d --name my-container nginx

# Run with environment variables
podman run -d --name my-container -e ENV_VAR=value nginx

# Run with volume mounts
podman run -d --name my-container -v /host/path:/container/path nginx

# Run with port mapping
podman run -d --name my-container -p 8080:80 nginx
```

### Inspecting Containers

```bash
# View container details
podman inspect my-container

# View container stats
podman stats my-container

# View container processes
podman top my-container
```

### Executing Commands in Containers

```bash
# Execute a command in a running container
podman exec -it my-container /bin/bash

# Execute a single command
podman exec my-container ls -la
```

## Working with Pods

Podman has native support for pods (similar to Kubernetes pods).

### Creating a Pod

```bash
# Create a pod
podman pod create --name my-pod -p 8080:80

# Add a container to the pod
podman run -d --pod my-pod --name my-container nginx
```

### Managing Pods

```bash
# List pods
podman pod ls

# Inspect a pod
podman pod inspect my-pod

# Stop a pod
podman pod stop my-pod

# Start a pod
podman pod start my-pod

# Remove a pod
podman pod rm my-pod
```

## Volumes and Networks

### Managing Volumes

```bash
# List volumes
podman volume ls

# Create a volume
podman volume create my-volume

# Inspect a volume
podman volume inspect my-volume

# Remove a volume
podman volume rm my-volume
```

### Managing Networks

```bash
# List networks
podman network ls

# Create a network
podman network create my-network

# Connect a container to a network
podman network connect my-network my-container

# Disconnect a container from a network
podman network disconnect my-network my-container

# Remove a network
podman network rm my-network
```

## Troubleshooting

### Permission Denied Errors

If you encounter permission errors with rootless Podman:

```bash
# Check if user namespaces are enabled
sysctl user.max_user_namespaces

# Enable user namespaces (requires root)
echo "user.max_user_namespaces=15000" | sudo tee -a /etc/sysctl.conf
sudo sysctl -p

# Check your user's subuid and subgid mappings
cat /etc/subuid
cat /etc/subgid
```

### Network Issues

If containers cannot communicate:

```bash
# Check Podman network configuration
podman network ls

# Check container network settings
podman inspect my-container | grep -A 20 NetworkSettings

# Create a custom network
podman network create nusa_network

# Restart containers with the new network
podman-compose down
podman-compose up -d
```

### Volume Permission Issues

For rootless Podman, ensure proper permissions:

```bash
# Fix volume permissions
sudo chown -R $USER:$USER ./data

# Or use the :z flag for SELinux
podman run -v ./data:/data:z my-container
```

### Port Already in Use

If you encounter port conflicts:

```bash
# Check what's using the port
sudo lsof -i :5432

# Stop the conflicting service
sudo systemctl stop postgresql

# Or use a different port in your .env file
```

### Container Not Starting

If a container fails to start:

```bash
# Check container logs
podman logs my-container

# Check container status
podman ps -a

# Inspect the container
podman inspect my-container

# Try running in foreground to see errors
podman run --rm my-container
```

## Migration Tips

### From Docker to Podman

1. **Install Podman**: Follow the installation guide in [PODMAN_SETUP.md](PODMAN_SETUP.md)

2. **Install podman-compose**: Required for docker-compose compatibility
   ```bash
   pip3 install podman-compose
   ```

3. **Use existing Dockerfiles**: Podman is compatible with Dockerfiles, no changes needed

4. **Use podman-compose.yml**: Copy your docker-compose.yml to podman-compose.yml (already done in this project)

5. **Update Makefile**: Add Podman targets (already done in this project)

6. **Test thoroughly**: Ensure all services work correctly with Podman

### Common Migration Issues

1. **Socket Path**: Docker uses `/var/run/docker.sock`, Podman uses `/run/user/UID/podman/podman.sock`
   - Most applications don't need the socket, but if yours does, update the path

2. **Network Names**: Podman may use different network names
   - Check with `podman network ls`

3. **Volume Paths**: Rootless Podman stores volumes in `~/.local/share/containers/storage/volumes`
   - Root Podman uses `/var/lib/containers/storage/volumes`

4. **Image Cache**: Podman has its own image cache
   - Images won't be shared between Docker and Podman

### Best Practices

1. **Use Rootless Podman**: More secure and doesn't require root privileges
2. **Enable User Namespaces**: Provides additional isolation
3. **Use Podman Pods**: Better resource management and networking
4. **Regularly Update**: Keep Podman and podman-compose updated
5. **Monitor Resources**: Use `podman stats` to monitor container resource usage
6. **Clean Up Regularly**: Remove unused images and volumes with `podman system prune`

## Additional Resources

- [Official Podman Documentation](https://docs.podman.io/)
- [Podman Compose Repository](https://github.com/containers/podman-compose)
- [Podman vs Docker Comparison](https://developers.redhat.com/articles/2022/01/27/podman-vs-docker)

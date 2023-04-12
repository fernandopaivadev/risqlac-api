echo "==> stop all containers"
podman stop -a

echo "==> prune containers"
podman container prune -f

echo "==> prune images"
podman image prune -f

# setup
up:
    docker compose up -d

# Do docker compose down
down: 
    docker compose down

# Build docker image to local development
build-local: 
    docker compose build --no-cache
 
# clean all
barth:
    docker compose down --rmi all --volumes --remove-orphans

# Tail docker compose logs
logs:
    docker compose logs -f

# Check container status
ps:
    docker compose ps

# Execute tests
test:
    go test -race -shuffle=on ./...


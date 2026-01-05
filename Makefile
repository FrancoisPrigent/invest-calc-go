# Variables
IMAGE_NAME=investment-calc
CONTAINER_NAME=investment-runner

.PHONY: build run clean

# Build the Docker image
build:
	@@DOCKER_BUILDKIT=1 docker build -q -t $(IMAGE_NAME) . > /dev/null

# Run the container and pass through flags
# Usage: make run ARGS="-p 20000 -r 7.5 -t 15 -i"
run:
	@docker run --rm --name $(CONTAINER_NAME) $(IMAGE_NAME) $(ARGS)

# Remove the image
clean:
	@docker rmi $(IMAGE_NAME) > /dev/null

# This will Build, then Run, then Cleanup the image.
calc: build run clean

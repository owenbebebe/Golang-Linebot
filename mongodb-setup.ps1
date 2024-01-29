# environemnt: powershell Windows 11 Pro
# type powershell -ExecutionPolicy Bypass -File mongodb-setup.ps1. to run locally
$IMAGE_NAME = "mongo"
$IMAGE_VERSION = "5.0"
$CONTAINER_NAME = "mongodbLine"

#pull mongo image from docker hub
docker pull "${IMAGE_NAME}:${IMAGE_VERSION}"

#run mongo image
docker run -d --name="${CONTAINER_NAME}" "${IMAGE_NAME}:${IMAGE_VERSION}"

#check if container is running
docker logs "${CONTAINER_NAME}"


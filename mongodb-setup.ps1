# environemnt: powershell Windows 11 Pro
# type powershell -ExecutionPolicy Bypass -File mongodb-setup.ps1. to run locally
$IMAGE_NAME = "mongo"
$IMAGE_VERSION = "5.0"
$CONTAINER_NAME = "mongodbLine"

#pull mongo image from docker hub
docker pull "${IMAGE_NAME}:${IMAGE_VERSION}"

#run mongo image
docker run --name="${CONTAINER_NAME}" -d -p 27017:27017 "${IMAGE_NAME}:${IMAGE_VERSION}"

#check if container is running
docker logs "${CONTAINER_NAME}"


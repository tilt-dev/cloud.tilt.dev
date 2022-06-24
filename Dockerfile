FROM alpine
WORKDIR /app
ADD build build
ENTRYPOINT ./build/cloud-tilt-dev

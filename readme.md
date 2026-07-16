# Sylvie
A minimal [Video-on-Demand](https://en.wikipedia.org/wiki/Video_on_demand) system complete with media transcoding, thumbnail generator, HLS support, and hover previews with WebVTT. Media transcoding, thumbnail generation, and HLS support is powered by `ffmpeg`.

This is a minimal system and I tried to follow a microservice architecture as correctly as I can. The major components of the system is separated into core services and communicate through a message queue using [RabbitMQ](https://www.rabbitmq.com).

## Why the name?
`Sylvie` is a reference to the Marvel Cinematic Universe Loki variant of the same name. Though she did not cause the multiverse, killing He Who Remains lead to branching of the Sacred Timeline into parallel universes. Creation of multiple versions of an uploaded media file is a recall to the creation of parallel universes and multiple variants of characters.

## Motivation
This project is a learning adventure to microservice architecture and video streaming systems.

## What's Next?
This checklist is a living document and may evolve over time as the projet progresses and requirements become clearer.

- [x] Implement a video upload service.
- [ ] Implment auth service to track.
- [ ] Implement search service.
- [x] Implement playback service.
- [ ] Implement a frontend to consume all these services.
- [ ] Generate Sprite sheet or WebVTT output for hover previous on video track
- [x] Generate HLS Playlist to support adaptive streaming
- [x] Read app configuration from a config file (sylvieconfig)
- [ ] Handle use-case when src file is in the cloud and transcoding output should be on the cloud
- [ ] Dockerize the program for easier setup

## Shared Directory
To limit the scope of the project, it currently does not integrate with cloud storage and CDNs. Instead, a shared directory in the root of the project is utilized.

## Prerequisite
1. `ffmpeg`
- On Linux
    - **Ubuntu/Debian:** `sudo apt install ffmpeg`

    - **Fedore/RHEL:** `sudo dnf install ffmpeg`

- On MacOS
`brew install ffmpeg`

- On Windows
    - Visit the [Official FFmpeg Website](https://www.ffmpeg.org/download.html) and follow the windows installtion guide.

2. `RabbitMQ`
- The preferred method is running it via `docker`.
- In this project, I used the image `rabbitmq:management-alpine`

## Environment Variables
A `.env` file should exist in the root of the project. The project will run even if this file is empty as long as it exists as there are default values for these varables. See `.env.example` for the default values.

To update the default values, `.env` file should have these variables:
```
PORT=

JSON_DB_FILE_PATH=
STORAGE_DIR=

RABBIT_CONNECTION_STR=
RABBIT_QUEUE_NAME=

FFMPEG_PATH=
FFPROBE_PATH=
```

## Running with Docker

### Prerequisites
- Docker
- Docker compose

### Start the application
Build the images and start all services.

```bash
docker compose up -d --build
```

This command starts:
- **API** - HTTP Server
- **Processor** - Background worker that processes transcoding jobs
- **RabbitMQ** - Message broker (with the management UI)

The API will be available at:
```
http://localhost:3000
```

RabbitMQ management UI:
```
http://localhost:15672
```

Default credentials:
- Username: `guest`
- Password: `guest`

### Stopping the application
```bash
docker compose down
```

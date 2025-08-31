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
- [ ] Implement playback service.
- [ ] Implement a frontend to consume all these services.
- [ ] Generate Sprite sheet or WebVTT output for hover previous on video track
- [x] Generate HLS Playlist to support adaptive streaming
- [x] Read app configuration from a config file (sylvieconfig)
- [ ] Handle use-case when src file is in the cloud and transcoding output should be on the cloud
- [ ] Dockerize the program for easier setup

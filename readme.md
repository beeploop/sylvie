## Sylvie

A minimal video transcoding and thumbnail generation program utilizing the `ffmpeg` capabilities.

## Motivation
This is a part of the whole learning process of creating a [Video-on-Demand](https://en.wikipedia.org/wiki/Video_on_demand) system similar what YouTube does.

This is the first part of the project and the goals of this is to subscribe to a message queue for uploaded videos and perform transcoding and thumbnail generation on it.

## What's Next?

- [ ] Generate Sprite sheet or WebVTT output for hover previous on video track
- [ ] Generate HLS Playlist to support adaptive streaming
- [ ] Read app configuration from a config file (sylvieconfig)
- [ ] Generate sprite_sheet/tile for the hover previous
- [ ] Dockerize the program for easier setup

## Requirements
- `ffmpeg`

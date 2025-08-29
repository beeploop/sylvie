## Sylvie
A minimal video transcoding and thumbnail generation program utilizing the `ffmpeg` capabilities. The program subscribe to a [RabbitMQ](https://www.rabbitmq.com) Queue for jobs. See the [message](#rabbitmq-message) section for the message structure.

## Why the name?
`Sylvie` is a reference to the Marvel Cinematic Universe Loki variant of the same name. Though she did not cause the multiverse, killing He Who Remains lead to branching of the Sacred Timeline into parallel universes. Creation of multiple versions of an uploaded media file is a recall to the creation of parallel universes and multiple variants of characters.

## Motivation
This is a part of the whole learning process of creating a [Video-on-Demand](https://en.wikipedia.org/wiki/Video_on_demand) system similar what YouTube does.

This is the first part of a bigger project and the goals of this is to subscribe to a message queue for uploaded videos and perform transcoding and thumbnail generation on it.

## What's Next?

- [ ] Generate Sprite sheet or WebVTT output for hover previous on video track
- [x] Generate HLS Playlist to support adaptive streaming
- [x] Read app configuration from a config file (sylvieconfig)
- [ ] Handle use-case when src file is in the cloud and transcoding output should be on the cloud
- [ ] Dockerize the program for easier setup

## Requirements
- `ffmpeg`
- `RabbitMQ`

## Configuration File
The tool automatically looks for a `sylvieconfig.yaml` file in the root of the directory.

### Example config
```yaml
out_dir: ./path/to/desired/directory
rabbit_connection_string: "amqp://guest:guest@localhost:5672"
transcoding_queue_name: "transcoding_jobs_queue"
publishing_queue_name: "transcoding_output_queue"
```
 - `out_dir` : directory where the transcoded media and result metadata is saved.
 - `rabbit_connection_string` : connection string to RabbitMQ.
 - `transcoding_queue_name` : queue name in RabbitMQ where sylvie will pull jobs from.
 - `publishing_queue_name` : queue name in RabbitMQ where sylvie will publish output to when done.

### Running with default config filename
```bash
sylvie
```

### Specifying a configuration file
```bash
sylvie --config {your-config-file.yaml}
```

## RabbitMQ Message
Sylvie subscribe to a RabbitMQ Queue for transcoding jobs. The job is in JSON format and structure looks like this:

```json
{
    "video_id": "af57f595-e548-421b-9410-7f50f719c7b3",
    "path": "/path/to/the/video/file",
    "resolutions": ["1080p", "720p", "480p"]
}
```
 - `video_id` - Unique ID of the client-uploaded video. This comes from the uploading service.
 - `path` - A sylvie-accessible path to the video file to be transcoded.
 - `resolutions` - List of resolutions to transcode to. Accepted values are: "1080p", "720p", "480p", "360p", "240p", "144p".

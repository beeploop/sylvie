## Media Transcoder
This transcoder subscribe to a [RabbitMQ](https://www.rabbitmq.com) message queue to pull jobs. I call these jobs `message` and this is a JSON formatted data that looks like this:

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

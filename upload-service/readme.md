## Upload Service
This service handles video uploads and stores it to the directory inside the `shared directory`.

## Configuration
This accepts/reads an `env` file with  the following structure: 

```env
PORT=3000
STORAGE_PATH=/path/to/upload/dir/in/resource
RABBIT_URL="amqp://localhost"
PUBLISH_QUEUE="transcoding_queue"
```

 - `PORT` : port the service should listen on.
 - `STORAGE_PATH` : path to the dedicated directory for file uploads inside the `shared directory`.
 - `RABBIT_URL` : connection url to RabbitMQ
 - `PUBLISH_QUEUE` : queue name in RabbitMQ where the service publishes jobs to. It should be identical to [transcoding_queue_name](../media-transcoder/readme.md#example-config).

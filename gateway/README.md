# Gateway

Handles incoming HTTP requests


| HTTP            | gRPC              |
| --------------- | ----------------- |
| `POST`-> `/shorten` | `CreateShortURL`  |
| `GET` -> `/{code}`   | `ResolveShortURL` |


Translate gRPC errors into HTTP errors

| gRPC Code       | HTTP Status |
| --------------- | ----------- |
| InvalidArgument | 400         |
| NotFound        | 404         |
| AlreadyExists   | 409         |
| Unavailable     | 503         |
| Internal        | 500         |


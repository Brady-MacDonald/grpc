# Gateway

Handles incoming HTTP requests


| HTTP            | gRPC              |
| --------------- | ----------------- |
| `POST`-> `/shorten` | `CreateShortURL`  |
| `GET` -> `/{code}`   | `ResolveShortURL` |


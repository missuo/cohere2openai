# Cohere2OpenAI
Used to convert the Cohere API to OpenAI compatible API. **Easily use Cohere with any OpenAI compatible client.**

## Before Using
You need to have a Cohere API key, if you don't have one, you can apply for [Trial Key](https://dashboard.cohere.com/api-keys). It is completely free at the moment and will not charge you any fees. You don't even need to bind your credit card.

## Demo (Public Convert API)
This is the public API I provide. I cannot guarantee stability, but you can use them for free without any deployment.

```bash
# Host: c.uid.si
# Endpoint: /v1/chat/completions
# Method: POST
# Headers: Content-Type: application/json, Authorization
curl https://c.uid.si/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer xxxxxxx" \
  -d '{
    "model": "command-r-plus",
    "messages": [
      {
        "role": "system",
        "content": "翻译为中文!"
      },
      {
        "role": "user",
        "content": "Hello!"
      }
    ],
    "stream": true
  }'
```

## Compatibility
Currently it is only compatible with the **Command** family of models, if you pass in any other model, the default will be to use **Command R+**. Supports streaming and non-streaming output.

### Models
```json
{
  "created": 1692901427,
  "id": "command-r",
  "object": "model",
  "owned_by": "system"
},
{
  "created": 1692901427,
  "id": "command-r-plus",
  "object": "model",
  "owned_by": "system"
},
{
  "created": 1692901427,
  "id": "command-light",
  "object": "model",
  "owned_by": "system"
},
{
  "created": 1692901427,
  "id": "command-light-nightly",
  "object": "model",
  "owned_by": "system"
},
{
  "created": 1692901427,
  "id": "command",
  "object": "model",
  "owned_by": "system"
},
{
  "created": 1692901427,
  "id": "command-nightly",
  "object": "model",
  "owned_by": "system"
}
```

## Request Example
```bash
curl http://127.0.0.1:6600/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer xxxxxxx" \
  -d '{
    "model": "command-r-plus",
    "messages": [
      {
        "role": "system",
        "content": "翻译为中文!"
      },
      {
        "role": "user",
        "content": "Hello!"
      }
    ],
    "stream": true
  }'
```


## Usage
### Docker

```bash
docker run -d --restart always -p 6600:6600 ghcr.io/missuo/cohere2openai:latest
```

```bash
docker run -d --restart always -p 6600:6600 missuo/cohere2openai:latest
```

### Docker Compose
It is recommended that you use docker version **26.0.0** or higher, otherwise you need to specify the version in the `compose.yaml` file.
```diff
+version: "3.9"
```

```bash
mkdir cohere2openai && cd cohere2openai
wget -O compose.yaml https://raw.githubusercontent.com/missuo/cohere2openai/main/compose.yaml
docker compose up -d
```

### Manual

Download the latest release from the [release page](https://github.com/missuo/cohere2openai/releases).

```bash
chmod +x cohere2openai
./cohere2openai
```

## License
[GPL 3.0](https://github.com/missuo/cohere2openai/blob/main/LICENSE)

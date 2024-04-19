# Cohere2OpenAI
Used to convert the Cohere API to OpenAI compatible API. **Easily use Cohere with any OpenAI compatible client.**

## Compatibility
Currently it is only compatible with the Cohere family of models, if you pass in any other model, the default will be to use **Command R+**. 

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

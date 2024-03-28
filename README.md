![Gateway](/.github/gateway.png)

<div align="center">

### Core infrastructure stack for building production-ready AI Applications

![Lint](https://github.com/missingstudio/gateway/actions/workflows/lint.yml/badge.svg)
![Release](https://github.com/missingstudio/gateway/actions/workflows/release.yml/badge.svg)
[![License](https://img.shields.io/github/license/missingstudio/gateway)](./LICENSE)
[![Version](https://img.shields.io/github/v/release/missingstudio/gateway?logo=semantic-release)](Version)
[![Discord](https://img.shields.io/discord/1136647481128661082)](https://discord.gg/yxV58ydzV7)
[![Twitter](https://img.shields.io/twitter/follow/_missingstudio)](https://twitter.com/_missingstudio)
</div>
<br><br>

## Introduction
🌈 A Robust cloud-native [AI Gateway](https://www.missing.studio) - core LLMOps infrastructure stack for building production-ready AI Applications. It provides an Universal API for inferencing 100+ LLMs(OpenAI, Azure, Cohere, Anthropic, HuggingFace, Replicate, Stable Diffusion).

## 🚀 Key Features

✅&nbsp; Seamless Integration with **Universal API** <br>
✅&nbsp; Reliable LLM Routing with **AI Router** <br>
✅&nbsp; **Load balance** across multiple models and providers <br>
✅&nbsp; **Automatic Retries** with exponential fallbacks <br>
✅&nbsp; High availability and resiliency using **production-ready LLMOps** <br>
🚧&nbsp; Detailed **Usage Analytics** <br>
🚧&nbsp; **PII detection and masking**<br>
🚧&nbsp; **Simple & Semantic Caching** for cost reduction <br>
🚧&nbsp; **No Vendor lock-in Observability** - Logging, monitoring and tracing <br>
✅&nbsp; **Enterprise-ready** with enhanced security, reliability, and scale with custom deployments support.

## Supported Providers
|                                               |   Provider     |   Provider Name  |   Support   |   Supported Endpoints    |
|-----------------------------------------------|----------------|    :---:         |    :---:    |--------------------------|
|<img src="assets/openai.png" width=16>         | OpenAI         |     openai       |      ✅     |  `/chat/completions`, `/chat/completions:stream`     |
|<img src="assets/groq.svg" width=16>           | Groq           |     groq         |      ✅     |  `/chat/completions`, `/chat/completions:stream`     |
|<img src="assets/anyscale.png" width=16>       | Anyscale       |    anyscale      |      ✅     |  `/chat/completions`     |
|<img src="assets/deepinfra.jpeg" width=16>     | Deepinfra      |    deepinfra     |      ✅     |  `/chat/completions`     |
|<img src="assets/togetherai.svg" width=16>     | Together AI	   |    togetherai    |      ✅     |  `/chat/completions`     |

> Not supported (yet): images, audio, files, fine-tunes, moderations


## Installation
AI gateway can be intall on macOS, Windows, Linux, OpenBSD, FreeBSD, and on any machine

### Binary (Cross-platform)

Download the appropriate version for your platform from [releases](https://github.com/missingstudio/gateway/releases) page. Once downloaded, the binary can be run from anywhere. Ideally, you should install it somewhere in your PATH for easy use. `/usr/local/bin` is the most probable location.

### MacOS

`gateway` is available via a Homebrew Tap, and as downloadable binary from the [releases](https://github.com/missingstudio/gateway/releases/latest) page:

```sh
brew install missingstudio/tap/gateway
```

To upgrade to the latest version:

```sh
brew upgrade gateway
```

### Linux

`gateway` is available as downloadable binaries from the [releases](https://github.com/missingstudio/gateway/releases/latest) page. Download the `.deb` or `.rpm` from the releases page and install with `sudo dpkg -i` and `sudo rpm -i` respectively.

### Windows

`gateway` is available via [scoop](https://scoop.sh/), and as a downloadable binary from the [releases](https://github.com/missingstudio/gateway/releases/latest) page:

```sh
scoop bucket add gateway https://github.com/missingstudio/scoop-bucket.git
```

To upgrade to the latest version:

```sh
scoop update gateway
```

### Docker

We provide ready to use Docker container images. To pull the latest image:

```sh
docker pull missingstudio/gateway:latest
```

To pull a specific version:

```sh
docker pull missingstudio/gateway:v0.0.1
```

### Docker compose

To start missing studio AI gateway, simply run the following command:

```sh
make up
```

> Your AI Gateway is now running on http://localhost:8080 💥

## Usage
Let's make a chat completion request to OpenAI through the AI Gateway using both REST and gRPC protocols

### Send a request using curl
```sh
curl \
--header "Content-Type: application/json" \
--header "x-ms-provider: openai" \
--header "Authorization: Bearer {{OPENAI_API_KEY}}" \
--data '{"model":"gpt-3.5-turbo","messages":[{"role":"user","content":"who are you?"}]}' \
http://localhost:8080/v1/chat/completions
```

### Send a request using grpcurl

```sh
grpcurl \
-d '{"model":"gpt-3.5-turbo","messages":[{"role":"user","content":"hi"}]}' \
-H 'x-ms-provider: openai' \
-H 'Authorization: Bearer {{OPENAI_API_KEY}}' \
-plaintext  localhost:8080  llm.v1.LLMService.ChatCompletions
```

## 🫶 Contributions
AI studio is an open-source project, and  contributions are welcome. If you want to contribute, you can create new features, fix bugs, or improve the infrastructure. 

It's still very early days for this so your mileage will vary here and lots of things will break. But almost any contribution will be beneficial at this point. Check the [current Issues](https://github.com/missingstudio/ai/issues) to see where you can jump in!

If you've got an improvement, just send in a pull request!

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'feat(module): add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

If you've got feature ideas, simply [open a new issues](https://github.com/missingstudio/ai/issues/new)!

Please refer to the [CONTRIBUTING.md](https://github.com/missingstudio/ai/blob/main/.github/CONTRIBUTING.md) file in the repository for more information on how to contribute.

<a href="https://github.com/missingstudio/ai/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=missingstudio/studio" />
</a>

## License
AI Studio is [Apache 2.0](https://github.com/missingstudio/ai/blob/main/LICENSE) licensed.
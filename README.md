![Gateway](/assets/gateway.svg)
![Missing studio](/assets/github.png)

## Introduction
🌈 A Robust Open Source [AI studio](https://www.missing.studio). A Universal API for inferencing 100+ LLMs(OpenAI, Azure, Cohere, Anthropic, HuggingFace, Replicate, Stable Diffusion).

## Supported Providers
|                                               |   Provider     |   Provider Name  |   Support   |   Supported Endpoints    |
|-----------------------------------------------|----------------|    :---:         |    :---:    |--------------------------|
|<img src="assets/openai.png" width=16>         | OpenAI         |     openai       |      ✅     |  `/chat/completions`     |
|<img src="assets/anyscale.png" width=16>       | Anyscale       |    anyscale      |      ✅     |  `/chat/completions`     |
|<img src="assets/deepinfra.jpeg" width=16>     | Deepinfra      |    deepinfra     |      ✅     |  `/chat/completions`     |
|<img src="assets/togetherai.svg" width=16>     | Together AI	   |    togetherai    |      ✅     |  `/chat/completions`     |

> Not supported (yet): images, audio, files, fine-tunes, moderations

## Run locally
Missing studio can be deployed in a variety of ways. It is deployable on bare metal, or in dockerized environments.

To start missing studio server, simply run the following command:
```sh
make compose-dev
```
> Your AI Gateway is now running on http://localhost:8080 💥

### Send a request
```sh
curl \
--header "Content-Type: application/json" \
--header "x-ms-provider: openai" \
--header "Authorization: Bearer {{OPENAI_API_KEY}}" \
--data '{"model":"gpt-3.5-turbo","messages":[{"role":"user","content":"who are you?"}]}' \
http://localhost:8080/v1/chat/completions
```


## 🚀 Features
- 🌐 *Universal API* - Call every LLM API like it's OpenAI
- *AI Gateway* for Security, Reliability and Observability
  - [ ] Load balancing across multiple provider and models   
  - [ ] Atomatic Retries with exponential backoff
  - [ ] Rate limiting
  - [ ] Caching
  - [ ] Failovers
  - [ ] Fallbacks
  - [ ] Monitoring
  - [ ] Alerting
  - [ ] Analytics
- AI Studio
- AI Agents
- AI Workflow builder
- OSS AI Models Inferancing
- Serving model api at Scale on Kubernetes 🦄️
- Building dev tools (CLI, SDK, API Client)
![Missing studio](/docs/assets/github.png)

## Introduction
🌈 A Robust Open Source [AI studio](https://www.missing.studio). A Universal API for inferencing 100+ LLMs(OpenAI, Azure, Cohere, Anthropic, HuggingFace, Replicate, Stable Diffusion).

## 🚀 Features
- 🌐 *Universal API* - Call every LLM API like it's OpenAI
  - [ ] Text completion (`/completions`) `inprogress`
    - [X] Non `stream` responses
    - [ ] `stream` responses
  - [ ] Not supported (yet): `images`, `audio`, `files`, `fine-tunes`, `moderations`
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
- Workflow builder to compose complex AI apps
- AI Agents
- LLM Studio
- Image studio
- Audio studio
- Video studio
- Storage to connect with providers like s3, gcs, vector DBs etc.
- Inference and serving to serve AI/ML models in production
- Finetune models
- Model Deployment at Scale on Kubernetes 🦄️
- Dev tools – all missing studio dev tools (CLI, SDK, API Client)

## Run Missing studio locally
Missing studio can be deployed in a variety of ways. It is deployable on bare metal, or in dockerized environments.

There are three officially supported ways of running missing studio locally:
1. Bare Metal
2. docker compose
3. Kubernetes via Tilt and Kind.


### Running with Compose
To start missing studio server, simply run the following command:
```sh
make compose-dev
```

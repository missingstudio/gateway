![Missing studio](/docs/assets/github.png)

## Introduction
[Missing studio](https://www.missing.studio) is a robust open source platform to build a developer-first way to create 
AI application using LLM, Image, Audio and Video Studios.


## Features
- 🌐 *Universal API* - Call every LLM API like it's OpenAI
  - [ ] Text completion (`/completions`) `inprogress`
    - Non `stream` responses 
    - `stream` responses
  - [ ] Not supported (yet): `images`, `audio`, `files`, `fine-tunes`, `moderations`
- *AI Gateway* for rate limit, atomatic retries, caching, failover etc
- Workflow builder to compose complex AI apps
- AI Agents
- LLM Studio to bring LLM applications to production
- Image studio
- Audio studio
- Video studio
- Storage to connect with providers like s3, gcs, vector DBs etc.
- Inference and serving to serve AI/ML models in production 
- Finetune models
- Model Deployment at Scale on Kubernetes 🦄️
- Dev tools – all missing studio dev tools (CLI, SDK, API Client)
# Contributing to AI studio

Thanks for being interested in contributing to AI studio. We'd love to accept your contributions!

We are extremely open to any and all contributions you might be interested in making.
To contribute to this project, please follow a ["fork and pull request"](https://docs.github.com/en/get-started/quickstart/contributing-to-projects) workflow. Please do not try to push directly to this repo unless you are a maintainer.

## Guidelines

### Code reviews

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

### Commit messages

Commit messages should follow
[Conventional Commits format](https://www.conventionalcommits.org/en/v1.0.0/#summary).

In particular, breaking changes should clearly be noted as “BREAKING CHANGE:” in
the commit message footer. Example:

```
fix(page): fix openai provider

This patch fixes openai provider so that it works with temperature.

Issues: #123

BREAKING CHANGE: openai now default temperature as 1 by default.
```

### Writing documentation

Documentation is generated using mintlify in docs branch. It is automatically
published to our documentation site on merge.

### Running the documentation site locally

To preview changes locally
1. Install Mintlify `npm i -g mintlify`.
2. run `mintlify dev` in docs folder of docs branch.
3. A local preview of your documentation will be available at http://localhost:3000


### Bug triage guidelines

The [issues](https://github.com/missingstudio/studio/issues) contain all current bugs, improvements, and feature requests. [Check incoming bug reports](https://github.com/missingstudio/studio/issues) that do not have a `confirmed` or `needs-feedback` label

1. Make sure the issue is labeled as either `bug` or `feature`.
2. If the issue does not have a clear repro or you cannot repro, ask for the repro and set the `needs-feedback` label.
3. Follow-up on the issues you previously asked for a feedback on (you should get a notification on GitHub when the user responds).
4. If the user does not provide feedback, the issue will be closed by the stale bot eventually.
5. If you are able to reproduce the issue, add the label `confirmed`.
6. If the bug is on the provider side, create a corresponding provider issue, label the GitHub issue with the `upstream` label, and post a link to provider issue in the comments.
7. If the issue is not related to either AI Studio or any provider, close the issue.
8. If the issue is about missing/incorrect documentation, label it as `documentation`.

### Getting Help

Contact a maintainer of AI studio with any questions or help you might need.
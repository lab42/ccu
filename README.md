![Conventional Commit Util](img/banner.svg)

<p align="center">
  <img src="https://img.shields.io/github/v/tag/lab42/ccu?label=latest%20tag" alt="Latest Tag" height="30" />
  <img src="https://github.com/lab42/ccu/actions/workflows/tag.yaml/badge.svg?event=push" alt="Build Status" height="30" />
  <img src="https://github.com/lab42/ccu/actions/workflows/main.yaml/badge.svg?branch=main&job=analyze" alt="CodeQL" height="30" />
  <img src="https://img.shields.io/github/go-mod/go-version/lab42/ccu" alt="Go Version" height="30" />
  <img src="https://img.shields.io/github/license/lab42/ccu" alt="License" height="30" />
  <a href="https://goreportcard.com/report/github.com/lab42/ccu">
    <img src="https://goreportcard.com/badge/github.com/lab42/ccu" alt="Go Report Card" height="30" />
  </a>
</p>


`ccu` is a tool for enforcing structured and compliant commit messages in Git repositories. It supports both command-line arguments and a YAML configuration file for flexibility, and can be easily integrated into CI/CD pipelines like GitHub Actions, GitLab CI, and Drone.

<h2 align="center">Installation</h2>

You can install `ccu` by downloading and running the `install.sh` script hosted in this repository. Choose either `curl` or `wget` based on your preference:

<h3 align="center">Install with `curl`</h3>

```sh
curl -fsSL https://raw.githubusercontent.com/lab42/ccu/refs/heads/main/install.sh | sh
```

<h3 align="center">Install with `wget`</h3>

```sh
wget -qO- https://raw.githubusercontent.com/lab42/ccu/refs/heads/main/install.sh | sh
```

The installation script automatically detects your OS and architecture, downloads the latest release of `ccu`, and places it in `/usr/local/bin`.

<h3 align="center">Install with Go</h3>

If you have Go installed, you can install the latest version of `ccu` directly:

```sh
go install github.com/lab42/ccu@latest
```

<h3 align="center">Run via Docker</h3>

`ccu` is also available as a Docker image:

```sh
docker pull ghcr.io/lab42/ccu:latest
```

Run `ccu` with Docker:

```sh
docker run --rm ghcr.io/lab42/ccu:latest [command] [options]
```

<h2 align="center">Configuration Options</h2>

`ccu` can be configured using command-line arguments, environment variables, or a YAML configuration file. The following options are available:

- `--type` / `CCU_TYPE`: Regular expression pattern for commit type (default: `build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test`)
- `--topic` / `CCU_TOPIC`: Regular expression pattern for commit topic (default: `(\([a-zA-Z0-9\-\.]+\))?(!)?`)
- `--message` / `CCU_MESSAGE`: Regular expression pattern for commit message (default: ` .*`)
- `--input` / `CCU_INPUT`: The commit message to validate
- `--config`: Path to config file (default: `$HOME/.ccu.yaml`)

<h3 align="center">Example YAML Configuration File</h3>

You can also use a YAML configuration file to specify options. Here's an example:

```yaml
# .ccu.yaml
type: "build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test"
topic: "(\([a-zA-Z0-9\-\.]+\))?(!)?"
message: " .*"
```

The config file will be automatically loaded from your home directory if named `.ccu.yaml`. Alternatively, you can specify a custom path:

```sh
ccu --config path/to/config.yaml
```

<h2 align="center">Usage</h2>

After installation, you can use `ccu` to validate commit messages in various ways.

<h3 align="center">Basic Command-Line Usage</h3>

```sh
ccu --input "feat(auth): add OAuth2 support"
```

<h3 align="center">Using Environment Variables</h3>

Set the configuration via environment variables:

```sh
export CCU_TYPE="build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test"
export CCU_TOPIC="(\([a-zA-Z0-9\-\.]+\))?(!)?"
export CCU_MESSAGE=" .*"
export CCU_INPUT="feat(auth): add OAuth2 support"
ccu
```

<h2 align="center">CI Integration</h2>

<h3 align="center">GitHub Actions</h3>

To integrate `ccu` with GitHub Actions, add a step in your workflow that validates the latest commit message. Here's an example workflow configuration:

```yaml
name: ccu

on:
  push:
    branches:
      - main
      - 'feature/*'
  pull_request:

jobs:
  validate_commit:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v2

      - name: Install ccu
        run: |
          curl -fsSL https://raw.githubusercontent.com/lab42/ccu/refs/heads/main/install.sh | sh

      - name: Run ccu
        env:
          CCU_INPUT: ${{ github.event.pull_request.title }}
        run: |
          ccu
```

<h3 align="center">GitLab CI</h3>

For GitLab CI integration, add a job to your `.gitlab-ci.yml` file:

```yaml
stages:
  - validate

CCU:
  stage: validate
  image: ghcr.io/lab42/ccu:latest
  variables:
    CCU_INPUT: "$CI_MERGE_REQUEST_TITLE"
  script:
    - ccu
  only:
    - merge_requests
```

<h2 align="center">Note on Windows Support</h2>

Please be aware that I do not use Windows as part of my workflow. As a result, I cannot provide support for Windows-related issues or configurations. However, I do generate Windows executables as a courtesy for those who need them.

Thank you for your understanding!

<h2 align="center">Contributing</h2>

I welcome contributions to this project! If you have ideas for new features or improvements, please submit a feature request or contribute directly to the project.

<h2 align="center">License</h2>

This project is licensed under the [MIT License](LICENSE).
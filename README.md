# Conventional Commit Util

`ccu` is a tool for enforcing structured and compliant commit messages in Git repositories. It supports both command-line arguments and a YAML configuration file for flexibility, and can be easily integrated into CI/CD pipelines like GitHub Actions, GitLab CI, and Drone.

## Installation

You can install `ccu` by downloading and running the `install.sh` script hosted in this repository. Choose either `curl` or `wget` based on your preference:

### Install with `curl`

```sh
curl -fsSL https://raw.githubusercontent.com/lab42/ccu/refs/heads/main/install.sh | sh
```

### Install with `wget`

```sh
wget -qO- https://raw.githubusercontent.com/lab42/ccu/refs/heads/main/install.sh | sh
```

The installation script automatically detects your OS and architecture, downloads the latest release of `ccu`, and places it in `/usr/local/bin`.

### Install with Go

If you have Go installed, you can install the latest version of `ccu` directly:

```sh
go install github.com/lab42/ccu@latest
```

### Run via Docker

`ccu` is also available as a Docker image:

```sh
docker pull ghcr.io/lab42/ccu:latest
```

Run `ccu` with Docker:

```sh
docker run --rm ghcr.io/lab42/ccu:latest [command] [options]
```

## Configuration Options

`ccu` can be configured using command-line arguments, environment variables, or a YAML configuration file. The following options are available:

- `--type` / `CCU_TYPE`: Regular expression pattern for commit type (default: `build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test`)
- `--topic` / `CCU_TOPIC`: Regular expression pattern for commit topic (default: `(\([a-zA-Z0-9\-\.]+\))?(!)?`)
- `--message` / `CCU_MESSAGE`: Regular expression pattern for commit message (default: ` .*`)
- `--input` / `CCU_INPUT`: The commit message to validate
- `--config`: Path to config file (default: `$HOME/.ccu.yaml`)

### Example YAML Configuration File

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

## Usage

After installation, you can use `ccu` to validate commit messages in various ways.

### Basic Command-Line Usage

```sh
ccu --input "feat(auth): add OAuth2 support"
```

### Using Environment Variables

Set the configuration via environment variables:

```sh
export CCU_TYPE="build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test"
export CCU_TOPIC="(\([a-zA-Z0-9\-\.]+\))?(!)?"
export CCU_MESSAGE=" .*"
export CCU_INPUT="feat(auth): add OAuth2 support"
ccu
```

## CI Integration

### GitHub Actions

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

### GitLab CI

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

## Note on Windows Support

Please be aware that I do not use Windows as part of my workflow. As a result, I cannot provide support for Windows-related issues or configurations. However, I do generate Windows executables as a courtesy for those who need them.

Thank you for your understanding!

## Contributing

I welcome contributions to this project! If you have ideas for new features or improvements, please submit a feature request or contribute directly to the project.

## License

This project is licensed under the [MIT License](LICENSE).
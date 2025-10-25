# Steam Hour Booster

## Overview

Steam Hour Booster is a headless Go application that simulates running Steam titles on your account so you can accumulate playtime without keeping the Steam client open. The tool is purpose built for automation: it loads configuration from flags, environment variables, or a YAML file, and persists session information so repeated logins are frictionless.

## Quick Start

1. Install Go 1.21 or newer and clone this repository.
2. Build the booster:
   ```bash
   go build -o booster ./cmd/booster
   ```
3. Export your credentials and launch the games you want to idle:
   ```bash
   export STEAM_USERNAME="your_steam_username"; export STEAM_PASSWORD="your_steam_password"; ./booster --appids "730,570" --max-concurrent 2
   ```
   On the first run you may be prompted in the terminal for a Steam Guard code. The booster saves the resulting sentry file and login key in the data directory (default `~/.local/steam-booster/`), letting subsequent runs skip the password and guard code entirely.
4. Once a session has been remembered you can run with just your username and desired flags:
   ```bash
   ./booster --appids "730,570" --max-concurrent 2
   ```

## Configuration

### Precedence

Configuration values are resolved in the following order (highest priority first):

1. Command-line flags
2. Environment variables
3. `config.yaml` (or the file provided via `--config`)

### Configuration file

Copy [`config.example.yaml`](./config.example.yaml) to `config.yaml` and adjust it to your needs:

```yaml
username: your_user
# password: your_pass  # optional; prefer ENV
appids: [730, 570]
max_concurrent: 2
heartbeat_sec: 60
backoff_min_sec: 1
backoff_max_sec: 60
log_level: info
socks5: "socks5://127.0.0.1:1080"
```

The configuration file pairs naturally with environment variables and flags—any value omitted here will fall back to the other sources.

### Environment variables

The following environment variables can be used to configure the booster:

| Variable | Description |
| --- | --- |
| `STEAM_USERNAME` | Steam account username. |
| `STEAM_PASSWORD` | Steam account password. Only required on the first run if a login key is not remembered. |
| `STEAM_APPIDS` | Comma-separated Steam app IDs to idle. |
| `STEAM_MAX_CONCURRENT` | Maximum number of games to idle simultaneously. |
| `STEAM_DATA_DIR` | Directory used to store session data and sentry files (default `~/.local/steam-booster/`). |
| `STEAM_SOCKS5` | SOCKS5 proxy URL to route Steam traffic through. |
| `STEAM_HEARTBEAT_SEC` | Seconds between heartbeats sent to Steam while idling. |
| `STEAM_BACKOFF_MIN_SEC` | Minimum seconds before retrying after a failure. |
| `STEAM_BACKOFF_MAX_SEC` | Maximum seconds before retrying after a failure. |
| `STEAM_LOG_LEVEL` | Log verbosity (`debug`, `info`, `warn`, `error`). |
| `HTTP_PROXY` / `HTTPS_PROXY` | Standard proxy environment variables respected by the Go HTTP client. |

Unset variables simply defer to the config file or internal defaults.

### Session caching & Steam Guard

During the first login you will be prompted for your Steam password (if not provided via flags or environment variables) and a Steam Guard one-time code. The booster then writes a sentry file and login key into the data directory. On subsequent runs you can omit the password and guard code; the stored session will be reused unless you pass `--no-remember` or manually delete the cache (see `--forget-session` under Flags).

## Flags

Every configuration value is also accessible via a flag. Use `./booster --help` to view the current list. The flags below correspond to the environment variables listed earlier.

| Flag | Description | Environment variable |
| --- | --- | --- |
| `--username string` | Steam account username. | `STEAM_USERNAME` |
| `--password string` | Steam account password. | `STEAM_PASSWORD` |
| `--appids string` | Comma-separated Steam app IDs to idle. | `STEAM_APPIDS` |
| `--max-concurrent int` | Maximum number of games to idle simultaneously. | `STEAM_MAX_CONCURRENT` |
| `--data-dir string` | Directory used to store session/sentry data. | `STEAM_DATA_DIR` |
| `--socks5 string` | SOCKS5 proxy URL. | `STEAM_SOCKS5` |
| `--heartbeat int` | Heartbeat interval in seconds. | `STEAM_HEARTBEAT_SEC` |
| `--backoff-min int` | Minimum retry backoff in seconds. | `STEAM_BACKOFF_MIN_SEC` |
| `--backoff-max int` | Maximum retry backoff in seconds. | `STEAM_BACKOFF_MAX_SEC` |
| `--log-level string` | Log verbosity (`debug`, `info`, `warn`, `error`). | `STEAM_LOG_LEVEL` |
| `--config string` | Path to a YAML configuration file. Defaults to `./config.yaml`. | — |
| `--no-remember` | Disable storing the login key and sentry file for future runs. | — |
| `--forget-session` | Delete cached session data before logging in. | — |

Example help output:

```text
$ ./booster --help
Usage of ./booster:
  -appids string
        Comma-separated Steam app IDs to idle (overrides STEAM_APPIDS)
  -backoff-max int
        Maximum retry backoff in seconds (overrides STEAM_BACKOFF_MAX_SEC)
  -backoff-min int
        Minimum retry backoff in seconds (overrides STEAM_BACKOFF_MIN_SEC)
  -config string
        Path to configuration file (default "./config.yaml")
  -data-dir string
        Directory used to store session/sentry data (overrides STEAM_DATA_DIR)
  -forget-session
        Forget any stored session before logging in
  -heartbeat int
        Heartbeat interval in seconds (overrides STEAM_HEARTBEAT_SEC)
  -log-level string
        Log level to use (overrides STEAM_LOG_LEVEL)
  -max-concurrent int
        Maximum number of games to idle simultaneously (overrides STEAM_MAX_CONCURRENT)
  -no-remember
        Do not persist session data between runs
  -password string
        Steam account password (overrides STEAM_PASSWORD)
  -socks5 string
        SOCKS5 proxy URL (overrides STEAM_SOCKS5)
  -username string
        Steam account username (overrides STEAM_USERNAME)
```

## Disclaimer

Boosting or idling playtime may violate Steam’s Terms of Service. Use this tool responsibly and at your own risk.

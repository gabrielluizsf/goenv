# goenv

Run any app with environment variables, identically (no more `VAR=x cmd` quirks).

## Usage

```
goenv [flags] -- <command> [args...]

```

| Flag | Description |
| --- | --- |
| `-e KEY=VALUE` | Set an environment variable (repeatable) |
| `-env-file PATH` | Load variables from a `.env` file |
| `-C DIR` | Working directory for the app |
| `-log PATH` | Log file for background mode (default `goenv.log`) |
| `-fg` | Run in foreground (do not detach) |

Precedence: current environment < `.env` file < `-e` flags.

### Examples

```
goenv -e PORT=8080 -e DB_URL=postgres://localhost/app -- go run ./cmd/api
goenv -env-file .env -- ./api
goenv -env-file .env -e PORT=9000 -C ./backend -- go run .

```

### `.env` format

```
# comments are ignored
PORT=8080
export DB_URL="postgres://localhost/app"
API_KEY='secret'

```

## Output

After starting, goenv prints the PID and a ready-to-copy kill command:

```
✔ App started (runtime.GOOS)
  PID: 12345
  Log: goenv.log

To stop (copy and paste):
  kill -TERM -- -12345          # UNIX (Linux / macOS)
  taskkill /PID 12345 /T /F     # Windows

```
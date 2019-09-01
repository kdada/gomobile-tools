# gomobile-tools


## cmdrunner

1. Build `./cmd/server` on target machine (ex. MacOS, Linux).
2. Build `./cmd/client` on development machine.
3. Run `./server -a 'ip:port' -h '/absolute/path/to/project'`
4. Run `COMMAND_RUNNER_ADDR=ip:port client make`. this command will run `make` at the target project.

## debuglogger

1. Build&Run `./cmd/server -a 'ip:port'`
2. Import `github.com/kdada/gomobile-tools/debuglogger/pkg/remote`
3. Call `remote.SetDebugLogger("ip:port")` in `func init`
4. All logs to `log` will redirect to the server.


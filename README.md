## PoC: Stream log file content (Python >> Go)

### Usage
- Generate logs via `cd python && python generate-logs.py`
- Consume logs via `cd go && go run consume-logs.go`

The python code snippet shows
- how to instantiate a logger that streams data into a file (`.shared-fs/user-logs.log`)

The go code snippets shows
- how to continuously read data from a file (`.shared-fs/user-logs.log`)
  - there may be more efficient ways!
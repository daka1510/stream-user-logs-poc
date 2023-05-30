## POC Stream Logs

### Usage
- Generate logs via `cd python && python generate-logs.py`
- Consume logs via `cd go && go run consume-logs.go`

The python code snippet shows
- how to instantiate a logger that streams data into a file

The go code snippets shows
- how to continuously read data from a file (there may be more efficient ways!)
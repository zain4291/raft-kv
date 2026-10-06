# Learning Notes

A running log of what I learn while building raft-kv.

## Go basics

- Files in the same folder with `package main` are compiled together. `go run .` runs the whole package (the `.` matters).
- A map must be initialized (`map[string]string{}` or `make(...)`) before writing, or writing to it panics.
- `value, ok := m[key]` returns the value and whether the key existed. Check `ok`, not `value`, because a missing key returns `""`, which could also be a real stored value.
- `delete(m, key)` is the built-in way to remove a key. It does nothing if the key is missing.
- Go refuses to compile if a variable is declared and never used.
- `:=` creates new variables; `=` reassigns existing ones.
- Methods use a receiver: `func (s *Store) Get(...)`, so `s` is the store the method is called on.
- Comments start with the function name (`// Put saves...`) by Go convention.

## Concurrency

- Writing to a map from many goroutines at once crashes with `fatal error: concurrent map writes`.
- Fix: a `sync.Mutex` in the struct. `Lock()` before touching the map, and `defer Unlock()` so it always releases.
- Only one goroutine can be inside the locked section at a time (the critical section). Reads need the lock too, since a read racing a write is also unsafe.
- `go run -race .` runs the race detector. It finds unsafe access even when the program doesn't crash. Silent output means clean; `DATA RACE` means a bug.
- Verified both ways: with the lock removed the detector reported races at `Put`; with the lock it was clean.

## Tooling gotchas

- `main.go: expected 'package', found 'EOF'` means the file was empty on disk. I hadn't saved it (Ctrl+S).
- `go run -race` without the `.` gives `no go files listed`.
- Create the GitHub repo empty (no README or license) to avoid push conflicts with local files.

## HTTP API
- Go's HTTP server runs each request in its own goroutine automatically, so many requests can hit the store at once. That's why the mutex matters.
- Handled PUT, GET, and DELETE requests
- 204 No Content indicates that the server successfully processed the request, but intentionally has no data or content to send back in the response body
- 404 Not Found indicates that the requested resource on the server could not be found, in this case, the key couldn't be found or isn't in the store 

## Project progress

- [x] Single-node in-memory store (`Get`, `Put`, `Delete`)
- [x] Mutex for concurrent safety, confirmed with `-race`
- [ ] Network interface (gRPC)
- [ ] Raft: leader election
- [ ] Raft: log replication
- [ ] Persistence
- [ ] Benchmarks

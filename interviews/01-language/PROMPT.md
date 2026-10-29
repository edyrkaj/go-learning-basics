# Language simulation (25 minutes)

Rules:

- Talk through your approach before editing.
- Use the standard library only.
- Leave `solution/` closed until you stop.
- From this directory, run `go run .`
- Done when the process exits 0. Every line of output is `PASS`.

## Task

`starter.go` compiles, and three functions are wrong. Keep the signatures.

1. `Prefix` returns a slice that shares the caller's backing array. Appending to the result must leave the original slice unchanged.
2. `Label` reports every key as present. A stored `0` and a missing key must be different results. Use the comma-ok map lookup.
3. `NewStore("")` returns a nil `*Mem` inside a non-nil interface, so `IsEmpty` is false. An empty name must return a nil interface.

## Follow-up

Explain why a nil `*Mem` stored in a `Store` is not equal to `nil`, and why `Prefix` must copy.

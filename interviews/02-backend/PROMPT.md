# Backend simulation (35 minutes)

Rules:

- Talk through your approach before editing.
- Use the standard library only.
- Leave `solution/` closed until you stop.
- From this directory, run `go run .`
- Done when the process exits 0. Every line of output is `PASS`.
- Do not listen on a port. The checks use `httptest`.

## Task

Implement the empty methods in `starter.go`.

- `MemoryStore` keeps jobs in a map guarded by `sync.Mutex`.
- Ids are `job-1`, `job-2`, and so on.
- `Put` assigns the id, sets `Status` to `running`, and returns the stored job.
- `Get` returns the job. A missing id wraps `ErrNotFound`.
- `Put` and `Get` return `fmt.Errorf("store: %w", ctx.Err())` when the context is already cancelled.
- `POST /jobs` with `{"name":"compile"}` responds `201` and the job JSON. An empty name or bad JSON responds `400`.
- `GET /jobs/{id}` responds `200` and the job, or `404` with body `not found` when the id is missing. Any other store error responds `500`.
- The checks require the first id to be `job-1`, HTTP `400` with body `bad request` for an empty name or bad JSON, and `errors.Is(err, context.Canceled)` from `Put`, `Get`, and `Cancel` when the context is already cancelled.

## Follow-up

`POST /jobs/{id}/cancel` sets `Status` to `cancelled` and responds `200` with the job. A missing id responds `404` with body `not found`. `Cancel` uses the same cancelled-context rule as `Put` and `Get`.

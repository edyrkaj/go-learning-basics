# Coding simulation (30 minutes)

Rules:

- Talk through your approach and the complexity before editing.
- Use the standard library only.
- Leave `solution/` closed until you stop.
- From this directory, run `go run .`
- Done when the process exits 0. Every line of output is `PASS`.

## Task

Implement `GroupAnagrams`. Words with the same letters and the same counts form one group. Build the map key by sorting each word's runes.

Return value:

- Sort the words inside each group lexicographically.
- Sort the groups by the first word of each group.
- `["eat","tea","tan","ate","nat","bat"]` returns `[["ate","eat","tea"],["bat"],["nat","tan"]]`.
- A nil or empty input returns an empty non-nil slice.

## Follow-up

State the cost on the function comment: time `O(n · k(log k + log n))`, space `O(n · k)`, where `k` is the length of the longest word. The extra `log n` is the cost of sorting the groups. The sorted output above is what the check compares.

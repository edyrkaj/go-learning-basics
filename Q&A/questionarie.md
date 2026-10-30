Here is a complete, compiled Q&A review guide covering all the core Go concepts from our session today:

---

**Q1: What is the difference between an unbuffered and a buffered channel in Go, and what happens when you try to send data to a full buffered channel?**

* **Answer:** An unbuffered channel requires a synchronous handoff where both sender and receiver must meet simultaneously. A buffered channel has internal capacity, allowing the sender to write without blocking until the buffer fills up. Sending to a full buffered channel **blocks** the calling goroutine until space opens; it does not panic.

**Q2: In Go, what is the difference between passing a parameter by value versus passing a pointer? What happens when you pass a slice to a function and modify its elements inside?**

* **Answer:** Everything in Go is passed by value. Passing a value creates a local copy, while passing a pointer creates a copy of the memory address to mutate the original caller data. When passing a slice, its 3-field header (pointer, length, capacity) is copied, but it still points to the same underlying array. Modifying elements within capacity updates the caller's array directly; however, appending beyond capacity allocates a new underlying array that will not be visible to the caller unless returned.

**Q3: What is the difference between a value receiver and a pointer receiver on a struct method? How does it affect interface satisfaction?**

* **Answer:** Value receivers operate on a copy of the struct, making them safe for read-only calls on small structs, and allow both values (`T`) and pointers (`*T`) to satisfy an interface. Pointer receivers operate on the original memory address to allow state mutation, but restrict interface satisfaction strictly to pointer types (`*T`).

**Q4: How does error handling work in Go? Specifically, what is error wrapping using `fmt.Errorf`, and how do `errors.Is` and `errors.As` differ?**

* **Answer:** Errors in Go are handled explicitly by returning them as ordinary values. `fmt.Errorf` with the `%w` verb wraps an error to preserve its stack context. `errors.Is` checks if a specific target error exists anywhere within the wrapped chain (equality check), whereas `errors.As` searches the chain for a specific error type and populates a target pointer (type assertion/cast).

**Q5: How does `defer` work in Go? In what order do deferred statements execute, when are their arguments evaluated, and how do you recover from a panic inside a deferred function?**

* **Answer:** `defer` schedules a function call to execute immediately after the surrounding function returns. Multiple deferred statements execute in Last-In, First-Out (LIFO) order. Arguments passed to `defer` are evaluated immediately at the moment the `defer` line is reached. Panics can be caught and stopped by calling `recover()` directly inside a deferred function to resume normal execution flow.

**Q6: What is the dynamic stack size of a Goroutine compared to an OS thread, and how does the Go scheduler work?**

* **Answer:** Goroutines start with a small, dynamic stack of ~2KB that grows and shrinks as needed, whereas standard OS threads carry fixed stacks of ~1–2MB. The Go runtime uses an M:N scheduler (mapping $M$ goroutines onto $N$ OS threads across $P$ logical processors) to handle light-weight, highly concurrent execution.

**Another 10/10 answer!** You captured the core engineering trade-offs perfectly:

* **Protobuf vs. JSON:** Protocol Buffers use a binary encoding format with field tags (integers) rather than field names (strings), producing drastically smaller payloads and requiring far less CPU overhead to serialize/deserialize than text-based JSON parsing.
* **Strict Schema Contracts:** `.proto` files define a clear interface contract. The `protoc` compiler auto-generates type-safe client stubs and server interface handlers in Go, eliminating runtime type mismatch bugs.
* **Transport & Streaming:** Built directly on top of **HTTP/2**, gRPC supports unary calls as well as client, server, and bi-directional streaming natively using Go channels and context controls.

---

### Master Q&A Cheat Sheet (Advanced Topics)

Here are the questions and ideal answers from this second round to add to your personal notes:

**Q7: How does Go's Escape Analysis and Garbage Collection work?**

* **Answer:** Go's compiler uses **escape analysis** at compile time to decide whether memory is allocated on the stack or the heap. Variables bounded by a function's stack frame are placed on the stack for zero-cost cleanup when the function returns. If a pointer escapes the function scope (e.g., returning a pointer), it is allocated on the heap. The heap is managed by a concurrent **tricolor mark-and-sweep Garbage Collector** (White = candidate for collection, Grey = work queue, Black = reachable/retained), designed for microsecond-level Stop-The-World (STW) pause times.

**Q8: What are the key advantages of HTTP/2 over HTTP/1.1?**

* **Answer:** HTTP/2 introduces **multiplexing** (sending/receiving multiple requests concurrently over a single TCP connection, eliminating HTTP-level Head-of-Line blocking), **binary framing** (replacing plain text parsing with compact binary frame encoding), and **HPACK header compression** (drastically reducing HTTP metadata overhead using static/dynamic header tables).

**Q9: Why is gRPC faster than traditional REST/JSON APIs, and how does it integrate with Go?**

* **Answer:** gRPC leverages **Protocol Buffers (Protobuf)** for compact, binary serialization, which is significantly smaller and faster to encode/decode than human-readable text JSON. It enforces strict, strongly-typed contracts defined in `.proto` files, compiled via `protoc` into native Go code and stubs. Built on top of HTTP/2, it supports high-throughput unary requests and bi-directional streaming seamlessly.

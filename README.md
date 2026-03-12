<div align="center">
  <h1>🚀 Go Learning Basics</h1>
  <p>A personal repository to master the fundamentals and advanced concepts of the Go Programming Language.</p>
</div>

<br />

## 📖 About This Repository

This repository, **go-learning-basics**, serves as a personal playground and reference guide for learning Go (Golang). It contains clear, concise, and heavily-commented code examples demonstrating various language features, ranging from basic string manipulations to complex concurrency patterns.

> Feel free to explore the code, run the examples, and use it as a learning reference!

---

## 📂 Repository Structure

The project is thoughtfully organized into modular directories based on topic:

### 🧵 [Concurrency & Channels](./channels)
An exploration of Go's powerful concurrency model using **goroutines** and **channels**.
*   **Basic Channels:** Asynchronous data sending and receiving.
*   **Buffered Channels:** Understanding channel capacity and non-blocking operations.
*   **Worker Pools:** Executing complex tasks in parallel utilizing bounded concurrency.
*   **Select Statements:** Handling operations over multiple channels simultaneously.
*   **Synchronization:** Safely orchestrating goroutines and avoiding deadlocks using `sync.WaitGroup`.

### 🔤 [Strings & Runes](./strings)
Understanding how Go safely handles text, bytes, and Unicode characters.
*   **String Manipulation:** Exploring the differences between `string`, byte slices (`[]byte`), and rune slices (`[]rune`).
*   **In-place Operations:** Utilizing Go 1.21+ standard packages (like `slices.Reverse`) to modify internal elements securely.

### 🛠️ [CLI Commands](./commands)
Examples demonstrating how to build robust Command-Line Applications in Go.
*   Follows standard Go project layouts (`cmd/` and `pkg/` directories).
*   Contains structured binaries easily compilable with `go build`.

---

## 🏃 Getting Started

To run any of the examples in this repository, simply navigate to the topic's folder and execute the Go files directly using the `go run` tool.

As an example, to see the channels demonstration in action:

```bash
# Clone the repository
git clone git@github.com:edyrkaj/go-learning-basics.git
cd go-learning-basics

# Navigate to a specific topic
cd channels

# Run the application
go run main.go
```

---

## 🧠 Prerequisites

Ensure you have the Go toolchain installed on your machine to execute this code.

*   [Go 1.21+](https://go.dev/doc/install) (The examples use recent standard library packages like `slices`).

---
<div align="center">
  <i>Written in Go 🐹</i>
</div>

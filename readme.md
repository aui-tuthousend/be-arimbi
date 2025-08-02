# Go Project with Google Wire Dependency Injection

This project uses [Google Wire](https://github.com/google/wire) for compile-time dependency injection in Go.

---

---

## ⚙️ Prerequisites

- Go 1.20+
- [Wire CLI](https://github.com/google/wire) installed

Install Wire once using:

```bash
go install github.com/google/wire/cmd/wire@latest
```

Install dependencies:
```bash
go mod tidy
```

Generate wire files:
```bash
cd container
wire
```

Run the application:
```bash
cd ..
go run main.go
```

or use air for hot reloading




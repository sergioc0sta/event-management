# Event Management (Go)

A simple event management system written in Go, designed to demonstrate event dispatching, handler registration, and event clearing using interfaces and test-driven development.

## Features
- Register event handlers for specific event names
- Prevent duplicate handler registration for the same event
- Clear all registered event handlers
- Unit tests using `testify/suite`

## Project Structure
```
go.mod
go.sum
pkg/
  events/
    events_dispatcher.go      # Event dispatcher implementation
    events_dispatcher_test.go # Unit tests for dispatcher
    interfaces.go             # Event and handler interfaces
```

## Getting Started
1. **Clone the repository**
   ```sh
   git clone <your-repo-url>
   cd event-management
   ```
2. **Install dependencies**
   ```sh
   go mod tidy
   ```
3. **Run tests**
   ```sh
   go test ./pkg/events
   ```

## Usage Example
```go
// Create dispatcher
ed := events.NewEventDispatcher()

// Register handler
err := ed.Register("MyEvent", &MyHandler{})
if err != nil {
    log.Fatal(err)
}

// Clear all handlers
err = ed.Clear()
if err != nil {
    log.Fatal(err)
}
```

## Technologies
- Go (Golang)
- Testify (for unit testing)

## License
MIT

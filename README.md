# Fitness Tracker

Fitness Tracker is a web application built with Go and the Fiber web framework. It provides a simple API to track fitness activities.

## Project Structure

- **api/**: Contains the HTTP handlers.
- **cmd/**: Contains the entry point of the application.
- **internal/**: Contains the application-specific business logic.
- **test/**: Contains test files.
- **go.mod**: The Go module file.
- **go.sum**: The Go module dependencies file.
- **Makefile**: Contains commands to run the application and tests.

## Getting Started

### Prerequisites

- Go 1.22.5 or later

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/fitness-tracker.git
    cd fitness-tracker
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

### Running the Application

To run the application, use the following command:

```sh
make run
```

# HealthChecker

HealthChecker is a command-line tool written in Go that checks whether a website is running or down by attempting to establish a TCP connection to a specified domain and port.

## Features

- **Domain and Port Checking**: Checks the status of a specified domain and port combination.
- **Flexible Input**: Allows users to specify the domain and optionally the port via command-line flags.
- **Timeout Handling**: Uses timeouts to control how long it waits for a connection to be established.
- **Error Reporting**: Provides detailed error messages when a connection cannot be made.

### Prerequisites

1. Make sure you have [Go](https://golang.org/dl/) installed on your machine.

### Installation

2. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/HealthChecker.git
   cd HealthChecker

### Run program

    To run the program: 
    ```bash
    go run . --domain example.com
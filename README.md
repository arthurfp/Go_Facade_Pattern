# Facade Pattern in Go

## Overview
This repository demonstrates the application of the Facade design pattern in Go. The project highlights how to provide a simplified interface to a complex subsystem, showcasing flexibility and best practices in design patterns and unit testing.

## Pattern Description
The Facade pattern provides a simplified interface to a complex subsystem, making it easier to use and understand. In this project, we have implemented subsystems and a Facade to demonstrate this pattern effectively.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), demonstrating the creation and usage of the Facade.
- **pkg/facade/**: Houses the Facade implementation and its subsystems.
  - **component.go**: Defines the `Component` interface.
  - **subsystem_a.go**: Implements the `SubsystemA`.
  - **subsystem_b.go**: Implements the `SubsystemB`.
  - **subsystem_c.go**: Implements the `SubsystemC`.
  - **facade.go**: Implements the `Facade`.
  - **subsystem_a_test.go**: Unit tests for `SubsystemA`.
  - **subsystem_b_test.go**: Unit tests for `SubsystemB`.
  - **subsystem_c_test.go**: Unit tests for `SubsystemC`.
  - **facade_test.go**: Unit tests for the `Facade`.

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).


### Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/arthurfp/Go_Facade_Pattern.git
cd Go_Facade_Pattern
```

### Running the Application
To run the application, execute:
```bash
go run cmd/main.go
```

### Running the Tests
To execute the tests and verify the functionality:
```bash
go test ./...
```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp
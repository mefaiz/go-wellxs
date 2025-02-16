# Project go-wellxs

A REST API backend for the WellXs mobile application, converted from Laravel to Golang using the Fiber framework. WellXs is a comprehensive wellness application designed for foreign workers in Malaysia, providing easy access to healthcare services and medical screenings.

## About WellXs

WellXs is FOMEMA's official mobile application that helps foreign workers in Malaysia:
- Schedule and manage medical examinations
- Access health screening results
- Find nearby panel clinics
- Get health-related notifications and reminders
- Store and access medical records securely
- Receive important healthcare updates

## Tech Stack

- **Framework**: [Fiber](https://gofiber.io/) - Express-inspired web framework for Go
- **ORM**: [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- **Database**: MySQL
- **Authentication**: JWT

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.21 or higher
- MySQL
- Docker (optional, for containerized database)

### Installation

1. Clone the repository
```bash
git clone https://github.com/yourusername/go-wellxs.git
```

2. Install dependencies
```bash
go mod download
```

3. Set up your environment variables (copy from example)
```bash
cp .env.example .env
```

## Development Commands

Build and test the application
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

Run integration tests
```bash
make itest
```

Live reload the application
```bash
make watch
```

Run the test suite
```bash
make test
```

Clean up binary from the last build
```bash
make clean
```

## Key Features

- JWT Authentication
- User Management
- RESTful API endpoints
- MySQL Database Integration
- Input Validation
- Error Handling

## Contributing

[Add contribution guidelines here]

## License

This project is licensed under the MIT License - see the LICENSE file for details

## Acknowledgments

- Original Laravel project team
- Fiber framework community
- GORM community

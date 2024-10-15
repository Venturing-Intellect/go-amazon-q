# Feedback Application

This is a full-stack feedback application built with Go for the backend and React for the frontend. It allows users to submit feedback which is then stored in a PostgreSQL database.

## Project Structure

```
├── Dockerfile
├── Makefile
├── README.md
├── controller/
│ ├── controller_test.go
│ └── main.go
├── docker-compose.yml
├── frontend/
│ ├── package-lock.json
│ ├── package.json
│ ├── public/
│ │ └── index.html
│ └── src/
│ ├── App.js
│ ├── components/
│ │ └── FeedbackForm.js
│ └── index.js
├── go.mod
├── go.sum
├── init_repo.sh
├── main.go
├── main_test.go
├── middleware/
├── migrations/
│ └── 20241015150636_init_table.sql
├── repository/
│ ├── main.go
│ └── repository_test.go
└── service/
├── main.go
└── service_test.go
```


## Backend (Go)

The backend is structured using a ports and adapters (hexagonal) architecture:

- `controller/`: HTTP handlers for the Feedback API
- `service/`: Business logic layer
- `repository/`: Data access layer
- `migrations/`: Database migration files

### Prerequisites

- Go 1.23 or later
- PostgreSQL

### Setup

1. Clone the repository
2. Navigate to the project root
3. Run `go mod download` to install dependencies

### Database Setup

1. Create a PostgreSQL database
2. Update the database connection details in your environment or configuration file

### Running Migrations

```
export POSTGRES_HOST=localhost
&& export POSTGRES_USER=postgres
&& export POSTGRES_PASSWORD=mysecretpassword
&& export POSTGRES_DB=feedback_db
&& make migrate-up
```


### Running the Application

```
go run main.go
```


### Running Tests

```
go test ./...
```


## Frontend (React)

The frontend is a React application located in the `frontend/` directory.

### Prerequisites

- Node.js 14 or later
- npm

### Setup

1. Navigate to the `frontend/` directory
2. Run `npm install` to install dependencies

### Running the Frontend

```
npm start
```


### Building for Production

```
npm run build

```


## Docker

The application can be run using Docker:

```
docker-compose up --build

```


This will start backend, as well as a PostgreSQL database.

## API Endpoints

- `POST /feedback`: Submit feedback
  - Request body: `{ "name": "string", "email": "string", "feedback": "string" }`

## Testing

### Unit Tests

Unit tests are available for each package. Run them using: `go test ./...`

### Integration Tests

Integration tests are in `main_test.go`. To run:

1. Set up the test database and environment variables:

```
export DB_USER=your_test_db_user
export DB_PASSWORD=your_test_db_password
export DB_HOST=your_test_db_host
export DB_PORT=your_test_db_port
export DB_NAME=your_test_db_name
```

2. Run the tests:

```
go test -v ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

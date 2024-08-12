# Notes Backend 🚀

Welcome to the backend of Notes! This guide will help you understand the file structure and get your environment up and running. Let's dive in!

## 🗂 Project Structure

Here's an overview of the backend file structure:

```
backend/
├── cmd/
│ └── server/
│ └── main.go # Entry point for the application
├── internal/
│ ├── auth/
│ │ └── auth.go # Authentication logic
│ ├── database/
│ │ ├── init.sql # Database initialization script
│ │ └── database.go # Database connection setup
│ ├── handlers/
│ │ ├── journal.go # Handlers for journal-related routes
│ │ └── user.go # Handlers for user-related routes
│ ├── middleware/
│ │ └── auth_middleware.go # Authentication middleware
│ └── models/
│ ├── journal.go # Journal model
│ └── user.go # User model
├── config/
│ └── config.go # Configuration settings
├── go.mod # Go module file
└── go.sum # Go dependencies checksum
```

## 🛠 Getting Started

Follow these steps to set up and run the backend locally.

### Prerequisites

Make sure you have the following tools installed:

-   [Docker](https://www.docker.com/get-started)
-   [Go](https://golang.org/doc/install)

### 🚀 Setup Instructions

1. **Clone the repository**

    ```sh
    git clone https://github.com/bartosz-skejcik/notes.git
    cd notes/backend
    ```

2. **Create a .env file**

    Create a .env file in the backend directory with the following content:

    ```sh
    POSTGRES_USER=myuser
    POSTGRES_PASSWORD=mypassword
    POSTGRES_DB=mydb
    ```

3. **Run Docker Compose**

    Start the Docker containers:

    ```sh
    docker-compose up --build
    ```

    This will set up a PostgreSQL database and run the initialization script to create the necessary tables and indexes.

4. **Run the Go server**

    In a new terminal, navigate to the backend directory and run:

    ```sh
    go run cmd/server/main.go
    ```

    This will start the Fiber server on `http://localhost:3000`.

### 🎨 Code Overview

-   main.go: The entry point of the application. Initializes the Fiber app and starts the server.

-   auth/: Contains authentication logic and middleware.

-   database/: Sets up the database connection and contains the initialization SQL script.

-   handlers/: Contains route handlers for different functionalities like user and journal management.

-   middleware/: Contains middleware for handling authentication.

-   models/: Defines the data models for the application.

-   config/: Manages configuration settings.

### 📋 Example Requests

Here are some example requests to test the API:

-   Get Users

    ```sh
    curl http://localhost:3000/users
    ```

-   Create a Journal Entry

    ```sh
    curl -X POST http://localhost:3000/journal \
    -H "Content-Type: application/json" \
    -d '{
        "user_id": "user-id-here",
        "content": "My first journal entry",
        "date": "2024-08-07",
        "type": "personal"
        }'
    ```

### 🔧 Development Tips

-   **Database Access**: To access the database directly, use the following command:

    ```sh
    docker exec -it my_postgres psql -U myuser -d mydb
    ```

-   **Rebuild Containers**: If you make changes to the Docker configuration or initialization script, rebuild the containers:

    ```sh
    docker-compose down
    docker-compose up --build
    ```

-   **Environment Variables**: Manage sensitive information like database credentials using environment variables.

### 💬 Need Help?

If you encounter any issues or have questions, feel free to open an issue or reach out to the project maintainers.

Happy coding! 🎉

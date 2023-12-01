

# Messaging App

Messaging App is a simple web application that allows a team of agents to respond to customer inquiries in a streamlined fashion. It uses Go as the programming language, Gin as the web framework, and PostgreSQL as the database.

## Features

- Simulate sending and receiving messages from customers via an API endpoint
- Store and retrieve messages from a database
- Display and respond to messages in a web interface
- Handle multiple agents and customers concurrently

## Installation

To install and run the application, you need to have Go, Gin, and PostgreSQL installed on your system. You also need to set up the database connection string as an environment variable named `DB_CONNECTION_STRING`. You can follow these steps:

- Clone the repository from GitHub:

```bash
git clone https://github.com/MikeMwita/messaging_app.git
```

- Change the directory to the project folder:

```bash
cd messaging_app
```

- Build the application and produce a binary executable file:

```bash
go build -o main .
```

- Run the application:

```bash
go run cmd/api/main.go
```

## Usage

To use the application, you can send and receive messages via the API endpoints or the web interface(Postman or Insomnia). The API endpoints are as follows:
- `GET /simulate/receive-messages`: This route returns a JSON array of messages from the customers that are stored in the database.
- `POST /simulate/send-message`: This route accepts a JSON object of a message from an agent and inserts it into the database. The JSON object should have the following fields: `user_id`, `time`, and `content`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

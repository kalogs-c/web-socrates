# Web Socrates

Web Socrates is a game server built using Golang and web sockets. It allows
players to create rooms, participate in a question-and-answer game, engage in
real-time chat, and includes authentication using JSON Web Tokens (JWT).

## Features

- Create and join rooms: Players can create their own rooms or join existing
  ones.
- Question transmission: Questions are sent to the room for players to answer.
- Answer submission: Players can submit their answers to the questions received.
- Real-time interaction: Web sockets enable real-time communication between
  players in the same room.
- Real-time chat: Players can chat with each other in real time within the game.
- Authentication with JWT: Secure player authentication using JSON Web Tokens.

## How to Use

1. Clone the repository:
   `git clone https://github.com/kalogs-c/web-socrates.git`
2. Navigate to the project directory: `cd web-socrates`
3. Start the container and entry on Go container: `make up`
4. Run the app `make dev`
5. Access the application in your browser: `http://localhost:8080`

## Dependencies

The app uses docker for containers and [sqlc](https://sqlc.dev/) to compile SQL
to Golang code

## Contributing

Contributions are welcome! If you find any issues or have suggestions for
improvements, please submit a pull request or open an issue on the
[GitHub repository](https://github.com/kalogs-c/web-socrates).

## License

This project is licensed under the [MIT License](LICENSE). Feel free to modify
and distribute it as needed.

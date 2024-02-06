# Go Search Service

This is a Go-based search service for finding books based on various criteria such as book name, author, category, etc.

## Routes

- **GET /api/search/books?bookName={bookName}**: Searches for books by name.
- **GET /api/search/books?author={author}**: Searches for books by author.
- **GET /api/search/books?category={category}**: Searches for books by category.
- **GET /api/search/books/best-sellers**: Retrieves best-selling books.
- **GET /api/search/books/language-original**: Retrieves books in the original language.
- **GET /api/search/books/popular**: Retrieves popular books.
- **GET /api/search/books/economical**: Retrieves economical books.
- **GET /api/search/books/free**: Retrieves free books.
- **GET /api/search/books/recently-added**: Retrieves recently added books.

## Usage

1. Clone the repository from GitHub.
2. Make sure Docker is installed on your machine.
3. Create a Docker network: `docker network create my_network`
4. Navigate to the project directory.
5. Use Docker Compose to start the service: `docker-compose up -d`

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

# Distributed Crawler API

This project provides a distributed crawler API. You can use this API to start crawling URLs and get information about available endpoints.

## Installation and Running

### Requirements

- Go 1.23.4 or higher
- Git

### Installation

1. Clone this repository.

    ```sh
    git clone https://github.com/yourusername/distributed-crawler.git
    cd distributed-crawler
    ```

2. Install the required packages.

    ```sh
    go mod tidy
    ```

### Running

1. Start the server.

    ```sh
    go run server/server.go
    ```

2. Once the server is running, the API will be available at [http://localhost:8080](http://_vscodecontentref_/0).

## Usage

### Endpoints

- `GET /`: Returns information about available endpoints and example usage.
- `POST /add-url`: Starts crawling a URL. The request body should include the `url` in JSON format.

### Example Requests

- Create Redis Server

    ``` bash
    docker-compose up -d 
    ```

- Get available endpoints

    ```sh
    curl http://localhost:8080/
    ```

- Start crawling a URL

    ```sh
    curl -X POST -H "Content-Type: application/json" -d '{"url":"http://example.com"}' http://localhost:8080/add-url
    ```

## File Structure

- [server.go](http://_vscodecontentref_/1): Main code file for the server.
- `server/usage.json`: JSON file describing available endpoints and example usage.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue.

## License

This project is licensed under the MIT License. See the [LICENSE](http://_vscodecontentref_/2) file for details.
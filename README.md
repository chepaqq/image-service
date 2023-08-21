## Jungle Test Task: Image Upload and Authentication Service

This is a test task for creating a service that allows users to register, authenticate, upload images, and retrieve their uploaded images.

## Endpoints

### POST /register

Register a new user with the provided credentials.

Request:

```json
{
  "username": "test",
  "password": "123456"
}
```

Response:

```json
{
  "id": 1
}
```

### POST /login

Authenticate the user and receive a JWT token for subsequent requests.

Request:

```json
{
  "username": "test",
  "password": "123456"
}
```

Response:

```json
{
  "token": "jwt string"
}
```

### POST /upload-picture

Upload an image and save its URL in the database.

Response:

```json
{
  "url": "http://0.0.0.0:9000/images/e04606e7-ff08-417d-a69c-a0bcf1c001be_mclovin.jpg"
}
```

### GET /images

Retrieve an array of all uploaded images for the authenticated user.

Response:

```json
[
  {
    "id": 4,
    "user_id": 10,
    "image_path": "mclovin.jpg",
    "image_url": "http://0.0.0.0:9000/images/e04606e7-ff08-417d-a69c-a0bcf1c001be_mclovin.jpg"
  }
]
```

## Getting Started

To run the service locally, you can use Docker Compose. The following steps assume you have Docker and Docker Compose installed.

1. Clone this repository to your local machine.

2. Create a `.env` file in the root directory based on the provided `.env.example` file. Modify the environment variables as needed.

3. Build and run the Docker containers using the following command:

   ```sh
   docker-compose up --build
   ```

4. The service should now be accessible at `http://localhost:{PORT}` where `{PORT}` is the value you specified in your `.env` file.

## Dependencies

This test task uses the following technologies and libraries:

- PostgreSQL: A relational database used for storing user information and image URLs.
- MinIO: An object storage server used for storing the uploaded images.
- Gorilla Mux: A popular HTTP router for creating clean and efficient API routes.
- Logrus: A structured logger for better logging in the application.
- Cleanenv: A library for reading environment variables into a Go struct.
- SQLx: A database toolkit for Go that provides a set of extensions on top of the standard `database/sql`.

```
# Tryout Backend API

This repository contains the backend API for the Tryout application. The API is built using Go and the Fiber web framework, SQLC for Database. It provides endpoints for managing tryouts, including creating, updating, deleting, and fetching tryout details.

## Installation

Clone the repository:

```sh
git clone https://github.com/doodpanda/tryout-backend.git
cd tryout-backend
```

Install dependencies:

```sh
go mod tidy
```

Copy the sample environment file and configure it:

```sh
cp .env.sample .env
```

Update the .env file with your database URL.

## Usage

### Running the Application

To run the application, use the following command:

```sh
make run
```

## Tryout Endpoints

**API BASE** : `/api/v1/`

- **Get Tryout List**: `GET /tryouts`
- **Get Tryout List Filtered**: `POST /tryouts`
- **Get Tryout By ID**: `GET /tryouts/:id`
- **Create New Tryout**: `POST /tryouts/new`
- **Update Tryout**: `PUT /tryouts/:id`
- **Delete Tryout**: `DELETE /tryouts/:id`
- **Login**: `POST /login`
- **Register**: `POST /register`

## License

This project is licensed under the Apache 2.0 License. See the LICENSE file for details.
```
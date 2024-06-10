# Pismo Back Challenge

## Running
Requirements: `docker`, `docker compose`.

Run: 
```sh
docker compose up --build
```

This will get the api up and running on `localhost:8080`.

## Usage
You can import the Postman collection found in the root folder [here](/pismo-challenge.postman_collection.json) to make HTTP requests.

## Development
Requirements: `go`, `docker`, `docker compose`

- Change the `CONNECTION_STRING` in the `.env` file to use the local one.
- Then run:
```
docker compose up postgres -d
go build -o api
./api
```

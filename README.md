# Hotel Reservation Backend

## Project outline
- users -> book form from a hotel
- admins -> going to check reservation/bookings
- Authentication & Authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding,  migration

## Project Environment Variables
```
HTTP_LISTEN_ADDRESS:=
JWT_SECRET=
MONGO_DB_NAME=
MONGO_DB_URL=
MONGO_DB_URL_TEST=
```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/lucaliebenberg/hotel-reservation-backend.git
   ```

2. Build the project:
  ```bash
  go build
  ```

3. Run the project:
  ```bash
  ./hotel-reservation-backend
  ```
## Contributing
If you would like to contribute to the project, please follow these steps:
  1. Fork the repository.
  2. Create a new branch (git checkout -b feature-branch).
  3. Make your changes.
  4. Commit your changes (git commit -am 'Add new feature').
  5. Push to the branch (git push origin feature-branch).
  6. Create a new Pull Request.

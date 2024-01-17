# Booking System

This repository contains the source code for a simple room booking system.
<p>The system is built using Go and PostgreSQL.</p>

## Getting Started

Follow the steps below to set up and run the application:

### Prerequisites

- Go should be installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org).
- PostgreSQL should be installed and running.

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/your-username/udemy-course-booking-system.git
```

### Database

The project uses PostgreSQL for its database. To access the PostgreSQL command-line interface, you can use the following command:

```bash
sudo -u postgres psql
```

### Migrations

The project uses [Soda](https://github.com/gobuffalo/pop) (Pop v6.1.1) for database migrations.

To run migrations, use the following command:

```bash
soda migrate
```

Change to the project directory:
```bash
cd udemy-course-booking-system
```

Build the application:
```bash
go build -o bookings cmd/web/*.go
```

### Configuration
Before starting the application, you may need to configure the database connection. Update the database connection details in the config/config.json file.

### Starting the Application
The application supports several command-line flags for configuration.
To start the application, run the following command:
```bash
./bookings -dbname=bookings -dbuser=username -dbpassword=DB-Password
```
```bash
./bookings -production=true -cache=true -dbhost=localhost -dbname=bookings -dbuser=username -dbpassword=password -dbport=5432 -dbssl=disable
```
Make sure PostgreSQL is running before starting the application.

### Usage

Visit [http://localhost:8080](http://localhost:8080) in your web browser to access the application.


### Features

User authentication
Room reservation
Template caching for improved performance
Admin Panel Features:

- **Delete Reservation:** Implement the ability to delete a reservation.
    ```go
    func DeleteReservation(id int) error
    ```

- **Update Processed for Reservation:** Update the processed status for a reservation.
    ```go
    func UpdateProcessedForReservation(id, processed int) error
    ```

- **Get Restrictions for Room by Date:** Retrieve restrictions for a room based on specified dates.
    ```go
    func GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error)
    ```

- **Delete Block for Room:** Delete a block for a room.
    ```go
    func DeleteBlockForRoom(id int) error
    ```

- **Insert Block for Room:** Insert a block for a room with a specified start date.
    ```go
    func InsertBlockForRoom(id int, startDate time.Time) error
    ```

These functions are part of the `DatabaseRepo` interface and provide the necessary functionality for managing reservations and restrictions in the admin panel.

### Project Dependencies

The project uses Go 1.21 and has the following dependencies specified in the `go.mod` file:

- **github.com/go-chi/chi/v5 v5.0.10**
- **github.com/alexedwards/scs/v2 v2.6.0** *(indirect)*
- **github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2** *(indirect)*
- **github.com/jackc/pgpassfile v1.0.0** *(indirect)*
- **github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a** *(indirect)*
- **github.com/jackc/pgx/v5 v5.5.0** *(indirect)*
- **github.com/jackc/puddle/v2 v2.2.1** *(indirect)*
- **github.com/justinas/nosurf v1.1.1** *(indirect)*
- **github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208** *(indirect)*
- **github.com/xhit/go-simple-mail/v2 v2.16.0** *(indirect)*
- **golang.org/x/crypto v0.9.0** *(indirect)*
- **golang.org/x/sync v0.1.0** *(indirect)*
- **golang.org/x/text v0.9.0** *(indirect)*




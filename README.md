# MyDrive 1.0.0

## Project Description

**MyDrive is a project built with the purpose of serving as an open-source personal cloud hosted drive.**

## Project Status

**The project is currently in DEVELOPMENT status**

For now the project is made to be run in development mode, since is not ready production
so proceed with caution when intending to use it for more than an experiment.

Moreover, the project is not yet fully tested, so there might be some bugs that need to be fixed.

## Features

- User Authentication
- JWT
- File Upload/Download
- Metrics (not yet implemented in the frontend)

## Upcoming features
- File Sharing
- File Versioning
- File Encryption
- File Search
- File Preview

## Colaboration and Contribution

If you find a bug, please open an issue.

If you want to contribute, please open a pull request.

### Technologies Used
- Go
- React
- Docker
- Postgres


### Prerequisites

- docker/docker-compose installed
- docker-desktop installed
- go installed
- node installed
- npm installed

## Setup

1. Clone the repository
2. Setup the environment variables:

In the project you cloned, create a file named `.env` and fill it with the following content:

```
POSTGRES_USER=**your_postgres_user**
POSTGRES_PASSWORD=**your_postgres_password**
POSTGRES_DB=**your_postgres_db**

DB_ADDR=postgres://{POSTGRES_USER}:{POSTGRES_PASSWORD}@localhost:5434/{POSTGRES_DB}?sslmode=disable
DB_MAX_OPEN_CONNS=30 // this is default you can change later
DB_MAX_IDLE_CONNS=30 // this is default you can change later
DB_MAX_IDLE_TIME=15m // this is default you can change later

!!! VERY IMPORTANT !!!
// Choose a secret key for the JWT token, make sure it is randomly chosen and is secured.
AUTH_TOKEN_SECRET=**your_secret_key**

!!! VERY IMPORTANT !!!
// This should be the root folder where your files will be uploaded to and downloaded from.
FILE_SYSTEM_ROOT_FOLDER=**your_root_folder**

// Here you can choose between development and production
DEV=development 
```

3. Run the following command to start the project:

```
docker-compose up --build // This will start your docker containers
air // This will start the go server (air is used for hot-reloading)
npm start dev // This will start the react server
```

4. Enjoy!
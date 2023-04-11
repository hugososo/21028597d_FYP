# Installation Guide

This guide will help you install and run the project. The project has three parts: Smart contract, Backend, Frontend

1. Clone the repo First

```
git clone <repository-url>
```

## Smart Contract

### Prerequisites

- Node.js version 16 or higher
- Ganache Installed

### Steps

1. Go to [Ganache](https://trufflesuite.com/ganache/) Download the software and start

2. node package install

   ```
   cd smart-contract
   npm i
   ```

3. Config
   Go to hardhat.config.js to config the ganache server address and private key

4. Go to scripts, copy the first two line of deploy.sh
   ```
   npx hardhat compile
   npx hardhat run --network localhost scripts/deploy.js
   ```

# Backend

### Prerequisites

- Golang installed
- Docker installed
- Infura API register

### Step

1. Download [Golang](https://go.dev/)
2. Download [Docker](https://www.docker.com/)
3. Register [Infura](https://www.infura.io/)
4. Install golang module

   ```
   cd backend
   go get .
   ```

5. Start PostgreSQL
   ```
   cd deployments
   docker-compose up
   ```
6. Config
   go to configs dir to config the variables
   Create a env folder and .env.local in it
   input the following variables
   ```
   POSTGRES_HOST=127.0.0.1
   POSTGRES_USER=postgres
   POSTGRES_PASSWORD=postgres
   POSTGRES_PORT=5432
   POSTGRES_DB_NAME=postgres
   IPFS_PRJ_ID = <Your infura id>
   IPFS_API_SECRET = <Your infura api secret>
   ```
7. Run the server
   ```
   cd cmd
   go run main.go
   ```

## Frontend

### Prerequisites

- Node.js version 16 or higher

### Step

1. Install package

   ```
   cd frontend
   npm i
   ```

2. Run the program
   ```
   npm run start
   ```

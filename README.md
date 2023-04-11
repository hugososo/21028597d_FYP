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

### Step

1. Download [Golang](https://go.dev/)
2. Download [Docker](https://www.docker.com/)
3. Install golang module

   ```
   cd backend
   go get .
   ```

4. Start PostgreSQL
   ```
   cd deployments
   docker-compose up
   ```
5. Run the server
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

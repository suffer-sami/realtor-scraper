# realtor-scraper
![banner](./assets/banner.svg)

## Overview
This project allows you to scrape all the real estate agents' information, including their contact details and stats from Realtor.com. We do this by reverse-engineering their internal API and bypassing the signing and verification of requests, thereby accessing all the data publicly.

<img alt="intro" src="./assets/intro.gif" width="600" />

We are using `libsql` to manage the database, which organizes data into various interconnected tables. The database will store crucial information, including agent details, office locations, areas of operation, sales data, and social media profiles. This structure ensures that data remains organized and easily accessible for analysis and future use.

![db_schema](./assets/db_schema.png)

## Prerequisites

Before running the project, make sure you have Go installed on your system.

### Install Go

1. Visit the official Go website: [https://golang.org/dl/](https://golang.org/dl/)
2. Download the appropriate installer for your operating system.
3. Follow the installation instructions for your platform.

To verify the installation, open your terminal or command prompt and run:

```bash
go version
```
This should return the installed version of Go.

### Clone the Project
Once Go is installed, you can clone this project to your local machine.

Run the following command in your terminal:
```bash
git clone https://github.com/suffer-sami/realtor-scraper.git
```

Navigate to the project directory:

```bash
cd realtor-scraper
```

### Install Go Dependencies
Before running the project, ensure that your Go dependencies are up to date. Run the following command to tidy up your Go modules:

```bash
go mod tidy
```
This will clean up any unnecessary dependencies and ensure everything is properly installed.

### Install Goose

This project uses Goose for database migrations. To install Goose, run the following command:


```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Run Migrations
Once Goose is installed, you can apply the migrations to set up the database.

Set the necessary environment variables:

- `GOOSE_DRIVER` — Specifies the database driver.
- `GOOSE_MIGRATION_DIR` — Directory where migration files are located.
- `GOOSE_DBSTRING` — Connection string to your database (local or remote)

Run the migrations with the following command:

```bash
# For local database
GOOSE_DRIVER="turso" GOOSE_MIGRATION_DIR="./sql/schema/" GOOSE_DBSTRING="file:./local.db" goose up

# OR

# For remote turso database
GOOSE_DRIVER="turso" GOOSE_MIGRATION_DIR="./sql/schema/" GOOSE_DBSTRING="libsql://dbname.turso.io?authToken=token" goose up
```

## Run the Project

Follow these steps to configure and run the project:

### Copy the `.env` File

Start by copying the example `.env` file to set up your environment variables:

```bash
cp .env.example .env
```

Edit the .env file to add your details, such as the database connection string, API tokens, and other necessary configuration.

### Run the Scraper
To build and run the scraper, use the following commands:

**Build**:
```bash
go build 
```

**Run**:
```bash
./realtor-scraper .

# OR

# Run with a custom number of concurrent threads (default: 3)
./realtor-scraper . 5
```

## Walkthrough
> [!NOTE]
> For this walkthrough, I'm using a 5-thread configuration to fetch 10 requests with LOG_LEVEL=INFO in a Docker container with rotating SOCKS5 proxies.

> [!CAUTION]
> **Avoid running the scraper without proxies.** If you're not using Docker w/ proxies, consider using any rotating system-level proxies or even a VPN.

![Walkthrough](./assets/demo.gif)

<details>
<summary>Detailed Log</summary>

**Note:** Below is an example of a dummy agent log entry with `LOG_LEVEL=DEBUG`.

```bash
2024/12/10 00:15:55 realtor-scraper  INFO: FETCHING: Total Results
2024/12/10 00:16:01 realtor-scraper  INFO: STATS: Total Agents: 1666411
2024/12/10 00:16:01 realtor-scraper  INFO: STATS: (Total Requests: 83321, Remaining Requests: 83320)
2024/12/10 00:16:01 realtor-scraper  INFO: FETCHING: Agents (offset 0, limit 20)
2024-12-09 14:46:25 INFO realtor-scraper Agent: Dummy Agent Name
2024-12-09 14:46:25 DEBUG realtor-scraper - sales data: 2024-11-25
2024-12-09 14:46:25 DEBUG realtor-scraper - listing data: 2024-12-08 23:59:59 +0000 UTC
2024-12-09 14:46:25 DEBUG realtor-scraper - social medias:
2024-12-09 14:46:25 DEBUG realtor-scraper - feed licences:
2024-12-09 14:46:25 DEBUG realtor-scraper - mls:
2024-12-09 14:46:25 DEBUG realtor-scraper   - MLS1
2024-12-09 14:46:25 DEBUG realtor-scraper   - MLS2
2024-12-09 14:46:25 DEBUG realtor-scraper - mls history:
2024-12-09 14:46:25 DEBUG realtor-scraper - languages:
2024-12-09 14:46:25 DEBUG realtor-scraper   - English
2024-12-09 14:46:25 DEBUG realtor-scraper   - Spanish
2024-12-09 14:46:25 DEBUG realtor-scraper - user languages:
2024-12-09 14:46:25 DEBUG realtor-scraper   - English
2024-12-09 14:46:25 DEBUG realtor-scraper   - Spanish
2024-12-09 14:46:25 DEBUG realtor-scraper - zips:
2024-12-09 14:46:25 DEBUG realtor-scraper   - 12345
2024-12-09 14:46:25 DEBUG realtor-scraper   - 67890
2024-12-09 14:46:25 DEBUG realtor-scraper   - 98765
2024-12-09 14:46:25 DEBUG realtor-scraper - areas:
2024-12-09 14:46:25 DEBUG realtor-scraper   - Cityville, CA
2024-12-09 14:46:25 DEBUG realtor-scraper   - Countyville, TX
2024-12-09 14:46:25 DEBUG realtor-scraper - marketing areas:
2024-12-09 14:46:25 DEBUG realtor-scraper   - Neighborhood1, Cityville, CA
2024-12-09 14:46:25 DEBUG realtor-scraper   - Neighborhood2, Countyville, TX
2024-12-09 14:46:25 DEBUG realtor-scraper - designations:
2024-12-09 14:46:25 DEBUG realtor-scraper - specializations:
2024-12-09 14:46:25 DEBUG realtor-scraper   - Residential Sales
2024-12-09 14:46:25 DEBUG realtor-scraper   - Commercial Leasing
2024-12-09 14:46:25 DEBUG realtor-scraper - address:
2024-12-09 14:46:25 DEBUG realtor-scraper   - {Street: 123 Main St, City: Anytown, State: CA, Zip: 98765}
2024-12-09 14:46:25 DEBUG realtor-scraper - phones:
2024-12-09 14:46:25 DEBUG realtor-scraper   - XXX-XXX-XXXX
2024-12-09 14:46:25 DEBUG realtor-scraper   - XXX-XXX-XXXX
2024-12-09 14:46:25 DEBUG realtor-scraper - broker:
2024-12:09 14:46:25 DEBUG realtor-scraper   - Dummy Brokerage
2024-12-09 14:46:25 DEBUG realtor-scraper - office address:
2024-12-09 14:46:25 DEBUG realtor-scraper   - {Street: 456 Elm St, City: Anytown, State: CA, Zip: 98765}
2024-12-09 14:46:25 DEBUG realtor-scraper - office:
2024-12-09 14:46:25 DEBUG realtor-scraper   - Dummy Brokerage
2024-12-09 14:46:25 DEBUG realtor-scraper - office phones:
2024-12-09 14:46:25 DEBUG realtor-scraper   - {Ext: Number: XXX-XXX-XXXX, Type: Office, IsValid: true}
2024-12-09 14:46:25 DEBUG realtor-scraper   - {Ext: Number: XXX-XXX-XXXX, Type: Mobile, IsValid: true}
# (cont... )

2024-12-09 14:46:25 INFO realtor-scraper DONE
```
</details>
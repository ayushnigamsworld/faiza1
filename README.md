# Synctera Tech Challenge

## Problem statement
The starter server exposed a `/transactions` endpoint that returned hard-coded mock data from memory. The challenge asks for a production-like shape where transaction data is ingested from a JSON file chosen at runtime, PANs are masked to the last four digits in all responses, transactions can be returned ordered by descending posted timestamp, tests accompany the behavior, and documentation explains how to build, test, and run the service.

## What changed and why
- Added JSON file ingestion so the API can be pointed at external mock transaction data via `--transactions <file>`.
- Implemented consistent PAN masking to return only the last four digits and protect sensitive card information.
- Added a `/transactions/posted-desc` endpoint to provide transactions sorted by `posted_timestamp` in descending order.
- Wrote unit tests for the masking logic and both transaction endpoints to guard the behavior.
- Documented build, test, and run steps for clarity.

## How this differs from real-world services
- The service is intentionally minimal and runs in a single file; a production service would likely include structured logging, validation, layered architecture, error handling contracts, observability hooks, configuration management, and persistent storage.
- Authentication and authorization are omitted; real APIs would protect transaction data.
- The JSON file acts as a mock data store instead of a database or downstream service.

## Build, test, and run
### Build
```bash
make build
```

### Test
```bash
make test
```

### Run
Use the default mock data file:
```bash
./main
```

Or point to a specific data source:
```bash
./main --transactions path/to/transactions.json
```

Then query the API:
```bash
curl http://localhost:8000/transactions
curl http://localhost:8000/transactions/posted-desc
```

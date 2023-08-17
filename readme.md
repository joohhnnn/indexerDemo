# indexerDemo

## Overview

This project is designed to index ERC20 transfer events from the Ethereum blockchain. It provides functionalities for connecting to the Ethereum network, interacting with ERC20 contracts, and storing events in a database.

## Directory Structure

- `cmd`: Contains the main entry point of the application.
- `config`: Houses configuration-related code and tests.
- `indexer`: The core part of the project, including various components:
  - `database`: Database connection and operations, including schema definition.
  - `types`: Definitions of custom types used in the project.
  - `abi`: Ethereum contract ABI definitions.
  - `api`: API server and routes.
  - `ethereum`: Code related to Ethereum interactions, such as event indexing and connection handling.
- `Integration_test_unfinished`: Contains unfinished integration tests and related files.



## TODO

1. **Integration Testing**: Implement integration tests that simulate the Ethereum environment to thoroughly test the application.
2. **Concurrency Optimization for Indexing Past Events**: Enhance the performance of indexing past events by introducing concurrency optimizations.
3. **Filtering Spam Transfer Events**: Investigate and implement a mechanism to filter out spam transfer events (e.g., events that only send without actual transfer).


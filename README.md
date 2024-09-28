DiceDB
===

DiceDB is an in-memory real-time database with SQL-based reactivity. It is hyper-optimized for building and scaling truly real-time applications on modern hardware while being a drop-in replacement for Redis.

We are looking for Early Design Partners, so, if you want to evaluate DiceDB, [block our calendar](https://cal.com/dicedb-arpit). always up for a chat.

> Note: DiceDB is still in development and it supports a subset of Redis commands. So, please do not use it in production. But, feel free to go through the [open issues](https://github.com/DiceDB/dice/issues) and contribute to help us speed up the development.

## Want to contribute?

We have multiple repositories where you can contribute. So, as per your interest, you can pick one and build a deeper understanding of the project on the go.

- [dicedb/docs](https://github.com/dicedb/docs) for [documentation and blogs](https://dicedb.io). a great way to build initial understanding / Stack - Go
- [dicedb/dice](https://github.com/dicedb/dice) for core database features and engine / Stack - Go
- [dicedb/playground-mono](https://github.com/dicedb/playground-mono) backend APIs for DiceDB playground / Stack - Go
- [dicedb/playground-web](https://github.com/dicedb/playground-web) frontend for DiceDB playground / Stack - NextJS

## How is it different from Redis?

Although DiceDB is a drop-in replacement of Redis, which means almost no learning curve and switching does not require any code change, it still differs in two key aspects and they are

1. DiceDB is multi-threaded and follows [shared-nothing architecture](https://en.wikipedia.org/wiki/Shared-nothing_architecture).
2. DiceDB supports a new command called `QWATCH` that lets clients listen to a SQL query and get notified in real-time whenever something changes.

With this, you can build truly real-time applications like [Leaderboard](https://github.com/DiceDB/dice/tree/master/examples/leaderboard-go) with simple SQL query.

![Leaderboard with DiceDB](https://github.com/user-attachments/assets/327792c7-d788-47d4-a767-ef2c478d75cb)

## Get started

### Using Docker

The easiest way to get started with DiceDB is using [Docker](https://www.docker.com/) by running the following command.

```
$ docker run dicedb/dicedb
```

The above command will start the DiceDB server running locally on the port `7379` and you can connect
to it using DiceDB CLI and SDKs, or even Redis CLIs and SDKs.

> Note: Given it is a drop-in replacement of Redis, you can also use any Redis CLI and SDK to connect to DiceDB.

### Multi-Threading Mode (Experimental)

Multi-threading is currently under active development. To run the server with multi-threading enabled, follow these steps:
```shell
$ git clone https://github.com/dicedb/dice
$ cd dice
$ go run main.go --enable-multithreading=true
```

**Note:** Only the following commands are optimised for multi-threaded execution: `PING, AUTH, SET, GET, GETSET, ABORT`

### Setting up DiceDB from source for development and contributions

To run DiceDB for local development or running from source, you will need

1. [Golang](https://go.dev/)
2. Any of the below supported platform environment:
    1. [Linux based environment](https://en.wikipedia.org/wiki/Comparison_of_Linux_distributions)
    2. [OSX (Darwin) based environment](https://en.wikipedia.org/wiki/MacOS)
    3. WSL under Windows

```
$ git clone https://github.com/dicedb/dice
$ cd dice
$ go run main.go
```

4. Install GoLangCI

```
$ sudo su
$ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /bin v1.60.1
```

### Live Development Server

DiceDB provides a hot-reloading development environment, which allows you to instantly view your code changes in a live server. This functionality is supported by [Air](https://github.com/air-verse/air)

To Install Air on your system you have following options.

1. If you're on go 1.22+
```sh
go install github.com/air-verse/air@latest
```

2. Install the Air binary
```sh
# binary will be installed at $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

Once `air` is installed you can verify the installation using the command `air -v`

To run the live DiceDB server for local development:

```sh
$ git clone https://github.com/dicedb/dice
$ cd dice
$ DICE_ENV=dev air
```

> The `DICE_ENV` environment variable is used set the environment, by default it is treated as production. `dev` is used to get pretty printed logs and lower log level.

### Local Setup with Custom Config

By default, DiceDB will look for the configuration file at `/etc/dice/config.toml`. (Linux, Darwin, and WSL)

```sh
$ # set up configuration file # (optional but recommended)
$ sudo mkdir -p /etc/dice
$ sudo chown root:$USER /etc/dice
$ sudo chmod 775 /etc/dice # or 777 if you are the only user
$ git clone https://github.com/DiceDB/dice.git
$ cd dice
$ go run main.go -init-config

```

#### For Windows Users:
If you're using Windows, it is recommended to use Windows Subsystem for Linux (WSL) or WSL 2 to run the above commands seamlessly in a Linux-like environment.

Alternatively, you can:

Create a directory at `C:\ProgramData\dice` and run the following command to generate the configuration file:
```bash
go run main.go -init-config
```
For a smoother experience, we highly recommend using WSL.

#### Additional Configuration Options:

If you'd like to use a different location, you can specify a custom configuration file path with the `-c flag`:

```bash
go run main.go -c /path/to/config.toml
```
If you'd like to output the configuration file to a specific location, you can specify a custom output path with the `-o flag`:

```bash
go run main.go -o /path/of/output/dir
```

## Setting up CLI

The best way to connect to DiceDB is using DiceDB CLI and you can install it by running the following command.

```
$ pip install dicedb-cli
```

> Because DiceDB speaks Redis dialect, you can connect to it with any Redis Client and SDK also.
> But if you are planning to use the `QWATCH` feature then you need to use the DiceDB CLI.

## Running Tests

Unit tests and integration tests are essential for ensuring correctness and in the case of DiceDB, both types of tests are available to validate its functionality.

For unit testing, you can execute individual unit tests by specifying the name of the test function using the `TEST_FUNC` environment variable and running the `make unittest-one` command. Alternatively, running `make unittest` will execute all unit tests.

### Executing one unit test

```
$ TEST_FUNC=<name of the test function> make unittest-one
$ TEST_FUNC=TestByteList make unittest-one
```

### Running all unit tests

```
$ make unittest
```

Integration tests, on the other hand, involve starting up the DiceDB server and running a series of commands to verify the expected end state and output. To execute a single integration test, you can set the `TEST_FUNC` environment variable to the name of the test function and run `make test-one`. Running `make test` will execute all integration tests.

### Executing a single integration test

```
$ TEST_FUNC=<name of the test function> make test-one
$ TEST_FUNC=TestSet make test-one
```

### Running all integration tests

```
$ make test
```

> Work to add more tests in DiceDB is in progress and we will soon port the
> test [Redis suite](https://github.com/redis/redis/tree/f60370ce28b946c1146dcea77c9c399d39601aaa) to this codebase to ensure full compatibility.

## Running Benchmark

```sh
$ go test -test.bench <pattern>
$ go test -test.bench BenchmarkListRedis -benchmem
```

## Getting Started

To get started with building and contributing to DiceDB, please refer to the [issues](https://github.com/DiceDB/dice/issues) created in this repository.

## The story

DiceDB started as a re-implementation of Redis in Golang and the idea was to - build a DB from scratch and understand the micro-nuances that come with its implementation. The database does not aim to replace Redis, instead, it will fit in and optimize itself for multi-core computations running on a single-threaded event loop.

## How to contribute

The Code Contribution Guidelines are published at [CONTRIBUTING.md](CONTRIBUTING.md); please read them before you start making any changes. This would allow us to have a consistent standard of coding practices and developer experience.

Contributors can join the [Discord Server](https://discord.gg/6r8uXWtXh7) for quick collaboration.

## Contributors

<a href = "https://github.com/dicedb/dice/graphs/contributors">
  <img src = "https://contrib.rocks/image?repo=dicedb/dice"/>
</a>

## Troubleshoot

### Forcefully killing the process

```
$ sudo netstat -atlpn | grep :7379
$ sudo kill -9 <process_id>
```

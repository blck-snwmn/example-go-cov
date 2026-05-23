# example-go-cov

A Go code coverage measurement sample. This project demonstrates how to measure test coverage in Go and automate it with GitHub Actions.

## How to Get Coverage

### Basic Coverage Check

To display the coverage rate for each package, run the following command:

```bash
go test ./... -cover
```

This command displays the coverage percentage for each package individually.

### Generating Detailed Coverage Reports

To generate a more detailed coverage report, use the following command:

```bash
go test ./... -coverprofile=coverage.out
```

This command generates a `coverage.out` file.

### Checking Coverage by Function

To display coverage information for each function from the generated coverage profile:

```bash
go tool cover -func=coverage.out
```

This command displays the coverage rate for each function and shows the total statement coverage for the entire project at the end.

## Development

CLI tools (`lefthook`) are managed by [aqua](https://aquaproj.github.io/) with versions pinned in [aqua.yaml](aqua.yaml).

### Install tools

Install aqua itself first (see the [aqua installation guide](https://aquaproj.github.io/docs/install)), then install the pinned tools:

```bash
aqua install
```

### Set up git hooks

[lefthook](lefthook.yml) runs coverage tests on staged `*.go` files before each commit. Register the hooks once after cloning:

```bash
lefthook install
```

## License

MIT

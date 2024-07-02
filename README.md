# automation-api-godog

### Requirements
- Go https://go.dev/doc/install
  ```sh
  $ brew install go
  ```
- Set your environment
  ```sh
  export PATH=$PATH:/usr/local/go/bin
  ```
- Reload your shell configuration
  ```
  source ~/.zshrc
  ```

### Getting started
#### Step 1 - Setup a go module
- Go to `automation-api-godog` by running `cd automation-api-godog`
- Initiate the go module inside the directory by running `go mod init automation-api-godog`

#### Step 2 - Run the test
Run the test in the directory to run the steps you have defined
```sh
go test -v main_test.go
```

### Directory Structure
        .
        ├── features/ # Feature files (test scenarios)
        │  ├── sample.feature # Example feature file
        │  └── ...
        │
        ├── go.mod
        ├── go.sum
        ├── godog_report.json
        ├── main_test.go
        └── main.go

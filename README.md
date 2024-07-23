# automation-api-godog

### Requirements
- Go https://go.dev/doc/install
  ```sh
  $ brew install go
  $ go install github.com/cucumber/godog/cmd/godog@latest
  ```
- Set your environment
  ```sh
  export PATH=$PATH:/usr/local/go/bin
  export PATH=$PATH:$(go env GOPATH)/bin'
  ```
- Reload your shell configuration
  ```sh
  $ source ~/.zshrc
  ```
- Check your godog version
  ```
  $ godog version
  ```

### Getting started
#### Step 1 - Setup a go module
- Go to `automation-api-godog` by running `cd automation-api-godog`
- Initiate the go module inside the directory by running `go mod init automation-api-godog`

#### Step 2 - Run the test
Run the test in the directory to run the steps you have defined
```sh
$ make tests
```
After `report.json` file generated in reports folder, you can run this command to see the report
```sh
$ npm run report
```
![Screen Shot 2024-07-23 at 17 32 19](https://github.com/user-attachments/assets/f630dc88-ec1a-440e-ae6e-1f15e51254c4)

### Directory Structure
        .
        ├── features
        │  ├── init.feature
        │  └── ...
        ├── helpers
        │  ├── generate_report.js
        │  └── ...
        ├── reports
        │  ├── report.json
        │  └── report.html
        ├── steps
        │  ├── steps.go
        │  └── ...
        ├── main_test.go
        ├── package.json
        ├── README.md
        └── Makefile

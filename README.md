# automation-api-godog
![Screen Shot 2024-07-23 at 17 35 19](https://github.com/user-attachments/assets/13bc6948-6862-4987-b2ef-5ce2c6dc2a9a)

### Overview
Welcome to the automation-api-godog project! This repository provides a robust framework for API testing using the godog library, enabling Behavior-Driven Development (BDD) with Gherkin syntax.

### Requirements
- Go Language: Install Go by following the instructions https://go.dev/doc/install
  ```sh
  $ brew install go
  $ go install github.com/cucumber/godog/cmd/godog@latest
  ```
- Environment Setup: Ensure your Go and Godog paths are correctly set
  ```sh
  export PATH=$PATH:/usr/local/go/bin
  export PATH=$PATH:$(go env GOPATH)/bin'
  ```
- Reload Shell Configuration: Apply the changes to your shell environment
  ```sh
  $ source ~/.zshrc
  ```
- Verify Godog Installation: Check your godog version to confirm installation
  ```
  $ godog version
  ```
  <img width="289" alt="Screen Shot 2024-07-23 at 17 36 09" src="https://github.com/user-attachments/assets/87992a4b-91ef-44dd-b7d6-d800eebb8cb4">

### Getting started
#### Step 1: Setup a go module
Navigate to your project directory and initialize a Go module
```sh
$ cd automation-api-godog
$ go mod init automation-api-godog
```

#### Step 2: Run the test
Execute the tests using the Makefile to run the defined steps
```sh
$ make tests
```
<img width="484" alt="Screen Shot 2024-07-23 at 17 36 34" src="https://github.com/user-attachments/assets/c80d25e7-82cb-4f13-8c47-447cdcabeeb4">

After `report.json` file generated in reports folder, you can run this command to see the report
```sh
$ npm run report
```
<img width="220" alt="Screen Shot 2024-07-23 at 17 37 00" src="https://github.com/user-attachments/assets/60ca6489-6038-4851-a592-5c19a0418f22">

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
        ├── go.mod
        ├── go.sum
        ├── README.md
        └── Makefile

This project structure organizes your feature files, step definitions, and reports to maintain clarity and ease of navigation.

### Conclusion
Follow these steps to set up, run, and generate reports for your API tests using the automation-api-godog framework. Happy testing!

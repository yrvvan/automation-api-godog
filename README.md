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
<img width="289" alt="Screen Shot 2024-07-23 at 17 36 09" src="https://github.com/user-attachments/assets/87992a4b-91ef-44dd-b7d6-d800eebb8cb4">

### Getting started
#### Step 1 - Setup a go module
- Go to `automation-api-godog` by running `cd automation-api-godog`
- Initiate the go module inside the directory by running `go mod init automation-api-godog`

#### Step 2 - Run the test
Run the test in the directory to run the steps you have defined
```sh
$ make tests
```
<img width="484" alt="Screen Shot 2024-07-23 at 17 36 34" src="https://github.com/user-attachments/assets/c80d25e7-82cb-4f13-8c47-447cdcabeeb4">

After `report.json` file generated in reports folder, you can run this command to see the report
<img width="220" alt="Screen Shot 2024-07-23 at 17 37 00" src="https://github.com/user-attachments/assets/60ca6489-6038-4851-a592-5c19a0418f22">
```sh
$ npm run report
```
![Screen Shot 2024-07-23 at 17 35 19](https://github.com/user-attachments/assets/37d42679-d090-4107-8ebe-574f4e64388c)

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

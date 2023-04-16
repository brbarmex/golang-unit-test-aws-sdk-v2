# Golang: Test unit using AWS SDK for GO.

<img src="/docs/red-green-refactor.png"  width="550" height="350">

This repository contains a Go project that shows an example of how to write unit tests for AWS SDK V2 operations in Go.

The project uses the AWS SDK V2 for Go and the testing framework testing to write unit tests for operations that interact with AWS services.

### Installation
To use this project, you need to have Go installed on your machine. If you don't have it already, you can download it from the official Go website: https://golang.org/dl/

Once you have Go installed, you can clone this repository using the following command:

```bash
git clone https://github.com/brbarmex/golang-unit-test-aws-sdk-v2.git
```

### Usage
After cloning the repository, you can run the tests using the following command:

```bash
go test -v
```

This command will run all the tests in the project and print the output to the console.

To run the application locally use docker using the following command:

```bash
docker-compose -f ./docker/docker-compose -up -d
```
This command will up compose after that build the infra component using the shell script `local-stack.sh`:

```bash
./docker/local-stack.sh
```


### Contributing

If you find a bug in the project or would like to contribute to it, you can create a pull request on the repository. Please make sure to include a detailed description of the bug or feature you are adding, as well as any relevant test cases.

### License

This project is licensed under the MIT License. See the LICENSE file for more information.

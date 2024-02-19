## Table of Contents

- [Description](#description)
- [Installation](#installation)
- [Build the application](#build-the-application)
- [Run the application](#run-the-application)
- [Tests](#tests)

## Description

The application takes an input file containing a set of integers, both negative and positive,  
arranged line by line.  

The output of the application includes:

1. The maximum value in the file.
2. The minimum value in the file.
3. The median.
4. The arithmetic mean.

The longest sequence of consecutive numbers:
5. Increasing sequence.
6. Decreasing sequence.

In case there are multiple sequences of the same length, the application outputs all of them.


## Installation

To get started, clone the repository:

```bash
git clone https://github.com/Riznyk01/numeric-analyzer.git
cd numeric-analyzer
```

## Build the application:

For Linux:
```bash
make build
```

For Windows:
```bash
make buildwin
```

If make isn't installed on the Windows OS, it can be installed using the Chocolatey packet manager.  
Run Windows PowerShell with admin rights and type:
```
choco install make
```

If Chocolatey is not installed on the Windows OS, follow the [instructions](https://chocolatey.org/install) (installs in one step).

## Run the application:

To run the application, use the following command in the application's directory:
```bash
./numan -path=path-to-file
```

## Tests

To run tests, use the following command:

```bash
make test
```
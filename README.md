===========================================
oyzem - Simple Memoization Library for Go
===========================================

[![License](https://img.shields.io/badge/License-BSD--3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![Open Source](https://img.shields.io/badge/Open%20Source-Yes-green.svg)](https://opensource.org/)

## How to Install

Install `oyzem` using `go get`:

```bash
go get -u github.com/simplyYan/oyzem
```
## Introduction

oyzem is a lightweight and easy-to-use Go library for memoization, providing a simple mechanism to cache function results and improve performance.
Features

- Easy: Simple API for memoizing functions.
- Fast: Efficient caching using a mutex for thread safety.
- Lightweight: Minimalistic design with a focus on simplicity.
- Objective: Aims to provide a clear and straightforward memoization solution.

## License

oyzem is distributed under the BSD-3-Clause License. See the LICENSE file for details.

## Example
```go
package main

import (
	"fmt"
	"github.com/simplyYan/oyzem"
)

func main() {
    // Create a new Memoizer instance
    memoizer := oyzem.New()

    // Memoize a function
    memoizedFn, _ := memoizer.Memoize(func(a, b int) int {
        fmt.Println("Performing calculation...")
        return a + b
    })

    // Run the memoized function
    result, _ := memoizer.Run(memoizedFn, 2, 3)
    fmt.Println("Result:", result)
}
```
## How to Contribute

Contributions to oyzem are welcome! If you want to add, fix, or improve features, follow these steps:

1. Fork the repository on GitHub.
2. Clone your forked repository: git clone https://github.com/your-username/oyzem.git.
3. Create a new branch: git checkout -b feature-name.
4. Make your changes and commit: git commit -m "Description of changes".
5. Push your branch to GitHub: git push origin feature-name.
6. Open a pull request on the official oyzem repository.

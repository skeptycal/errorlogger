# errorlogger

> Tricky and fun utilities for Go programs.

---

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/skeptycal/errorlogger/Go) ![Codecov](https://img.shields.io/codecov/c/github/skeptycal/errorlogger)

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](code-of-conduct.md)

![Twitter Follow](https://img.shields.io/twitter/follow/skeptycal.svg?label=%40skeptycal&style=social) ![GitHub followers](https://img.shields.io/github/followers/skeptycal.svg?style=social)

---

## Getting Started

### Prerequisites

-   [Go](https://golang.org/) - I recommend the latest stable release. At least 1.18 is required for generic features.
-   [Git](https://git-scm.com/) - I recommend the latest stable release.
-   [GitHub CLI](https://cli.github.com/) - you can use git on the command line if you wish, but the github cli is very nice ...

Developed with Go 1.16.4. Updated to Go 1.18. Go is _extremely_ backwards compatible and semver stable. Nearly any v1.x should work fine.

---

### Installation



To download the repo and use it as a dependency in your module:

```bash
#
go get github.com/skeptycal/errorlogger
```

To build and install this package as a utility on your machine so you can use it from the command line:

```bash

$ gh repo clone skeptycal/errorlogger

$ cd errorlogger

# test results and coverage info
./go.test.sh

go build -race

# install as a utility package
go install
```

### Use Template

To use this repo as a template for your own project:

```sh
gh repo create -y --public --template "https://github.com/skeptycal/errorlogger"
```

### Contributing

To clone this repo to test and contribute, use `git clone` or `gh clone`.

```sh
gh repo clone skeptycal/errorlogger

cd errorlogger

git branch -b dev

```

Use the [Issues][issues] and [PR][pr] templates on the GitHub repo page to contribute.

---

### Basic Usage

> This is a copy of the example script available in the `cmd/example/errorlogger` folder:

```go
package main

import "github.com/skeptycal/errorlogger"

func main() {
    errorlogger.Example()
}

```

To try it out:

```sh
# change to the sample folder
cd cmd/example/errorlogger

# run the main.go program
go run ./main.go

# to compile as an executable
go build
```

---

## Code of Conduct and Contributing

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us. Please read the [Code of Conduct](CODE_OF_CONDUCT.md) for details before submitting anything.

---

## Versioning

We use SemVer for versioning. For the versions available, see the tags on this repository.

---

## Contributors and Inspiration

-   Michael Treanor ([GitHub][github] / [Twitter][twitter]) - Initial work, updates, maintainer
-   [Francesc Campoy][campoy] - Inspiration and great YouTube videos!

See also the list of contributors who participated in this project.

---

## License

Licensed under the MIT <https://opensource.org/licenses/MIT> - see the [LICENSE](LICENSE) file for details.

[twitter]: (https://www.twitter.com/skeptycal)
[github]: (https://github.com/skeptycal)
[campoy]: (https://github.com/campoy)
[fatih]: (https://github.com/fatih/color)
[issues]: (https://github.com/skeptycal/errorlogger/issues)
[pr]: (https://github.com/skeptycal/errorlogger/pulls)

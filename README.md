<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# algorithmsAndDataJunctionStructures

A Go workspace of textbook data structure and sorting algorithm implementations built while following the imooc Play With Data Structures course, plus one LeetCode solution, each covered by its own unit tests.

**English** · [简体中文](README.zh-CN.md)

[![CI](https://github.com/anyingiit/algorithmsAndDataJunctionStructures/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/algorithmsAndDataJunctionStructures/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/algorithmsAndDataJunctionStructures)](LICENSE)

[Report a bug](https://github.com/anyingiit/algorithmsAndDataJunctionStructures/issues/new?template=bug_report.yml) · [Request a feature](https://github.com/anyingiit/algorithmsAndDataJunctionStructures/issues/new?template=feature_request.yml)

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About The Project

algorithmsAndDataJunctionStructures is a Go workspace of data structure and algorithm exercises, kept under `imooc/playWithDataStructures` after the imooc "Play With Data Structures" course it was written alongside: a dynamic array, a singly linked list, a stack and a queue (both slice- and link-backed), a binary search tree, a max-heap, a priority queue, a union-find structure, a hash set, and three sorting algorithms (bubble sort, merge sort, heap sort). Several structures are kept in multiple numbered versions — for example `UnionFind/QuickUnionV1` through `QuickUnionV5`, and `BST/versions/V4` — to show the same idea refined step by step, and `leetcode/DeleteOrderArrayRepeats` holds one solved LeetCode problem. Nearly every package ships its own `_test.go` file, and `go test ./...` exercises all of them.

See the [open issues](https://github.com/anyingiit/algorithmsAndDataJunctionStructures/issues) for planned features and known issues.

## Getting Started

### Prerequisites

- Go 1.18 or newer, the version pinned in `go.mod`
- Nothing else: the module has no `go.sum` and every import in the tree resolves to the Go standard library

### Installation

```sh
git clone https://github.com/anyingiit/algorithmsAndDataJunctionStructures.git
cd algorithmsAndDataJunctionStructures
go build ./...
```

## Usage

Run the whole test suite, or a single example program directly:

```sh
go test ./...
go run ./imooc/playWithDataStructures/LinkedList/entry
```

Most data structures have a matching `entry/entry.go` program under their
package (for example `imooc/playWithDataStructures/Queue/entry`,
`imooc/playWithDataStructures/Stack/entry`) that builds and exercises it the
same way.

## Contributing

Contributions are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md) for how to open an issue or a pull request, and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for the standards expected of everyone taking part.

Please do not report security issues in public issues or pull requests. [SECURITY.md](SECURITY.md) explains how to report them privately.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.

## Contact

Project link: [https://github.com/anyingiit/algorithmsAndDataJunctionStructures](https://github.com/anyingiit/algorithmsAndDataJunctionStructures)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

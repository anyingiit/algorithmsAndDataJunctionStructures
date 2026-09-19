[English](README.md) · **简体中文**

> 英文版是规范版本。本页与 [README.md](README.md) 不一致时，以英文版为准。

<!-- translation-of: README.md sha256:e5c6c6afbc797798 -->

<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# algorithmsAndDataJunctionStructures

一个 Go 语言的数据结构与排序算法练习仓库，是跟随 imooc《玩转数据结构》课程编写的教科书式实现，外加一道 LeetCode 题解，每个包都配有自己的单元测试。

[![CI](https://github.com/anyingiit/algorithmsAndDataJunctionStructures/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/algorithmsAndDataJunctionStructures/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/algorithmsAndDataJunctionStructures)](LICENSE)

[报告问题](https://github.com/anyingiit/algorithmsAndDataJunctionStructures/issues/new?template=bug_report.yml) · [提出需求](https://github.com/anyingiit/algorithmsAndDataJunctionStructures/issues/new?template=feature_request.yml)

<details>
  <summary>目录</summary>
  <ol>
    <li><a href="#about-the-project">关于本项目</a></li>
    <li><a href="#getting-started">开始使用</a></li>
    <li><a href="#usage">用法</a></li>
    <li><a href="#contributing">参与贡献</a></li>
    <li><a href="#license">许可证</a></li>
    <li><a href="#contact">联系方式</a></li>
  </ol>
</details>

## 关于本项目

algorithmsAndDataJunctionStructures 是一个 Go 语言的数据结构与算法练习仓库，代码放在
`imooc/playWithDataStructures` 目录下，这个目录名来自编写时跟随的 imooc《玩转数据结构》课程：
包含动态数组、单向链表、栈与队列（分别用切片和链表实现）、二分搜索树、大顶堆、优先队列、并查集、
哈希集合，以及三种排序算法（冒泡排序、归并排序、堆排序）。其中一些结构保留了多个编号版本——例如
`UnionFind/QuickUnionV1` 到 `QuickUnionV5`，以及 `BST/versions/V4`——用来展示同一个思路是如何被
逐步改进的；`leetcode/DeleteOrderArrayRepeats` 则保存了一道已解决的 LeetCode 题目。几乎每个包都
自带 `_test.go` 测试文件，`go test ./...` 会把它们全部跑一遍。

计划中的功能与已知问题，见 [open issues](https://github.com/anyingiit/algorithmsAndDataJunctionStructures/issues)。

## 开始使用

### 环境要求

- Go 1.18 或更高版本，即 `go.mod` 中固定的版本
- 不需要其他依赖：该模块没有 `go.sum`，代码树中的每一处导入都指向 Go 标准库

### 安装

```sh
git clone https://github.com/anyingiit/algorithmsAndDataJunctionStructures.git
cd algorithmsAndDataJunctionStructures
go build ./...
```

## 用法

运行整个测试套件，或者直接运行某一个示例程序：

```sh
go test ./...
go run ./imooc/playWithDataStructures/LinkedList/entry
```

大多数数据结构在自己的包下都配有一个对应的 `entry/entry.go` 示例程序（例如
`imooc/playWithDataStructures/Queue/entry`、`imooc/playWithDataStructures/Stack/entry`），
构建和运行方式与上面相同。

## 参与贡献

欢迎参与。[CONTRIBUTING.md](CONTRIBUTING.md) 说明如何提交 issue 或 pull request，[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 说明对所有参与者的行为要求。

请不要在公开的 issue 或 pull request 中报告安全问题。[SECURITY.md](SECURITY.md) 说明了私下报告的方式。

## 许可证

以 MIT 许可证分发。详见 [LICENSE](LICENSE)。

## 联系方式

项目地址：[https://github.com/anyingiit/algorithmsAndDataJunctionStructures](https://github.com/anyingiit/algorithmsAndDataJunctionStructures)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

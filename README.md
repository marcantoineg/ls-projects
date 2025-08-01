# 📚 ls-projects &nbsp; <img height="20px" src="https://img.shields.io/badge/Golang-FFFFFF?logo=go&style=flat">

![ci badge](https://github.com/marcantoineg/ls-projects/actions/workflows/ci.yml/badge.svg)
[![Release](https://github.com/marcantoineg/ls-projects/actions/workflows/release.yml/badge.svg)](https://github.com/marcantoineg/ls-projects/actions/workflows/release.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A simple Go app to list projects and open them in a new window of VS Code.

This app uses [Bubble Tea](https://github.com/charmbracelet/bubbletea/) as its UI framework.

<img src=".github/demo/home-demo.gif" width="100%" align="center" />

## Configuration

The CLI will read the file at `~/.config/ls-projects/.project.json`, don't forget to copy your config file if you edited the one provided in this project.

## Usage
You can either use it with `go run main.go` or by exporting it as a binary in your PATH using `go build -o <binary_path/binary_name>`.

## Motivation
<img src="https://user-images.githubusercontent.com/16008095/208336763-22bec39c-6a44-4469-96bc-675b0f2e85de.png" />

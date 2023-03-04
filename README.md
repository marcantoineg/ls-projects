# 📚 ls-projects

![ci badge](https://github.com/marcantoineg/list-my-projects/actions/workflows/ci.yml/badge.svg)
[![Release](https://github.com/marcantoineg/list-my-projects/actions/workflows/release.yml/badge.svg)](https://github.com/marcantoineg/list-my-projects/actions/workflows/release.yml)
<img height="20px" src="https://img.shields.io/badge/Golang-FFFFFF?logo=go&style=flat">

A simple Go app to list projects and open them in a new window of VS Code.

This app uses [Bubble Tea](https://github.com/charmbracelet/bubbletea/) as its UI framework.

<img src=".github/demo/home-demo.gif" width="100%" align="center" />

<p align="right"><em><sub>without a meme, is it really a nice repo?</sub></em></p>
<img src="https://user-images.githubusercontent.com/16008095/208336763-22bec39c-6a44-4469-96bc-675b0f2e85de.png" align="right" width="37%" />

## Configuration

The CLI will read the file at `~/.config/list-my-project/.project.json`, don't forget to copy your config file if you edited the one provided in this project.

## Usage
You can either use it with `go run main.go` or by exporting it as a binary in your PATH using `go build -o <binary_path/binary_name>`.

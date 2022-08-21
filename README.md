# list-my-projects
![ci badge](https://github.com/marcantoineg/list-my-projects/actions/workflows/ci.yml/badge.svg)

A simple shell app to list projects and open them in a new window of VS Code.

This app uses [Bubble Tea](https://github.com/charmbracelet/bubbletea/) as its UI framework.

## Configuration
The CLI will read the file at `~/.config/list-my-project/.project.json`, don't forget to copy your config file if you edited the one provided in this project.

## Usage
You can either use it with `go run main.go` or by exporting it as a binary in your PATH using `go build -o <binary_path/binary_name>`.

## Screenshots

| **List** | **New** | **Edit** |
| - | - | - |
| <img height="250px" src="https://user-images.githubusercontent.com/16008095/185778404-088da6ad-4a7f-4575-a09d-39c3a3a9921d.png"> | <img height="250px" src="https://user-images.githubusercontent.com/16008095/185778416-0609afc3-1a96-437d-a408-361b9f4408e6.png"> | <img height="250px" src="https://user-images.githubusercontent.com/16008095/185778433-d1bb9553-9d14-44c6-9430-887ab4fd47ac.png"> |

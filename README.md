# list-my-projects
![ci badge](https://github.com/marcantoineg/list-my-projects/actions/workflows/ci.yml/badge.svg)

A simple shell app to list projects and open them in a new window of VS Code.

This app uses [Bubble Tea](https://github.com/charmbracelet/bubbletea/) as its UI framework.

## Configuration
The CLI will read the file at `~/.config/list-my-project/.project.json`, don't forget to copy your config file if you edited the one provided in this project.

## Usage
You can either use it with `go run main.go` or by exporting it as a binary in your PATH using `go build -o <binary_path/binary_name>`.

## Screenshots

| **List** | **Form** |
| - | - |
| <img height="250px" src="https://user-images.githubusercontent.com/16008095/186266791-2ed13ab6-9b87-4004-998c-eeffa9d3fa13.png"> | <img height="250px" src="https://user-images.githubusercontent.com/16008095/186266988-c0722604-84aa-47e0-a6c9-c4101b03d91f.png">
| <img height="250px" src="https://user-images.githubusercontent.com/16008095/186266868-467a9f86-dab0-474c-aed8-28ace1d6c3b7.png"> | <img height="250px" src="https://user-images.githubusercontent.com/16008095/186267110-3b90a244-0322-4904-9a09-e641449b823a.png"> |

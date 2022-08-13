# list-my-projects

A simple cli app to list projects and open them in a new window of VS Code.

This app uses [Bubble Tea](https://github.com/charmbracelet/bubbletea/) as its UI framework.

## Configuration
The CLI will read the file at `~/.config/list-my-project/.project.json`, don't forget to copy your config file if you edited the one provided in this project.

## Usage
You can either use it with `go run main.go` or by exporting it as a binary in your PATH using `go build -o <binary_path/binary_name>`.

## Screenshots

<img height="300" src="https://user-images.githubusercontent.com/16008095/184463459-f6d2eaeb-6bb3-4f6c-b90f-b4baacf5a555.png">
<img height="300" src="https://user-images.githubusercontent.com/16008095/184463467-481ac1aa-8205-44ad-b1f8-3dcd2051b893.png">

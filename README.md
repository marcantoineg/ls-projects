# list-my-projects

A simple cli app to list projects and open them in a new window of VS Code.

## Configuration
You will need to create a new directory at `~/.config/go-apps/` and copy the file `.projects.json` provided in this repository.
To add your first project, just edit the json provided and launch the CLI.

_Note: the CLI will read the file at `~/.config/go-apps/.project.json`, don't forget to copy your config file if using the one provided in this repo._

## Usage
You can either use it with `go run main.go` or by exporting it as a binary in your PATH using `go build -o <binary_path/binary_name>`.

## Screenshots

<img height="300" src="https://user-images.githubusercontent.com/16008095/184043384-10b4fd31-b6f6-4644-9a72-dee966abc1cf.png">

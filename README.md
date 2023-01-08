<img src="https://user-images.githubusercontent.com/16008095/208336763-22bec39c-6a44-4469-96bc-675b0f2e85de.png" align="right" width="37%" />

# 📚 list-my-projects
![ci badge](https://github.com/marcantoineg/list-my-projects/actions/workflows/ci.yml/badge.svg)
<br><em><sub>without a meme, is it really a repo?</sub></em>

A simple Go app to list projects and open them in a new window of VS Code.

This app uses [Bubble Tea](https://github.com/charmbracelet/bubbletea/) as its UI framework.

<br/>

## Configuration

The CLI will read the file at `~/.config/list-my-project/.project.json`, don't forget to copy your config file if you edited the one provided in this project.

## Usage
You can either use it with `go run main.go` or by exporting it as a binary in your PATH using `go build -o <binary_path/binary_name>`.

## Screenshots 📸

### List

<img width="80%" src="https://user-images.githubusercontent.com/16008095/186266791-2ed13ab6-9b87-4004-998c-eeffa9d3fa13.png">
<img width="80%" src="https://user-images.githubusercontent.com/16008095/186266868-467a9f86-dab0-474c-aed8-28ace1d6c3b7.png">

### Form
<img width="50%" src="https://user-images.githubusercontent.com/16008095/186266988-c0722604-84aa-47e0-a6c9-c4101b03d91f.png">
<img width="50%" src="https://user-images.githubusercontent.com/16008095/186267110-3b90a244-0322-4904-9a09-e641449b823a.png">

### Search
<img width="80%" src="https://user-images.githubusercontent.com/16008095/211213373-aa7d409b-31fa-46ed-b2f1-0f82e8f24b55.png">

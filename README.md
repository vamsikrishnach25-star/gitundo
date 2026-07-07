# GitUndo

GitUndo is a lightweight CLI tool built in Go that simplifies common Git undo operations. Instead of remembering multiple git flags and commands, GitUndo wraps them into simple, memorable commands.

## 🚀 Features

- **Undo last commit** — revert your most recent commit while keeping changes staged
- **Undo file changes** — discard changes made to a specific file
- **Unstage files** — remove a file from the staging area without losing changes
- **Revert by commit hash** — revert any specific commit using its hash
- **View commit log** — quickly view recent commit history

## 🛠️ Installation & Usage

Clone the repository and run it with Go:

```bash
git clone https://github.com/vamsikrishnach25-star/gitundo.git
cd gitundo
go run main.go --help
```

### Commands

| Command | Description |
|---|---|
| `gitundo last-commit` | Undo the last commit, keeping changes staged |
| `gitundo file <filename>` | Discard changes made to a specific file |
| `gitundo unstage <filename>` | Remove a file from the staging area |
| `gitundo revert <commit-hash>` | Revert a specific commit by its hash |
| `gitundo log` | View recent commit history |

## 🧠 How it works

GitUndo is built using Go's `os/exec` package to wrap native Git commands under a simpler, more intuitive CLI interface. Each command maps to one or more underlying `git` operations, handling argument parsing and execution internally so users don't need to remember exact Git syntax.

## 📦 Tech Stack

- **Language:** Go
- **Core dependency:** Git (must be installed and available in PATH)

## 🤝 Contributing

Contributions are welcome! Feel free to fork the repo, open issues, or submit pull requests.

## 📄 License

MIT (or update this to match your actual license, if any)

# 📌 Task

**task** is a command-line tool written in Go for efficient management of your TODOs.

![](task.gif)

## Installation

To install the `task` command-line tool, ensure you have Go installed on your system. If not, you can download and install it from the [official Go website](https://golang.org/dl/).

Once you have Go installed, open a terminal or command prompt and run the following command:

```bash
go install github.com/oussamaM1/task/
```

This command will download the repository, build the `task` executable, and place it in your Go binary directory. Make sure your Go binary directory is in your system's PATH so that you can execute `task` from any directory.

## Usage

```bash
task [command]
```

## Available Commands

- `add`: Add a new task to your TODO list
- `completion`: Generate the autocompletion script for the specified shell
- `do`: Mark a task as complete, in-progress, or open
- `help`: Get help about any command
- `list`: List all of your incomplete tasks

## Flags

- `-h, --help`: Show help for the `task` command

Use `task [command] --help` for more detailed information about a specific command.

## Example

Here's an example of how to use `task`:

```bash
# Add a new task
task add "Write README.md"
✅ The task has been successfully added.

# List all tasks
task list
╭────────┬─────────────────────────────┬───────────────────────╮
│ ID     │         DESCRIPTION         │         STATE         │
├────────┼─────────────────────────────┼───────────────────────┤
│      1 │ Write README.md             │ 📌 Open               |
│        │                             │                       |
╰────────┴─────────────────────────────┴───────────────────────╯

# Mark a task in-progress
# format: task do <ID> <Open, In-progress, Completed>
task do 1 "In-progress"
✅ The task has been successfully updated.

# List all tasks
task list
╭────────┬─────────────────────────────┬───────────────────────╮
│ ID     │         DESCRIPTION         │         STATE         │
├────────┼─────────────────────────────┼───────────────────────┤
│      1 │ Write README.md             │ 🚀 In-progress        │
│        │                             │                       │
╰────────┴─────────────────────────────┴───────────────────────╯

# Mark a task completed
# format: task do <ID> <Open, In-progress, Completed>
task do 1 "Completed"
✅ The task has been successfully updated.

# List all tasks
task list
╭────────┬─────────────────────────────┬───────────────────────╮
│ ID     │         DESCRIPTION         │         STATE         │
├────────┼─────────────────────────────┼───────────────────────┤
│      1 │ Write README.md             │ ✅ Completed          │
│        │                             │                       │
╰────────┴─────────────────────────────┴───────────────────────╯
```

# git-commit-flood

Do you want to fill your activity history?

Fill your git repository with LOADS of commits!

## Build

```bash
go build
```

... with go 1.16+

... or just use `go run`!

## Usage

Execute the binary, it will generate the file needed to fill your git
repository, and a script named `commit.sh` that will automatically commit
the changes for you.

Please move the generated `output` directory elsewhere before executing
`commit.sh`.

## Note

**This simple program violated the *Open Close Principle*. Planned to have it 
refactored in near future.** (2022-06-02)

Still WIP, may not work as expected (but I will try my best to make it work).

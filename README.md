# b

> A tool that picks branch names for your pull requests.

## Installation

### Go

```shell
go install github.com/roj1512/b@latest
```

### Binary

1. Download the latest release from
   [here](https://github.com/roj1512/b/releases).
2. Rename it to `b`.
3. Optionally strip it.
4. Add it to path.

## Example Usage

```shell
git checkout -b $(b fix: make the ThingyThing handle the X.Y.Z correctly)
```

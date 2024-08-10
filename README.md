# go-dicewarelib

Contains a few basic functions that allow for diceware password generation, with sensible defaults that work with most websites.

**This is `v0.0.X` software whose API may change at any time, and is not suitable for any serious production use.** It's really just meant for my own reuse in other similar applications.

## Setup

If you're going to use this in your application, you should download `words-complex.txt` and/or `words-simple.txt` depending on your needs. These are managed via `git lfs`.

Obtain it via `go get`:

```bash
go get git.cmcode.dev/cmcode/go-dicewarelib
```

## Example

See `./example/main.go` for this full example.

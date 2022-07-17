# Gordle

Go implementation of Wordle game

## Building instructions

```bash
git clone https://git.dmitriy.icu/dm1sh/gordle
cd gordle
go mod tidy
go build .
```

## Running instructions

```bash
% ./gordle --help

Usage of ./gordle:
  -f string
        Path to dictionary file with words of length -n. If specified, -n must be specified too (if the latter must not be equal to default value (default "./dictionary/5.txt")
  -n int
        Word size (default 5)
```

## Testing instructions

```bash
go test # [-cover]
```

## Dictionaries

To run game you also need dictionary files. By default gordle searches for them in `./dictionary/<n>.txt` file, where `<n>` - number of characters, specified by `-n` flag. If want to use custom one, provide in `-f` flag path to txt file with strings separated by new line characters

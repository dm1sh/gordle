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
  -n int
        Word size (default 5)
```

## Testing instructions

```bash
go test # [-cover]
```

## Dictionaries

To run game you also need dictionary files. By default gordle searches for them in `./dictionary/<n>.txt` file, where `<n>` - number of characters, specified by `-n` flag

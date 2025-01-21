# my hack for kitty terminal emulator

## build
```sh
./dev.sh deps         # install dependencies
./build.sh clean
./build.sh kitty.app
fd -u -t d kitty.app
```

## docs (incomplete)
```sh
./dev.sh deps -for-docs && ./dev.sh docs
```

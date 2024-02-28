# Prof of concept

Making it possible to have the same codebase, but injecting different behaviors depending on build tags.

## How to build
> make sure you have google wire cmd installed

```
make wire
```

### To build the free version:
```
go build -tags free
```

### To build the pro version:
```
go build -tags free
```

## How to run
```
./ioc
```

## Features

### Free
* endpoint running on port `8000` that says `Hello, Jordan`

## Pro
* endpoint running on port `8000` that says `Hello, Premium Jordan`
* new endpoint `https://localhost:8000/hey`

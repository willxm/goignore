Generate .gitignore in command line
===
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)][license]

[license]: https://github.com/willxm/goignore/blob/master/LICENSE

## Features
- Golang
- C

## TODO
- Auto sync .gitignore from [gitignore](https://github.com/github/gitignore)
- Auto analysis of current directory project types


## Install
```
$ go get github.com/willxm/goignore
$ cd GOPATH/src/github.com/willxm/goignore
$ go install
```

## Usage
```
$ goignore go
.gitignore is existed, do you want override or update it ?
update
```

## Acknowledgements

* [gitignore](https://github.com/github/gitignore)
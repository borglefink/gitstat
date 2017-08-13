## Description

*gitstat* is a tiny command line utility that finds all git repositories below the given path, and for each runs "git status -s".

## Usage

Give a directory as a parameter. If none is given, the current directory is used.

```
gitstat [directory]
```

## Install

Clone the repository into your GOPATH somewhere and resolve dependencies (see below),
and do a **go install**.

## Dependencies

_cntsrc_ is dependent upon Michael T Jones' fast parallel filesystem traversal package. 
See [github.com/MichaelTJones/walk](https://github.com/MichaelTJones/walk). 

It also uses Brian Downs' spinner package, 
[github.com/briandowns/spinner](https://github.com/briandowns/spinner), 
for showing progress while walking the given directory. 

Resolve by doing:
```
go get github.com/MichaelTJones/walk
go get github.com/briandowns/spinner
```

## License

A MIT license is used here - do what you want with this. 
Nice though if improvements and corrections could trickle back to me somehow. :-)

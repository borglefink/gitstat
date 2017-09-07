## Description

*gitstat* is a tiny command line utility that finds all git repositories below the given path, and for each repository runs *git status -s*, the briefer version of *git status*.

## Usage

Give a directory as a parameter. If none is given, the current directory is used.

```
gitstat [options] [dirname]
```
Option *-a* shows directory name for every git repository, even if there is no status.
Normally 'git status -s' is shown for each repository, the option *-l* changes it to show the normal (long) 'git status'. 
To list ignored files, use the *-i* option.
Use *gitstat -?* for help.

## Install

Clone the repository into your GOPATH somewhere and resolve dependencies (see below),
and do a **go install**.

## Dependencies

*gitstat* is dependent upon Michael T Jones' fast parallel filesystem traversal package. 
See [github.com/MichaelTJones/walk](https://github.com/MichaelTJones/walk). 

It also uses Brian Downs' spinner package, 
[github.com/briandowns/spinner](https://github.com/briandowns/spinner), 
for showing progress while walking the given directory. 

Resolve by doing:
```
go get github.com/MichaelTJones/walk
go get github.com/briandowns/spinner
```

## Background

When working with a lot of different git repositories simultaneously, 
it's nice to hav a tool that helps to get a larger view of what is going on.
As usual, this was a tool I needed, and it was fun to make.

## License

A MIT license is used here - do what you want with this. 
Nice though if improvements and corrections could trickle back to me somehow. :-)

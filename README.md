go-bitbucket-backup
===================

A Go command-line tool to backup a collection of git repos to BitBucket

## Install

```
go get github.com/jpillora/go-bitbucket-backup
```

## Usage

```
$ go-bitbucket-backup help

NAME:
   bitbucket-backup - BitBucket Backup

USAGE:
   bitbucket-backup [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --directory    Directory containing Git repositories (defaults to cwd)
   --username     BitBucket username
   --password     BitBucket password
   --namespace    BitBucket namespace (used to create repositories, defaults to username)
   --reset    Resets the repositories on BitBucket before pushing (deletes then re-creates)
   --version, -v  print the version
   --help, -h   show help
```

## Todo

* Parallelize

#### MIT License

Copyright &copy; 2014 Jaime Pillora &lt;dev@jpillora.com&gt;

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
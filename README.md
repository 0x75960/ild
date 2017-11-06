ild
====

inquiry each lines to line-based database

usage
-----

### installation

```sh
go get -u github.com/0x75960/ild
```

### how to use

```console
ILD
        inquiry each lines to line-based database

Usage of this:
        ild -i [file to inquiry] -db [database file]

  -db string
        database file path
  -i string
        input file path
  -v    pick up not founded on database
```

#### what to do

This tool do roughly same as trailing command

```sh
cat <(cat /path/to/database_file | sort | uniq) <(cat /path/to/inquiry | sort | uniq) | sort | uniq -d # or uniq -u when specified -v option
```

I make this because I want do this in Windows.

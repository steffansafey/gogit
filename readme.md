# Gogit (git, but in go!)

Very much a work-in-progress. Capable of hashing & writing objects to the object store via `commit`, haven't added proper branching yet...)


Compile using
```
go build
```

View available commands by running `./gogit`

```
gogit is a simple git clone

Usage:
  gogit [command]

Available Commands:
  cat-file    Cats a file if it exists in the objects directory
  commit      Create a new commit object
  completion  Generate the autocompletion script for the specified shell
  hash-object Compute object ID and creates a blob from a file
  help        Help about any command
  init        Initialize a new repository
  read-tree   Reads the tree object with the given oid
  write-tree  Writes the current state of the index to the objects directory
```

Try it out by making a new directory and running the below commands. Objects will be written to the `.gogit` folder.

``` bash
mkdir test-repo && cd test-repo

../gogit init

echo "this is a test file" > test.txt

../gogit commit -m "initialize the repo"
```
# Trie Search

```
In computer science, a trie, also called digital tree or prefix tree, is a kind of search tree—an ordered tree data structure used to store a dynamic set or associative array where the keys are usually strings. Unlike a binary search tree, no node in the tree stores the key associated with that node; instead, its position in the tree defines the key with which it is associated; i.e., the value of the key is distributed across the structure. All the descendants of a node have a common prefix of the string associated with that node, and the root is associated with the empty string. Keys tend to be associated with leaves, though some inner nodes may correspond to keys of interest. Hence, keys are not necessarily associated with every node. For the space-optimized presentation of prefix tree, see compact prefix tree.
```

More about this data structure: https://en.wikipedia.org/wiki/Trie

## Running

Arguments:

- --search -> search query
- --dir -> directory to scan

Check Makefile

## Examples

```
~/Projects/lab/go-tries > make run
go run src/*.go
[docs/danube.txt docs/germany.txt docs/netherlands.txt docs/romania.txt]
[~] Elapsed: 686.557µs
[!] Ready. Loaded 4 tries
[!] Performing search for term: romania
[!] Found results in 1 file(s)
~ docs/romania.txt
[~] Elapsed: 25.875µs
```

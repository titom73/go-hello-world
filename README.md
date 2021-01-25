# Simpple GO Hello-World

First script using GO to get data from Github API

Inputs:

- Repository name
- Milestone

And program will print list of all closed PR related to given Milestone

```bash
$ go run read-gh-json.go
-> Enter Repository name (org/name): aristanetworks/ansible-avd
-> Enter milestone: v1.1.2

PR name: Release 1.1.2
PR state: closed
PR ID: 471
Author: carlbuchmann
...
```

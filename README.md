## tw_set_deps

Simple [Go](https://golang.org/) program to manage dependencies in [taskwarrior.](https://taskwarrior.org/)

Let's say you have the following tasks:

| 1 | Do homework |
| 2 | Write report |
| 3 | Submit report |

To enforce the condition that task 3 depends on task 2, which inturn depends on task 1, you can write:

```
task 3 modify depends task 2
task 2 modify depends task 1
```

This can get annoying when there's too many tasks!

Using `tw_set_deps`, you can set task dependencies like so:

```
./tw_set_deps 1 2 3
```

Changelog
=========

1. rename package `ebony` -> `rbtree`
2. change location from `github.com/logrusorgru/ebony` -> `github.com/logrusorgru/rbtree`
3. change license from WTFPL to Unlicense
4. use generics
5. use go modules
6. change count type `uint` -> `int`
7. `Exists` -> `IsExists`, following Go naming recommendations
8. don't call `runtime.GC()` in `Tree.Empty()`
9. change `Walker` -> `WalkFunc`, following Go naming recommendations
10. change `Range` -> `Slice`
11. improve documentation
12. rename `Count` -> `Len`

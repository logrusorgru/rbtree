Ebony
=====

[![GoDoc](https://godoc.org/github.com/logrusorgru/ebony?status.svg)](https://godoc.org/github.com/logrusorgru/ebony)
[![WTFPL License](https://img.shields.io/badge/license-wtfpl-blue.svg)](http://www.wtfpl.net/about/)

Golang red-black tree with uint index, not thread safe

### Methods

| Method name | Time |
|:-----------:|:----:|
| Set   | O(log*n*) |
| Del   | O(log*n*) |
| Get   | O(log*n*) |
| Exist | O(log*n*) |
| Count | O(1) |
| Move  | O(2log*n*) |
| Range | O(log*n* + *m*) |
| Max   | O(log*n*) |
| Min   | O(log*n*) |
| Flush | O(1) |

### Memory usage

O(*n*&times;node),

node = 3&times;ptr_size +
       uint_size +
       bool_size +
       data_size

### Licensing

Copyright &copy; 2015 Konstantin Ivanov <ivanov.konstantin@logrus.org.ru>
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE.md file for more details.

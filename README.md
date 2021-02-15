# jalopy

Command line utility that can be used to generate a catchy chord progression for a song.

```
% go test ./... && go install && jalopy song -key G
?       jalopymusic.com/jalopy  [no test files]
ok      jalopymusic.com/jalopy/numerals (cached)
ok      jalopymusic.com/jalopy/tonality (cached)
[]int{1, 4, 5, 1}
[]int{1, 4, 5, 1}
[]int{1, 4, 5, 1}
[]int{1, 4, 5, 1}
----
[]int{1, 1, 1, 6}
[]int{1, 2, 5, 5}
[]int{1, 2, 5, 5}
[]int{1, 2, 5, 5}
----
G   C   D   G
G   C   D   G
G   C   D   G
G   C   D   G
----
G   G   G   Em
G   Am  D   D
G   Am  D   D
G   Am  D   D
```

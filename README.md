# Reggenerator
Generate random string based on supplied regex

## Quick start

### Lib

`go get github.com/toncek345/reggenerator`  
Import and call `Generate(regex string, count int)`.  

If the random generation function is not being changed, `rand.Int` will be used.  
Seeding with `rand.Seed(time.Now().UnixNano())` is then recommended.

#### How to change the random number generation

```go
reggenerator.RandFn = randomFn
---
reggenerator.RandFn = func() int { return X }
```

### Standalone

`go install github.com/toncek345/reggenerator/reggenerator@v1.1.0`
```
> reggenerator --help
Usage of cmd:
  -count int
    	number of random string (default 1)
  -regex string
    	regex for random string

```


## Syntax

`.` Match any character except newline  
`[` Start character class definition  
`]` End character class definition  
`?` 0 or 1 quantifier  
`{` Start min/max quantifier  
`}` End min/max quantifier  

Within the class:  
`^` Negate the class, but only if the first character  
`-` Indicates character range

## Examples

```
> /[-+]?[0-9]{1,16}[.][0-9]{1,6}/
-1752643936.096896
9519688.31
+1.7036
+65048.3876
-6547028036936294.111
07252345.650
-27557.78
7385289878518.439775
13981103761187.90
4100273498885.614

> /[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{8}/
eb4bbfa4-11d4-bce9-0f56e2dc
d896fd6b-9bfe-d0ae-6fb05b52
b36dcaeb-5654-73aa-c9ec7de2

> /.{8,12}/
(<W[+]%i
7QEyw0th
rEF\Ly(C

> /[^aeiouAEIOU0-9]{5}/
#DTdH
B[n<F
rsQgV

> /[a-f-]{5}/
-cfbc
ab---
----f
```

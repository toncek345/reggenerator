# Reggenerator
Generate random string based on supplied regex

## Quick start

Clone and run as binary as `go run cmd/main -regex="/something/" -count=2` or include as a library  
and supply arguments to `Generate(regex string, count int)` function.

If using as a library you can supply your own random function  
`reggenerator.RandFn = randomFn`. Otherwise `math.Int()` is being used.

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

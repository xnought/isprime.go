# isprime.go

```bash
go build
```

For example I want to see if the numbers in `nums` are prime

```txt
2
3
4
17
5
6
79999999
81
73939133
0
```

I can pipe the nums into isprime like


```bash
cat nums | ./isprime
```

outputs

```txt
1
1
0
1
1
0
0
0
1
0
```

where 1 means the number is prime and 0 means is not.
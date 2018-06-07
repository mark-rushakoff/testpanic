# testpanic

Repository to reproduce a panic when running go test.

With this version of go (or earlier; I haven't bisected to find when this was introduced):
```
$ gotip version
go version devel +7b08e619bb Thu Jun 7 16:17:04 2018 +0000 darwin/amd64
```

Clone the repository and run `go test .`.

```
$ GOTRACEBACK=crash gotip test .
panic: Error parsing builtin CA x509: RSA key missing NULL parameters

goroutine 1 [running]:
panic(0x1271080, 0xc000095050)
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/panic.go:537 +0x2cb fp=0xc000149df0 sp=0xc000149d60 pc=0x102cbfb
github.com/mark-rushakoff/testpanic/vendor/github.com/elazarl/goproxy.init.0()
        /Users/mr/go/src/github.com/mark-rushakoff/testpanic/vendor/github.com/elazarl/goproxy/certs.go:10 +0x1bc fp=0xc000149e58 sp=0xc000149df0 pc=0x1240b0c
github.com/mark-rushakoff/testpanic/vendor/github.com/elazarl/goproxy.init()
        <autogenerated>:1 +0x250 fp=0xc000149f78 sp=0xc000149e58 pc=0x1240d70
github.com/mark-rushakoff/testpanic.init()
        <autogenerated>:1 +0x4a fp=0xc000149f88 sp=0xc000149f78 pc=0x1243faa
main.init()
        <autogenerated>:1 +0x54 fp=0xc000149f98 sp=0xc000149f88 pc=0x1244254
runtime.main()
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/proc.go:189 +0x1bd fp=0xc000149fe0 sp=0xc000149f98 pc=0x102e8ad
runtime.goexit()
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/asm_amd64.s:1333 +0x1 fp=0xc000149fe8 sp=0xc000149fe0 pc=0x105b941

goroutine 2 [force gc (idle)]:
runtime.gopark(0x12eb408, 0x14ea090, 0x1410, 0x1)
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/proc.go:298 +0xeb fp=0xc000032f80 sp=0xc000032f60 pc=0x102eceb
runtime.goparkunlock(0x14ea090, 0x1410, 0x1)
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/proc.go:304 +0x53 fp=0xc000032fb0 sp=0xc000032f80 pc=0x102ed93
runtime.forcegchelper()
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/proc.go:251 +0xb3 fp=0xc000032fe0 sp=0xc000032fb0 pc=0x102eb63
runtime.goexit()
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/asm_amd64.s:1333 +0x1 fp=0xc000032fe8 sp=0xc000032fe0 pc=0x105b941
created by runtime.init.4
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/proc.go:240 +0x35

goroutine 18 [GC sweep wait]:
runtime.gopark(0x12eb408, 0x14ea1e0, 0x102140c, 0x1)
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/proc.go:298 +0xeb fp=0xc00002e780 sp=0xc00002e760 pc=0x102eceb
runtime.goparkunlock(0x14ea1e0, 0x131140c, 0x1)
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/proc.go:304 +0x53 fp=0xc00002e7b0 sp=0xc00002e780 pc=0x102ed93
runtime.bgsweep(0xc000080000)
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/mgcsweep.go:52 +0x8f fp=0xc00002e7d8 sp=0xc00002e7b0 pc=0x1020b5f
runtime.goexit()
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/asm_amd64.s:1333 +0x1 fp=0xc00002e7e0 sp=0xc00002e7d8 pc=0x105b941
created by runtime.gcenable
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/mgc.go:216 +0x58

goroutine 34 [finalizer wait]:
runtime.gopark(0x12eb408, 0x1507da0, 0x129140f, 0x1)
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/proc.go:298 +0xeb fp=0xc000032728 sp=0xc000032708 pc=0x102eceb
runtime.goparkunlock(0x1507da0, 0x140f, 0x1)
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/proc.go:304 +0x53 fp=0xc000032758 sp=0xc000032728 pc=0x102ed93
runtime.runfinq()
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/mfinal.go:175 +0x99 fp=0xc0000327e0 sp=0xc000032758 pc=0x10185b9
runtime.goexit()
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/asm_amd64.s:1333 +0x1 fp=0xc0000327e8 sp=0xc0000327e0 pc=0x105b941
created by runtime.createfing
        /Users/mr/gotip/src/github.com/golang/go/src/runtime/mfinal.go:156 +0x61
FAIL    github.com/mark-rushakoff/testpanic     0.019s
```

In contrast, go 1.10 seems to work properly:

```
$ go version
go version go1.10.1 darwin/amd64

$ go test .
?       github.com/mark-rushakoff/testpanic     [no test files]
```

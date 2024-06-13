# TryCatch-go
Try/Catch/Else/Finally implementation for Go

Main credits go to Peter Verhas, more specifically this page: https://dzone.com/articles/try-and-catch-in-golang.

Changes made:
- Have a boolean that tells if the code panicked or not, so that Tcef can know if the Catch function  should execute or
not. This is because panic(nil) is a valid panic (nil interface for example), but when a function returns normal, the
recover() function returns nil. So recover() returns nil whether there is a panic or not. The idea takes care of that
(checks if the function panicked or not). EDIT: it seems that with Go 1.21+, `panic(nil)` no longer panics with `nil` as
pointed out by VonC in my answer [here](https://stackoverflow.com/a/76913649/8228163). But as I'm still using Go 1.20,
I'm keeping the check here.
- Also added Python's Else clause, which executes if the Try clause didn't panic.
- Now the whole Tcef doesn't need the Catch clause != nil to catch the panic. It will always catch the panic and execute
the Catch clause if there is one. If there is no Catch clause, it will just recover the panic and do nothing.

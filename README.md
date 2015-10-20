# Lights Off
The last one to leave turns the lights off.

## Why?

`LightsOff` is built for closing channels when you have a "fan in" architecture. If you have several go routines writing to a channel, then when do you close that channel? Ideally, you don't close it on a go routine that is reading because the writing go routines will panic.

The built-in methods `recover()` and `sync.WaitGroup` both have issues.

`recover()` implies that you close early and the writing go routines get ended early. You don't let the panic end your process because your `recover` it, but you also don't let them finish. This tends to be a hack and doesn't let the data complete its flow.
 
Using `sync.WaitGroup` means you have to burn a go routine. The idea is that you would do the following:

```
var wg sync.WaitGroup
go func(){
  wg.Wait()
  close(someChannel)
}{}
```

While this works fine, it does allocate a go routine. If you use the pattern quite a bit, you end up leaking go routines. This adds up in the go runtime.

Using `LightsOff` costs quite a bit less.

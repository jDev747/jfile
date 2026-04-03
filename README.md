This is a way to turn a folder into a json file.

It does not compress.

Check out the [wiki](https://github.com/jDev747/jfile/wiki/)

How to import:

```go
import "github.com/jDev747/jfile"
```
From then on you can use it like this:

```go

func main() {
    var item jfile.JItem = jfile.DirToJDir("test")
}
```

If you want to know what which part of the module does what, check the [wiki](https://github.com/jDev747/jfile/wiki/)

Note: It is unfortunately not done yet.

Any Issues or Pull requests welcome!

-- jDev747

# nanoid
Golang port of [ai/nanoid](https://github.com/ai/nanoid) (originally written in JavaScript).

# Description
A tiny, secure URL-friendly unique string ID generator for Golang.

**Safe.** It uses cryptographically strong random APIs
and guarantees a proper distribution of symbols.

**Compact.** It uses more symbols than UUID (`A-Za-z0-9_~`)
and has the same number of unique options in just 22 symbols instead of 36.

**No third party dependencies** No need to pollute your $GOPATH

## Install

```bash
$ go get github.com/jkomyno/nanoid
```

## Testing

``` bash
$ go test -v -bench=.
```

You should be able to see a log similar to the following one:
```
=== RUN   TestGeneratesURLFriendlyIDs
--- PASS: TestGeneratesURLFriendlyIDs (0.00s)
=== RUN   TestHasNoCollisions
--- PASS: TestHasNoCollisions (0.21s)
=== RUN   TestFlatDistribution
--- PASS: TestFlatDistribution (0.33s)
goos: linux
goarch: amd64
pkg: github.com/jkomyno/nanoid
BenchmarkNanoid-4        1000000              1704 ns/op
PASS
ok      github.com/jkomyno/nanoid       2.265s
```

## Usage
**This packages tries to offer an API as close as possible to the original JS module.**

### Normal

The Nanoid() function uses URL-friendly symbols (`A-Za-z0-9_~`) and returns an ID
with 22 characters (to have the same collisions probability as UUID v4).
Please note that it also returns an error, which (hopefully) will be `nil`.

```go
import "github.com/jkomyno/nanoid"

id, err := nanoid.Nanoid() //=> "Uakgb_J5m9g~0JDMbcJqLJ"
```

Symbols `-,.()` are not encoded in URL, but in the end of a link
they could be identified as a punctuation symbol.

If you want to reduce ID length (and increase collisions probability),
you can pass length as argument:

```go
import "github.com/jkomyno/nanoid"

id, err := nanoid.Nanoid(10) //=> "IRFa~VaY2b"
```

### Custom Alphabet or Length

If you want to change the ID alphabet or the length
you can use low-level `Generate` function.

```go
import "github.com/jkomyno/nanoid"

id, err := nanoid.Generate("1234567890abcdef", 10) //=> "4f90d13a42"
```

Alphabet must contain less than 256 symbols.

### Custom Random Bytes Generator

You can replace the default safe random generator using the `Format` function.

```go
import (
    "crypto/rand"

    "github.com/jkomyno/nanoid"
)

func random(size int) ([]byte, error) {
	var randomBytes = make([]byte, size)
	_, err := rand.Read(randomBytes)
	return randomBytes, err
}

id, err := nanoid.Format(random, "abcdef", 10) //=> "fbaefaadeb"
```

Note that `random` function must follow this spec:
```go
type RandomType func(int) ([]byte, error)
```

If you want to use the same URL-friendly symbols with `format`,
or take a look at the other defaults value, you can use `GetDefaults`.

```go
import "github.com/jkomyno/nanoid"

var defaults *nanoid.DefaultsType
defaults = nanoid.GetDefaults()
/*
	&DefaultsType{
		Alphabet: "_~0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		Size:     22,
		MaskSize: 5,
	}
*/
```

## Credits

[ai](https://github.com/ai) - [nanoid](https://github.com/ai/nanoid)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.

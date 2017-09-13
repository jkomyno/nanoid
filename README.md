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

## Usage

### Normal

The main module uses URL-friendly symbols (`A-Za-z0-9_~`) and returns an ID
with 22 characters (to have the same collisions probability as UUID v4).
Please note that it also returns an error, which (hopefully) will be `nil`.

```go
import "nanoid"

id, err := nanoid.Nanoid() //=> "Uakgb_J5m9g~0JDMbcJqLJ"
```

Symbols `-,.()` are not encoded in URL, but in the end of a link
they could be identified as a punctuation symbol.

If you want to reduce ID length (and increase collisions probability),
you can pass length as argument:

```go
id, err := nanoid.Nanoid(10) //=> "IRFa~VaY2b"
```

### Custom Alphabet or Length

If you want to change the ID alphabet or the length
you can use low-level `Generate` module.

```go
id, err := nanoid.Generate("1234567890abcdef", 10) //=> "4f90d13a42"
```

Alphabet must contain less than 256 symbols.

### Custom Random Bytes Generator

You can replace the default safe random generator using the `Format` module.

```go
import (
    "nanoid"
    "crypto/rand"
)

func random(size int) ([]byte, error) {
	var randomBytes = make([]byte, size)
	_, err := rand.Read(randomBytes)
	return randomBytes, err
}

id, err := Format(random, "abcdef", 10) //=> "fbaefaadeb"
```

Note that `random` function must follow this spec:
```go
type RandomType func(int) ([]byte, error)
```

If you want to use the same URL-friendly symbols with `format`,
or take a look at the other defaults value, you can use `GetDefaults`.

```go
var defaults *nanoid.DefaultsType
defaults = nanoid.getDefaults()
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

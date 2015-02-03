LoGo
====

LoGo is a Golang logging library.

License
-------

LoGo is licensed under the [Simplified BSD License](http://choosealicense.com/licenses/bsd-2-clause/).

How to Install
--------------

Make Sure You Have golang installed!


    $ go get github.com/Nurza/LoGo

How to Upgrade
--------------

Make Sure You Have golang installed!


    $ go get -u github.com/Nurza/LoGo

How to Use
--------------

See the "examples" directory.

Examples
--------------

- Simple console logging

```go
type SimpleStructure struct {
	A string
	B bool
	C int
	D float64
	F byte
}

func main() {
	logo.Log(true)                                                 // bool
	logo.Log(42)                                                   // int
	logo.Log(42.42)                                                // float64
	logo.Log('A')                                                  // byte
	logo.Log(SimpleStructure{"Hello World", true, 42, 42.42, 'A'}) // Structure
	logo.Log("Hello", true, 42, "World")                           // Mixed

	logo.Log("Hello", "World", "!") // Is the same as :
	logo.Log("Hello World !")
}
```

![Simple console logging](http://files.nurza.fr/github/logo/simple-console-log.png)

- Console logging

```go
func main() {
	var l logo.Logger            // Create a simple Logger
	l.AddTransport(logo.Console) // Add a transport: Console
	l.EnableAllLevels()          // Enable all logging levels

	l.Log(logo.Info, "Hello World !") // This is the same as:
	l.Info("Hello World !")
}
```

![Console logging](http://files.nurza.fr/github/logo/logger-console.png)

- Console logging with color:

```go
func main() {
	var l logo.Logger                 // Create a simple Logger
	t := l.AddTransport(logo.Console) // Add a transport: Console
	t.AddColor(logo.ConsoleColor)     // Add a color: Console color
	l.EnableAllLevels()               // Enable all logging levels

	l.Silly("Silly")
	l.Debug("Debug")
	l.Verbose("Verbose")
	l.Info("Info")
	l.Warn("Warn")
	l.Error("Error")
	l.Critical("Critical")
}
```

![Console logging with color](http://files.nurza.fr/github/logo/logger-console-color.png)

- Logging with timer:

```go
func main() {
	var l logo.Logger                  // Create a simple Logger
	t := l.AddTransport(logo.Console)  // Add a transport: Console
	t.AddColor(logo.ConsoleColor)      // Add a color: Console color
	l.AddTime("[2006-01-02 15:04:05]") // Add time template
	l.EnableAllLevels()                // Enable all logging levels

	l.Silly("Silly")
	l.Debug("Debug")
	l.Verbose("Verbose")
	l.Info("Info")
	l.Warn("Warn")
	l.Error("Error")
	l.Critical("Critical")
}
```

![Logging with timer](http://files.nurza.fr/github/logo/logger-addtime.png)


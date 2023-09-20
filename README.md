## Installing and using the Package

```go
import (
	woke "onboardbase/woke"
)
func main(){
  cliFunc := func() error {
      // wrap your Cli
      cmd.example()
  }
  woke.Wrap(cliFunc)
}
```

## Development

### 1. Install GO

```
Go to go website, https://go.dev/doc/install download and install go
```

### 2. Install dependencies 

```
go mod tidy
```

Once all things are done

Running the dummy.go file and the package of woke
Is Just Creating a View using tea, bubbletea, glamour, lipgloss

so this is totally the dependency of the projects 

## Questions on dev

### HowTo handle error during init-phase on main.go
```go
var (
	log *zap.SugaredLogger
)

func init() {
	log, _ = logger.New("create-agreement")
}
```

<h1 align="center">
  conceal
</h1>

<h4 align="center">
  Hide sensitive user data
</h4>

<p align="center">
  <a href="http://pkg.go.dev/github.com/gabriel-vasile/conceal">
    <img alt="Docs" src="https://pkg.go.dev/badge/github.com/gabriel-vasile/conceal">
  </a>
  <a href="https://github.com/gabriel-vasile/conceal/actions">
    <img alt="Build" src="https://github.com/gabriel-vasile/conceal/workflows/Go/badge.svg">
  </a>
  <a href="https://goreportcard.com/report/github.com/gabriel-vasile/conceal">
    <img alt="Report" src="https://goreportcard.com/badge/github.com/gabriel-vasile/conceal">
  </a>
  <a href="LICENSE">
    <img alt="License" src="https://img.shields.io/badge/License-MIT-green.svg">
  </a>
</p>

**conceal** takes an `io.Writer` and returns a new `io.Writer` which masks any sensitive user data
written to it.

Currently, **conceal** can hide:
- payment card numbers,
- social security numbers,
- email addreses.

```bash
go get github.com/gabriel-vasile/conceal
```
```go
log := logrus.New()
log.SetOutput(conceal.New(os.Stdout, conceal.CardNumber, conceal.SSN, conceal.Email))
log.Info("My credit card: 5555 5555 5555 4444, my SSN: 234-49-2324, my email: spam.me@email.com.")
//output: My credit card: XXXX XXXX XXXX XXXX, my SSN: SSN-XX-XXXX, my email: sXXXXXe@eXXXl.com.
```

Visit [Go Playground](https://go.dev/play/p/BCwlfPtvKrk) to see the code in action.

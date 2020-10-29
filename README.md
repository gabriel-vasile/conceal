<h1 align="center">
  conceal
</h1>

<h4 align="center">
  Hide sensitive user data
</h4>

<p align="center">
  <a href="LICENSE">
    <img alt="License" src="https://img.shields.io/badge/License-MIT-green.svg">
  </a>
</p>


```go
log := logrus.New()
log.SetOutput(conceal.New(os.Stdout, conceal.CardNumber, conceal.SSN, conceal.Email))
log.Info("My credit card: 5555 5555 5555 4444, my SSN: 234-49-2324, my email: spam.me@email.com.")
// output: My credit card: XXXX XXXX XXXX XXXX, my SSN: SSN-XX-XXXX, my email: sXXXXXe@eXXXl.com.
```

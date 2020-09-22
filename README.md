# Gormrus

This package provides an adapter for Logrus so that it can be used as logger in GORM v2.

## Usage

Simple use case:

```go
package main

import (
    "github.com/sirupsen/logrus"
    "github.com/uniwise/gormrus"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
        Logger: gormrus.New(),
    })
    if err != nil {
        panic("Failed to connect database")
    }

    ...
}
```

Bring you own logrus instance:

```go
package main

import (
    "github.com/sirupsen/logrus"
    "github.com/uniwise/gormrus"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    logger := logrus.New()
    logger.SetLevel(logrus.DebugLevel)

    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
        Logger: gormrus.NewWithLogger(logger),
    })
    if err != nil {
        panic("Failed to connect database")
    }

    ...
}
```

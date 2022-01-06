

```
├── application
│   └── usecase
│       ├── list_order_menu.go
│       └── order_menu.go
├── consts
│   └── order.go
├── domain
│   ├── model
│   │   ├── menu.go
│   │   ├── order.go
│   │   └── user.go
│   └── valueobj
│       ├── rank_score.go
│       └── rts.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── database.go
│   ├── infraservice
│   │   ├── paymentgateway
│   │   │   └── stripe.go
│   │   └── repository
│   │       ├── menu.go
│   │       ├── order_menu.go
│   │       └── user.go
│   └── logger.go
├── prerror
│   └── error.go
└── presentation
    └── handler
        └── order.go
```

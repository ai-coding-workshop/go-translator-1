# Project Structure

```
translator-service/
├── cmd/
│   └── translator/
│       └── main.go          # Application entry point
├── internal/
│   ├── handlers/            # HTTP handlers
│   ├── services/            # Business logic
│   ├── models/              # Data structures
│   └── config/              # Configuration
├── web/
│   ├── static/              # CSS, JS, images
│   └── templates/           # HTML templates
├── configs/                 # Configuration files
├── go.mod                   # Go module definition
└── go.sum                   # Go module checksums
```
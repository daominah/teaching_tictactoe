# Project name
Project summary

### Structure

##### `cmd`
Executables  
`cmd/example`: main app

##### `conf`
Environment variables are used in the app initialization

##### `pkg/core`
Business logic. Can be tested without external elements
(database, network, ..)

##### `pkg/driver`
External elements implement core interfaces.
package db

// Connector allows to use various kind of databases for Twitter application
type Connector interface {
	Connect() error
	Disconnect()
}

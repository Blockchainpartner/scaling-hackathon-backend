package db

// Init initialize connection to the DB and create collection objects if needed
func Init() {
	InitClient()
	InitCollections()
}

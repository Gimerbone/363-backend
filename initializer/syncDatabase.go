package initializer

func SyncDatabase() {
	DB.AutoMigrate()
}

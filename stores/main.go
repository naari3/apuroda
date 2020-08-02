package stores

var (
	// UserStore UserStore
	UserStore *UserDB
	// FileStore FileStore
	FileStore *FileDB
)

// Init Init
func Init() {
	UserStore = NewUserDB()
	FileStore = NewFileDB()
}

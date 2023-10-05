package dto

//Database Database
type Database struct {
	Host      string
	User      string
	Password  string
	DBName    string
	DBNumber  int
	Port      int
	DebugMode bool
}
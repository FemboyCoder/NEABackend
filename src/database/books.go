package database

type Book struct {
	ID           int    `json:"id"`           // id of the book
	Title        string `json:"title"`        // title of the book
	Author       string `json:"author"`       // author of the book
	CurrentOwner string `json:"currentOwner"` // username of current book owner. If empty, book is available for loan
}

func CreateBooksTable() error {
	// create the books table if it doesn't exist
	_, err := DatabaseConnection.Database.Exec(
		`create table if not exists books( 
					id integer primary key autoincrement, 
					title text not null, 
					author text not null,
					currentOwner string,
					foreign key (currentOwner) references users(username)
				)`,
	)
	return err
}

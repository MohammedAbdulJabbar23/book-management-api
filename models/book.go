package models

import "github.com/MohammedAbdulJabbar23/book-management-api/config"

type Book struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Year    int    `json:"year"`
    Cover string `json:"cover"`
	PDFPath string `json:"pdf_path"`
}

func (book *Book) Create() error {
	_, err := config.DB.Exec(`INSERT INTO books (title,author pdf_path) VALUES ($1,$2,$3);`,book.Title, book.Author,book.PDFPath);
    return err;
}

func (book *Book) GetBooks() ([]Book, error) {
    rows,err := config.DB.Query("SELECT * FROM books;");
    if err != nil {
        return nil, err;
    }    
    defer rows.Close();
    var books []Book;
    for rows.Next() {
        var book Book;
        if err := rows.Scan(&book.ID, &book.Title,&book.Author,&book.Year, book.PDFPath); err != nil {
            return nil, err;
        }
        books = append(books, book);
    }
    return books, nil;
}

func (book *Book) UpdatePDFPath(filePath string) error {
    book.PDFPath = filePath;
    _, err := config.DB.Exec("UPDATE books SET pdf_path=$1 WHERE id=$2", book.PDFPath, book.ID);
    return err;
}

func (book *Book) GetPDFPath() error {
    return config.DB.QueryRow("SELECT pdf_path FROM books WHERE id=$1", book.ID).Scan(&book.PDFPath);
}

func (book *Book) GetBook() error {
    query := "SELECT * FROM books WHERE id=$1";
    err := config.DB.QueryRow(query,book.ID).Scan(&book.ID, book.Title,&book.Author, &book.Year, &book.Cover, &book.PDFPath);
    if err != nil {
        return err;
    }
    return nil;
}
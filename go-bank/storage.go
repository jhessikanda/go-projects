package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Init() error {
	return s.createAccountTable()	
}

func (s *PostgresStorage) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY, 
		first_name VARCHAR(50), 
		last_name VARCHAR(50), 
		number INTEGER, 
		balance INTEGER, 
		created_at TIMESTAMP

	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateAccount(acc *Account) error {
	query := `INSERT INTO accounts (
		first_name, 
		last_name, 
		number, 
		balance, 
		created_at
	) VALUES ($1, $2, $3, $4, $5)`
	 
	_, err := s.db.Query(query,
		acc.FirstName, 
		acc.LastName, 
		acc.Number, 
		acc.Balance, 
		acc.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts")
	
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		acc := new(Account)
		err = rows.Scan(&acc.ID, &acc.FirstName, &acc.LastName, &acc.Number, &acc.Balance, &acc.CreatedAt)
		if err != nil {
			return nil, err
		}
		
		accounts = append(accounts, acc)
	}

	return accounts, nil
}

func (s *PostgresStorage) UpdateAccount(a *Account) error {
	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStorage) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
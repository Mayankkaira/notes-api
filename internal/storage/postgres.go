package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Mayankkaira/notes-api/application"
	"github.com/Mayankkaira/notes-api/note"
)

type PostgresStorage struct {
	db *sql.DB
}

var _ application.Storage = (*PostgresStorage)(nil)

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	newStorage := &PostgresStorage{db: db}
	return newStorage
}

func (p *PostgresStorage) Create(ctx context.Context, input application.CreateNoteInput) (note.Note, error) {
	var id int
	query := `INSERT INTO notes(title,content,created_at,updated_at)VALUES($1,$2,$3,$4)RETURNING id;`
	err := p.db.QueryRowContext(ctx, query, input.Title, input.Content, input.CreatedAt, input.UpdatedAt).Scan(&id)
	if err != nil {
		return note.Note{}, fmt.Errorf("create note: %w", err)
	}
	createdNote := note.Note{
		ID:        id,
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
	return createdNote, nil
}

func (p *PostgresStorage) Get(ctx context.Context, id int) (note.Note, error) {
	query := `SELECT id,title,content,created_at,updated_at FROM notes WHERE id=$1;`
	var n note.Note
	err := p.db.QueryRowContext(ctx, query, id).Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return note.Note{}, application.ErrStorageNotFound
	}
	if err != nil {
		return note.Note{}, fmt.Errorf("get note error %w", err)
	}
	return n, nil
}

func (p *PostgresStorage) List(ctx context.Context) ([]note.Note, error) {
	notes := make([]note.Note, 0)
	query := `select id,title,content,created_at,updated_at from notes order by id asc`
	rows, err := p.db.QueryContext(ctx, query)
	if err!=nil{
		return notes,fmt.Errorf("list note: %w",err)
	}
	defer rows.Close()
    for rows.Next(){
	var n note.Note
          err:=rows.Scan(&n.ID,&n.Title,&n.Content,&n.CreatedAt,&n.UpdatedAt)
          if err!=nil{
            return nil, fmt.Errorf("row scan %w",err)
          }
          notes=append(notes,n)
    }
    if err:=rows.Err();err!=nil{
         return []note.Note{},fmt.Errorf("row iteration %w",err)
    }
    return notes,nil
}

func(p * PostgresStorage)Update(ctx context.Context, ID int, updateInput application.UpdateNoteInput) (note.Note, error){
	var n note.Note
	query:=`update  notes set title=coalesce($1,title),content=coalesce($2,content),updated_at=$3 where id=$4 RETURNING id ,title,content,created_at,updated_at;`
	err:=p.db.QueryRowContext(ctx,query,updateInput.Title,updateInput.Content,updateInput.UpdatedAt,ID).Scan(&n.ID,&n.Title,&n.Content,&n.CreatedAt,&n.UpdatedAt)
	if errors.Is(err,sql.ErrNoRows){
		return note.Note{},application.ErrStorageNotFound
	}
	if err!=nil{
		return note.Note{},fmt.Errorf("update note: %w",err)
	}
    return n,nil
}

func (p *PostgresStorage)Delete(ctx context.Context, ID int) error{
	query:=`delete from notes where id=$1`
	res,err:=p.db.ExecContext(ctx,query,ID)
	if err!=nil{
		return fmt.Errorf("delete note: %w",err)
	}
	rowsAffected,err:=res.RowsAffected()
	if err!=nil{
		return fmt.Errorf("row affected %w",err)
	}
	if rowsAffected==0{
		return application.ErrStorageNotFound
	}
	return nil
}

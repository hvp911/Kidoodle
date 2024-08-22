package service

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"

	"content/internal/routes/contents"
)

const (
	getContentsQuery = `WITH selectContents AS (
					SELECT c.id AS content_id,
						c.protection_system,
						c.encryption_key,
						encode(c.encrypted_payload, 'hex') AS encrypted_payload
					FROM contents c
					ORDER BY c.id)`
)

// Reader abstraction of database connection
type Reader struct {
	dbConnection *sql.DB
}

// NewReaderService creates a service with db connection
func NewReaderService(dbConnection *sql.DB) *Reader {
	return &Reader{dbConnection: dbConnection}
}

// GetContent Retrieve contents from database and returns list of Content objects back to the handler
func (r Reader) GetContent(ctx context.Context) ([]contents.Content, error) {
	getContentsQueryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("content_id",
		"protection_system",
		"encryption_key",
		"encrypted_payload").
		From("selectContents").
		Prefix(getContentsQuery)

	getContentsSql, args, err := getContentsQueryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.dbConnection.QueryContext(ctx, getContentsSql, args...)
	if err != nil {
		return nil, err
	}

	// Once done with reading rows, we use defer to close the row.
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("failed to close rows: %v", err)
		}
	}(rows)

	var contentList []contents.Content
	for rows.Next() {
		var content contents.Content
		var sqlEncryptionKey sql.NullString
		var sqlEncryptedPayload sql.NullString
		err = rows.Scan(
			&content.ID,
			&content.ProtectionSystemID,
			&sqlEncryptionKey,
			&sqlEncryptedPayload,
		)
		if err != nil {
			return nil, err
		}
		if sqlEncryptionKey.Valid {
			content.EncryptionKey = sqlEncryptionKey.String
		}
		if sqlEncryptedPayload.Valid {
			content.EncryptedPayload = sqlEncryptedPayload.String
		}
		contentList = append(contentList, content)
	}
	return contentList, nil
}

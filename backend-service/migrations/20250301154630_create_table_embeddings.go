package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20250301154630",
		Up:      mig_20250301154630_create_table_embeddings_up,
		Down:    mig_20250301154630_create_table_embeddings_down,
	})
}

func mig_20250301154630_create_table_embeddings_up(tx *sql.Tx) error {
	// Create embeddings table
	_, err := tx.Exec(`
        CREATE TABLE embeddings (
            uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
            content TEXT NOT NULL,
            embedding vector(1536) NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        );
    `)
	if err != nil {
		return err
	}

	// Create index for similarity search
	_, err = tx.Exec("CREATE INDEX ON embeddings USING ivfflat (embedding vector_cosine_ops) WITH (lists = 100);")
	return err
}

func mig_20250301154630_create_table_embeddings_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS embeddings;")
	return err
}

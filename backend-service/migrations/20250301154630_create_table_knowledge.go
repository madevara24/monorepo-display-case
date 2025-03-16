package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20250301154630",
		Up:      mig_20250301154630_create_table_knowledge_up,
		Down:    mig_20250301154630_create_table_knowledge_down,
	})
}

func mig_20250301154630_create_table_knowledge_up(tx *sql.Tx) error {
	// Create knowledge table
	_, err := tx.Exec(`
        CREATE TABLE knowledge (
            uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
            category TEXT,
			granularity TEXT,
			content TEXT,
            embedding vector(1536) NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        );
    `)
	if err != nil {
		return err
	}

	// Create index for similarity search
	_, err = tx.Exec("CREATE INDEX ON knowledge USING ivfflat (embedding vector_cosine_ops) WITH (lists = 100);")
	return err
}

func mig_20250301154630_create_table_knowledge_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS knowledge;")
	return err
}

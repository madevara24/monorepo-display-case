package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20250315224212",
		Up:      mig_20250315224212_create_table_questions_up,
		Down:    mig_20250315224212_create_table_questions_down,
	})
}

func mig_20250315224212_create_table_questions_up(tx *sql.Tx) error {
	// Create questions table
	_, err := tx.Exec(`
        CREATE TABLE questions (
            uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
            question TEXT,
            embedding vector(1536) NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        );
    `)
	return err
}

func mig_20250315224212_create_table_questions_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS questions;")
	return err
}

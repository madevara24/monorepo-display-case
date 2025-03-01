package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20250301154608",
		Up:      mig_20250301154608_add_extension_pgvector_up,
		Down:    mig_20250301154608_add_extension_pgvector_down,
	})
}

func mig_20250301154608_add_extension_pgvector_up(tx *sql.Tx) error {
	// Enable pgvector extension
	_, err := tx.Exec("CREATE EXTENSION IF NOT EXISTS vector;")
	return err
}

func mig_20250301154608_add_extension_pgvector_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP EXTENSION IF EXISTS vector;")
	return err
}

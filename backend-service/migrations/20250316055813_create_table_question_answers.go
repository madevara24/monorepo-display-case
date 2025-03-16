package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20250316055813",
		Up:      mig_20250316055813_create_table_question_answers_up,
		Down:    mig_20250316055813_create_table_question_answers_down,
	})
}

func mig_20250316055813_create_table_question_answers_up(tx *sql.Tx) error {
	// Create question_answers table
	_, err := tx.Exec(`
        CREATE TABLE question_answers (
            uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
            question_uuid UUID REFERENCES questions(uuid),
            knowledge_uuid UUID REFERENCES knowledge(uuid),
            score FLOAT NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (question_uuid) REFERENCES questions(uuid) ON DELETE CASCADE,
            FOREIGN KEY (knowledge_uuid) REFERENCES knowledge(uuid) ON DELETE CASCADE
        );
    `)
	return err
}

func mig_20250316055813_create_table_question_answers_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS question_answers;")
	return err
}

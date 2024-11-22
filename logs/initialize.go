package logs

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open("sqlite3", "db_logs.db")
	if err != nil {
		log.Fatalf("erro ao abrir o banco de dados: %v\n", err)
	}

	createTables()

}

func createTables() {

	query := `

		-- Tabela para armazenar logs
		CREATE TABLE IF NOT EXISTS logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			timestamp DATETIME,
			level TEXT,
			status TEXT,
			message TEXT
		);

		-- Tabela para armazenar usuários
		CREATE TABLE IF NOT EXISTS users (
			email TEXT PRIMARY KEY,
			name TEXT,
			username TEXT
		);

		-- Tabela para armazenar repositórios
		CREATE TABLE IF NOT EXISTS repositories (
			id INTEGER PRIMARY KEY,
			name TEXT,
			full_name TEXT,
			git_url TEXT,
			allow_forking BOOLEAN,
			archive_url TEXT,
			assignees_url TEXT,
			blobs_url TEXT,
			branches_url TEXT,
			clone_url TEXT,
			collaborators_url TEXT,
			comments_url TEXT,
			commits_url TEXT,
			compare_url TEXT,
			contents_url TEXT,
			contributors_url TEXT,
			created_at REAL,
			default_branch TEXT,
			deployments_url TEXT,
			description TEXT,
			downloads_url TEXT,
			events_url TEXT,
			fork BOOLEAN,
			forks INTEGER,
			forks_count INTEGER,
			forks_url TEXT,
			html_url TEXT,
			language TEXT,
			languages_url TEXT,
			merges_url TEXT,
			milestones_url TEXT,
			mirror_url TEXT,
			node_id TEXT,
			notifications_url TEXT,
			open_issues INTEGER,
			open_issues_count INTEGER,
			owner_email TEXT,  -- Foreign key to users table
			private BOOLEAN,
			pulls_url TEXT,
			pushed_at REAL,
			releases_url TEXT,
			size INTEGER,
			ssh_url TEXT,
			stargazers INTEGER,
			stargazers_count INTEGER,
			stargazers_url TEXT,
			statuses_url TEXT,
			subscribers_url TEXT,
			subscription_url TEXT,
			svn_url TEXT,
			tags_url TEXT,
			teams_url TEXT,
			trees_url TEXT,
			updated_at TEXT,
			url TEXT,
			visibility TEXT,
			watchers INTEGER,
			watchers_count INTEGER,
			web_commit_signoff_required BOOLEAN,
			FOREIGN KEY (owner_email) REFERENCES users (email)
		);

		-- Tabela para armazenar commits
		CREATE TABLE IF NOT EXISTS commits (
			id TEXT PRIMARY KEY,
			message TEXT,
			timestamp DATETIME,
			author_email TEXT,
			committer_email TEXT,
			is_distinct BOOLEAN,
			tree_id TEXT,
			url TEXT,
			repository_id INTEGER,
			FOREIGN KEY (author_email) REFERENCES users (email),
			FOREIGN KEY (committer_email) REFERENCES users (email),
			FOREIGN KEY (repository_id) REFERENCES repositories (id)
		);

		-- Tabela para armazenar o webhook payload
		CREATE TABLE IF NOT EXISTS webhook_payloads (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			after TEXT,
			before TEXT,
			compare TEXT,
			created BOOLEAN,
			deleted BOOLEAN,
			forced BOOLEAN,
			head_commit_id TEXT,
			pusher_email TEXT,
			ref TEXT,
			repository_id INTEGER,
			sender_email TEXT,
			FOREIGN KEY (head_commit_id) REFERENCES commits (id),
			FOREIGN KEY (pusher_email) REFERENCES users (email),
			FOREIGN KEY (sender_email) REFERENCES users (email),
			FOREIGN KEY (repository_id) REFERENCES repositories (id)
		);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("erro ao criar tabelas de logs: %v\n", err)
	}

}

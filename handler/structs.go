package handler

import "time"

// Definição das structs para mapear o payload completo
type WebhookPayload struct {
	After      string     `json:"after"`
	Before     string     `json:"before"`
	Commits    []Commit   `json:"commits"`
	Compare    string     `json:"compare"`
	Created    bool       `json:"created"`
	Deleted    bool       `json:"deleted"`
	Forced     bool       `json:"forced"`
	HeadCommit Commit     `json:"head_commit"`
	Pusher     User       `json:"pusher"`
	Ref        string     `json:"ref"`
	Repository Repository `json:"repository"`
	Sender     User       `json:"sender"`
}

type Commit struct {
	Added     []string  `json:"added"`
	Author    User      `json:"author"`
	Committer User      `json:"committer"`
	Distinct  bool      `json:"distinct"`
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Modified  []string  `json:"modified"`
	Removed   []string  `json:"removed"`
	Timestamp time.Time `json:"timestamp"`
	TreeID    string    `json:"tree_id"`
	URL       string    `json:"url"`
}

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type Repository struct {
	AllowForking             bool     `json:"allow_forking"`
	ArchiveURL               string   `json:"archive_url"`
	AssigneesURL             string   `json:"assignees_url"`
	BlobsURL                 string   `json:"blobs_url"`
	BranchesURL              string   `json:"branches_url"`
	CloneURL                 string   `json:"clone_url"`
	CollaboratorsURL         string   `json:"collaborators_url"`
	CommentsURL              string   `json:"comments_url"`
	CommitsURL               string   `json:"commits_url"`
	CompareURL               string   `json:"compare_url"`
	ContentsURL              string   `json:"contents_url"`
	ContributorsURL          string   `json:"contributors_url"`
	CreatedAt                float64  `json:"created_at"`
	DefaultBranch            string   `json:"default_branch"`
	DeploymentsURL           string   `json:"deployments_url"`
	Description              *string  `json:"description"` // Usado *string para permitir nil
	DownloadsURL             string   `json:"downloads_url"`
	EventsURL                string   `json:"events_url"`
	Fork                     bool     `json:"fork"`
	Forks                    int      `json:"forks"`
	ForksCount               int      `json:"forks_count"`
	ForksURL                 string   `json:"forks_url"`
	FullName                 string   `json:"full_name"`
	GitCommitsURL            string   `json:"git_commits_url"`
	GitRefsURL               string   `json:"git_refs_url"`
	GitTagsURL               string   `json:"git_tags_url"`
	GitURL                   string   `json:"git_url"`
	HasDownloads             bool     `json:"has_downloads"`
	HasIssues                bool     `json:"has_issues"`
	HasPages                 bool     `json:"has_pages"`
	HasProjects              bool     `json:"has_projects"`
	HasWiki                  bool     `json:"has_wiki"`
	Homepage                 *string  `json:"homepage"`
	HooksURL                 string   `json:"hooks_url"`
	HTMLURL                  string   `json:"html_url"`
	ID                       int      `json:"id"`
	IsTemplate               bool     `json:"is_template"`
	IssueCommentURL          string   `json:"issue_comment_url"`
	IssueEventsURL           string   `json:"issue_events_url"`
	IssuesURL                string   `json:"issues_url"`
	KeysURL                  string   `json:"keys_url"`
	LabelsURL                string   `json:"labels_url"`
	Language                 *string  `json:"language"`
	LanguagesURL             string   `json:"languages_url"`
	License                  *string  `json:"license"`
	MergesURL                string   `json:"merges_url"`
	MilestonesURL            string   `json:"milestones_url"`
	MirrorURL                *string  `json:"mirror_url"`
	Name                     string   `json:"name"`
	NodeID                   string   `json:"node_id"`
	NotificationsURL         string   `json:"notifications_url"`
	OpenIssues               int      `json:"open_issues"`
	OpenIssuesCount          int      `json:"open_issues_count"`
	Owner                    User     `json:"owner"`
	Private                  bool     `json:"private"`
	PullsURL                 string   `json:"pulls_url"`
	PushedAt                 float64  `json:"pushed_at"`
	ReleasesURL              string   `json:"releases_url"`
	Size                     int      `json:"size"`
	SSHURL                   string   `json:"ssh_url"`
	Stargazers               int      `json:"stargazers"`
	StargazersCount          int      `json:"stargazers_count"`
	StargazersURL            string   `json:"stargazers_url"`
	StatusesURL              string   `json:"statuses_url"`
	SubscribersURL           string   `json:"subscribers_url"`
	SubscriptionURL          string   `json:"subscription_url"`
	SVNURL                   string   `json:"svn_url"`
	TagsURL                  string   `json:"tags_url"`
	TeamsURL                 string   `json:"teams_url"`
	Topics                   []string `json:"topics"`
	TreesURL                 string   `json:"trees_url"`
	UpdatedAt                string   `json:"updated_at"`
	URL                      string   `json:"url"`
	Visibility               string   `json:"visibility"`
	Watchers                 int      `json:"watchers"`
	WatchersCount            int      `json:"watchers_count"`
	WebCommitSignoffRequired bool     `json:"web_commit_signoff_required"`
}

type LogEntry struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
}

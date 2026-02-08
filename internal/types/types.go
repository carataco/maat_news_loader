package types

type PostgresConfig struct {
	Host   string
	User   string
	Pwd    string
	Schema string
	Table  string
}

type LoadConfig struct {
	PG PostgresConfig
}

type ExtractConfig struct {
	S3 S3Config
}

type S3Config struct {
	Bucket string
}

type Config struct {
	SourceType string
	SourceIDS  []string
	TargetType string
	AWSRegion  string
	S3Bucket   string
	Extract    ExtractConfig
	Load       LoadConfig
}

type Event struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

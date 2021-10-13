package configuration

type Configuration struct {
	LogFilePath    string `json:"log_file_path"`
	BackupFilePath string `json:"backup_file_path"`
	BackupInterval int    `json:"backup_interval"`
}

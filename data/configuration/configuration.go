package configuration

type Configuration struct {
	BackupFilePath string `json:"backup_file_path"`
	BackupInterval int    `json:"backup_interval"`
}

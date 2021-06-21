package repository

type FilesystemRepository interface {
	Move(fromPath string, toPath string, filename string)
	Run(path string, filename string)
}

type filesystemRepository struct {
}

func NewFilesystemRepository() FilesystemRepository {
	return &filesystemRepository{}
}

func (r *filesystemRepository) Move(fromPath string, toPath string, filename string) {

}
func (r *filesystemRepository) Run(path string, filename string) {

}

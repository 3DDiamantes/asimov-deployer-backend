package repository

type FilesystemRepository interface {
	Move(path string)
	Run(path string, app string)
}

type filesystemRepository struct {
}

func NewFilesystemRepository() FilesystemRepository {
	return &filesystemRepository{}
}

func (r *filesystemRepository) Move(path string) {

}
func (r *filesystemRepository) Run(path string, app string) {

}

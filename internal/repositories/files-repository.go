package repositories

import "github.com/marechal-dev/adoptgram-files/internal/entities"

type FilesRepository interface {
	create(file *entities.File)
}

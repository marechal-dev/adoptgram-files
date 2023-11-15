package services

import (
	"fmt"

	"github.com/google/uuid"
)

func GetFileKeyService(fileName string) (string, error) {
	fileUUID, err := uuid.NewUUID()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s", fileUUID.String(), fileName), nil
}

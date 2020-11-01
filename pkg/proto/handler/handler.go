package handler

import (
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func getErrorParser(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Resouce Not Found: %v", err))
	}
	return status.Errorf(codes.Internal,
		fmt.Sprintf("Unknown internal error: %v", err))
}

func deleteErrorParser(err error) error {
	if _, ok := err.(*orm.NoRowsAffected); ok {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Resouce Not Found: %v", err),
		)
	} else {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
}

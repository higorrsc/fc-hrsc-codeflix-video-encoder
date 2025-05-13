package domain_test

import (
	"testing"
	"time"

	"github.com/higorrsc/fc-hrsc-codeflix-video-encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotAUUID(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "abc"
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoResourceIdIsEmpty(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = "resource_id"
	video.FilePath = "path/to/file"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}

package utils_test

import (
	"testing"

	"github.com/higorrsc/fc-hrsc-codeflix-video-encoder/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{
				"id": "273c831b-1378-58e4-a063-5ef2571bdbc4",
				"file_path": "convite.mp4",
				"status": "pending"
			}`
	err := utils.IsJson(json)
	require.Nil(t, err)

	json = `wes`
	err = utils.IsJson(json)
	require.Error(t, err)
}

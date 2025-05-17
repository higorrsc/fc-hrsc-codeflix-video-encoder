package services

import (
	"io"
	"log"
	"os"

	"github.com/higorrsc/fc-hrsc-codeflix-video-encoder/application/repositories"
	"github.com/higorrsc/fc-hrsc-codeflix-video-encoder/domain"

	"context"

	"cloud.google.com/go/storage"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(bucketName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	bucket := client.Bucket(bucketName)
	object := bucket.Object(v.Video.FilePath)
	reader, err := object.NewReader(ctx)
	if err != nil {
		return err
	}

	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	file, err := os.Create(os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.ID + ".mp4")
	if err != nil {
		return err
	}

	_, err = file.Write(body)
	if err != nil {
		return err
	}

	defer file.Close()
	log.Printf("Video %s downloaded from bucket %s", v.Video.ID, bucketName)

	return nil
}

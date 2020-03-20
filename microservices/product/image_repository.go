package product

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio-go/v6"
)

type ImageRepository interface {
	Store(image []byte) (url string, err error)
	Delete(url string) error

	InitRepository() error
}

type imageRepositoryMinIO struct {
	client        *minio.Client
	bucketName    string
	publicBaseURL string
	hostname      string
}

// TODO: read via io.Reader (how to get image md5 sum and size?)
func (ir *imageRepositoryMinIO) Store(image []byte) (string, error) {
	if len(image) == 0 {
		return "", fmt.Errorf("empty image is not allowed")
	}
	if len(image) >= 128_000_000 {
		return "", fmt.Errorf("image size is too large: %d byte", len(image))
	}
	var extension, contentType string
	switch image[0] {
	case byte(0xFF):
		extension = "jpg"
		contentType = "image/jpeg"
	case byte(0x89):
		extension = "png"
		contentType = "image/png"
	case byte(0x47):
		extension = "gif"
		contentType = "image/gif"
	default:
		extension = "unknown"
		contentType = "application/octet-stream"
	}

	tmp := make([]byte, 0, len(image)+50)
	tmp = append(tmp, []byte(strconv.Itoa(int(time.Now().UnixNano())))...)
	tmp = append(tmp, image...)
	filename := fmt.Sprintf("%X.%s", md5.Sum(tmp), extension)

	r := bytes.NewReader(image)
	if _, err := ir.client.PutObject(
		ir.bucketName,
		filename,
		r,
		int64(len(image)),
		minio.PutObjectOptions{ContentType: contentType},
	); err != nil {
		return "", err
	}

	url, err := ir.FilenameToURL(filename)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (ir *imageRepositoryMinIO) Delete(url string) error {
	filename, err := ir.URLToFilename(url)
	if err != nil {
		return err
	}
	if err := ir.client.RemoveObject(ir.bucketName, filename); err != nil {
		return err
	}
	return nil
}

func (ir *imageRepositoryMinIO) FilenameToURL(filename string) (string, error) {
	return fmt.Sprintf("%s/%s", ir.publicBaseURL, filename), nil
}

func (ir *imageRepositoryMinIO) URLToFilename(urlstr string) (string, error) {
	u, err := url.Parse(urlstr)
	if err != nil {
		return "", err
	}
	if u.Hostname() != ir.hostname {
		return "", fmt.Errorf("this is not own url")
	}

	pathArray := strings.Split(u.Path, "/")
	filename := pathArray[len(pathArray)-1]

	return filename, nil
}

func (ir *imageRepositoryMinIO) InitRepository() error {
	location := "none"
	if err := ir.client.MakeBucket(ir.bucketName, location); err != nil {
		exists, errBucketExists := ir.client.BucketExists(ir.bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Already own %s", ir.bucketName)
		} else {
			return err
		}
	} else {
		log.Printf("Successfully created %s", ir.bucketName)
	}

	polmap := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Sid":       "PublicRead",
				"Effect":    "Allow",
				"Principal": "*",
				"Action": []string{
					"s3:GetObject",
				},
				"Resource": []string{
					fmt.Sprintf("arn:aws:s3:::%s/*", ir.bucketName),
				},
			},
		},
	}
	pol, err := json.Marshal(polmap)
	if err != nil {
		return err
	}
	if err := ir.client.SetBucketPolicy(ir.bucketName, string(pol)); err != nil {
		return err
	}
	log.Printf("Apply public read bucket policy to %s", ir.bucketName)
	return nil
}

type ImageRepositoryMinIOConfig struct {
	Host            string
	Port            int
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	UseSSL          bool
	PublicHost      string
	PublicPort      int
}

func (c *ImageRepositoryMinIOConfig) Connect() (ImageRepository, error) {
	endpoint := fmt.Sprintf("%s:%d", c.Host, c.Port)
	client, err := minio.New(endpoint, c.AccessKeyID, c.SecretAccessKey, c.UseSSL)
	if err != nil {
		return nil, err
	}
	return &imageRepositoryMinIO{
		client:        client,
		bucketName:    c.BucketName,
		publicBaseURL: fmt.Sprintf("https://%s:%d/%s", c.PublicHost, c.PublicPort, c.BucketName),
		hostname:      c.PublicHost,
	}, nil
}

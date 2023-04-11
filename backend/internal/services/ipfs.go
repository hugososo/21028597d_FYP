package services

import (
	"bytes"
	"comp4913-backend/configs"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"path/filepath"
)
import ipfsApi "github.com/ipfs/go-ipfs-api"

type IPFSClient struct {
	shell *ipfsApi.Shell
}

func NewIPFSClient(projectId, projectSecret string) *IPFSClient {
	return &IPFSClient{ipfsApi.NewShellWithClient(configs.IPFS_URL, NewClient(projectId, projectSecret))}
}

// NewClient creates an http.Client that automatically perform basic auth on each request.
func NewClient(projectId, projectSecret string) *http.Client {
	return &http.Client{
		Transport: authTransport{
			RoundTripper:  http.DefaultTransport,
			ProjectId:     projectId,
			ProjectSecret: projectSecret,
		},
	}
}

// authTransport decorates each request with a basic auth header.
type authTransport struct {
	http.RoundTripper
	ProjectId     string
	ProjectSecret string
}

func (t authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.ProjectId, t.ProjectSecret)
	return t.RoundTripper.RoundTrip(r)
}

func (s IPFSClient) UploadItem(data []byte) string {
	reader := bytes.NewReader(data)
	cid, err := s.shell.Add(reader)
	if err != nil {
		log.Fatal(err)
	}
	return cid
}

func (s IPFSClient) UploadFolder(folderPath string) string {
	cid, err := s.addFolderToIPFS(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	return cid
}

func (s IPFSClient) addFolderToIPFS(folderPath string) (string, error) {

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = s.shell.Add(file)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	cid, err := s.shell.AddDir(folderPath)
	if err != nil {
		return "", err
	}

	return cid, nil
}

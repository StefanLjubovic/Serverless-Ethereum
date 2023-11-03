package service

import (
	model "backend/model"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ipfs/go-cid"
	"github.com/web3-storage/go-w3s-client"
)

type NFTService struct{}

func NewNFTHandler() *NFTService {
	return &NFTService{}
}

func (nftService *NFTService) ReceiveCertificate(user *model.User, course *model.Course, img []byte) (string, error) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweDk5MDQzODQ3ZDhGNzMwRjU3NmVFNmRiNjhENjIxQTE0NmIwODgzNDciLCJpc3MiOiJ3ZWIzLXN0b3JhZ2UiLCJpYXQiOjE2OTg2NjA0MTc5MDMsIm5hbWUiOiJzZXJ2ZXJsZXNzIn0.ZgWRdtfmOh4VZuFs5CwCCV9sVrItzE7VtEKAt7bIumA"
	client, err := w3s.NewClient(w3s.WithToken(token))
	if err != nil {
		fmt.Println(err)
	}
	certFileName := getFileName(course.Certificate.ImagePath)
	cid1, err, newNameCert := writeBytesToFile(img, certFileName, client)
	if err != nil {
		fmt.Println(err)
	}
	path := "https://" + cid1.String() + ".ipfs.w3s.link/" + newNameCert
	metadata, err := generateNFTMetadata(course.Certificate, path)
	cid, err, name := writeBytesToFile(metadata, "data*.json", client)
	retVal := "https://" + cid.String() + ".ipfs.w3s.link/" + name
	return retVal, nil
}

func getFileName(filename string) string {
	urlParts := strings.Split(filename, "/")
	fileName := urlParts[len(urlParts)-1]
	fileNameParts := strings.Split(fileName, ".")
	name := fileNameParts[0]
	ext := "." + fileNameParts[len(fileNameParts)-1]
	return fmt.Sprintf("%s*%s", name, ext)
}

func writeBytesToFile(data []byte, name string, client w3s.Client) (cid.Cid, error, string) {
	fmt.Println(data)
	tmpFile, err := ioutil.TempFile("", name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tmpFile.Name())
	defer os.Remove(tmpFile.Name())
	_, _ = tmpFile.Write(data)

	_, _ = tmpFile.Seek(0, 0)
	parts := strings.Split(tmpFile.Name(), "/")
	cid, err := client.Put(context.Background(), tmpFile)
	return cid, err, parts[len(parts)-1]
}

func generateNFTMetadata(cert model.Certificate, path string) ([]byte, error) {
	nftMetadata := model.NFTMetadataScheme{
		Title: "Asset Metadata",
		Type:  "object",
		Properties: model.NFTMetadata{
			Name: model.NFTMetadataProperty{
				Type:        "string",
				Description: cert.Name,
			},
			Description: model.NFTMetadataProperty{
				Type:        "string",
				Description: cert.Description,
			},
			Image: model.NFTMetadataProperty{
				Type:        "string",
				Description: path,
			},
		},
	}
	return json.MarshalIndent(nftMetadata, "", "\t")
}

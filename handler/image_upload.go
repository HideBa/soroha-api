package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/HideBa/soroha-api/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iam/v1"
)

var (
	iamService         *iam.Service
	serviceAccountName string
	serviceAccountID   string
	uploadableBucket   string
)

func (h *Handler) SignHandler(c echo.Context) error {
	contentType := c.FormValue("content_type")
	if contentType == "" {
		return c.JSON(http.StatusBadRequest, "error: you must set content type")
	}
	key := uuid.New().String()
	if ext := c.FormValue("ext"); ext != "" {
		key += fmt.Sprintf(".%s", ext)
	}

	url, err := storage.SignedURL(uploadableBucket, key, &storage.SignedURLOptions{
		GoogleAccessID: serviceAccountName,
		Method:         "PUT",
		Expires:        time.Now().Add(15 * time.Minute),
		ContentType:    contentType,
		SignBytes: func(b []byte) ([]byte, error) {
			response, err := iamService.Projects.ServiceAccounts.SignBlob(serviceAccountID, &iam.SignBlobRequest{BytesToSign: base64.StdEncoding.EncodeToString(b)},).Context(c.Request().Context()).Do()
		}
		if err != nil {
			return nil, err
		}
		return base64.StdEncoding.DecodeString(response.Signature)
	})
	if err != nil {
		log.Printf("sign: failed to sign, err = %v\n", err)
		return c.JSON(http.StatusInternalServerError, "failure to sign")
	}
	return c.JSON(http.StatusOK)
}

func main() {
	cred, err := google.DefaultClient(context.Background(), iam.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}
	iamService, err = iam.New(cred)
	if err != nil {
		log.Fatal(err)
	}

	uploadableBucket = os.Getenv("UPLOADABLE_BUCKET")
	serviceAccountName = os.Getenv("SERVICE_ACCOUNT")
	serviceAccountID = fmt.Sprintf("projects/%s/serviceAccounts/%s", os.Getenv("GOOGLE_CLOUD_PROJECT"), serviceAccountName)
}
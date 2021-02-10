package util

import "google.golang.org/api/iam/v1"

var (
	iamService         *iam.Service
	serviceAccountName string
	serviceAccountID   string
	uploadableBucket   string
)

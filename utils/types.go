package utils

import (
	"mime/multipart"
)

const (
	AWSREGION = "aws_region"
	AWSBUCKET = "aws_bucket_name"
	KEYID     = "aws_key_id"
	SECRETKEY = "aws_secret_key"
)

type (
	Response struct {
		Success bool        `json:"success"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
		Meta    *Pagination `json:"meta,omitempty"`
	}
	Pagination struct {
		CurrentPage    int `json:"current_page,omitempty"`
		TotalPages     int `json:"total_pages,omitempty"`
		TotalDataCount int `json:"total_data_count,omitempty"`
	}
	S3UploadReq struct {
		File        multipart.File
		FileName    string
		ContentType string
	}
)

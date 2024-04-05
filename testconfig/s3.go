package testconfig

import (
	"context"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Mock struct {
	GetObjectFunc           func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	PutObjectFunc           func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	GetObjectAttributesFunc func(ctx context.Context, params *s3.GetObjectAttributesInput, optFns ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error)
	ListObjectsV2Func       func(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	DeleteObjectFunc        func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
}

func (m S3Mock) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	if m.GetObjectFunc != nil {
		return m.GetObjectFunc(ctx, params, optFns...)
	}
	return &s3.GetObjectOutput{}, nil
}

func (m S3Mock) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	if m.PutObjectFunc != nil {
		return m.PutObjectFunc(ctx, params, optFns...)
	}
	return &s3.PutObjectOutput{}, nil
}

func (m S3Mock) GetObjectAttributes(ctx context.Context, params *s3.GetObjectAttributesInput, optFns ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error) {
	if m.GetObjectAttributesFunc != nil {
		return m.GetObjectAttributesFunc(ctx, params, optFns...)
	}
	return &s3.GetObjectAttributesOutput{}, nil
}

func (m S3Mock) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	if m.ListObjectsV2Func != nil {
		return m.ListObjectsV2Func(ctx, params, optFns...)
	}
	return &s3.ListObjectsV2Output{}, nil
}

func (m S3Mock) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	if m.DeleteObjectFunc != nil {
		return m.DeleteObjectFunc(ctx, params, optFns...)
	}
	return &s3.DeleteObjectOutput{}, nil
}

type PresignedS3Mock struct {
	PresignGetObjectFunc func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignPutObjectFunc func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

func (m PresignedS3Mock) PresignGetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if m.PresignGetObjectFunc != nil {
		return m.PresignGetObjectFunc(ctx, params, optFns...)
	}
	return &v4.PresignedHTTPRequest{}, nil
}

func (m PresignedS3Mock) PresignPutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if m.PresignPutObjectFunc != nil {
		return m.PresignPutObjectFunc(ctx, params, optFns...)
	}
	return &v4.PresignedHTTPRequest{}, nil
}

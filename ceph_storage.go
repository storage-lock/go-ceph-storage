package ceph_storage

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
)

// CephStorage ceph对象存储实现 (S3兼容)
type CephStorage struct {
	*s3_storage.S3Storage
}

func NewCephStorage(ctx context.Context, options *s3_storage.S3StorageOptions) (*CephStorage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &CephStorage{S3Storage: s3Storage}, nil
}

const StorageName = "ceph-storage"

func (x *CephStorage) GetName() string {
	return StorageName
}

// Capabilities 声明存储支持的能力

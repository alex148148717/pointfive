package domain

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
)

type ImportJobFileRepository interface {
	InsertImportJobFile(ctx context.Context, path string) (*ImportJobFile, error)
}
type FileRepository interface {
	UploadFile(ctx context.Context) (io.WriteCloser, string, error)
}
type ImportWorkerRepository interface {
	InsertWorkerJob(ctx context.Context, id int) (*WorkerJob, error)
}

type ServiceImpl struct {
	importJobFileRepository ImportJobFileRepository
	fileRepository          FileRepository
	importWorkerRepository  ImportWorkerRepository
}

func NewService(importJobFileRepository ImportJobFileRepository, fileRepository FileRepository, importWorkerRepository ImportWorkerRepository) *ServiceImpl {
	return &ServiceImpl{importJobFileRepository: importJobFileRepository, fileRepository: fileRepository, importWorkerRepository: importWorkerRepository}
}

func (c *ServiceImpl) ImportFile(ctx context.Context, in io.Reader) (*ImportJobFile, error) {

	out, path, err := c.fileRepository.UploadFile(ctx)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	err = TransformFileLineByLine(in, out, func(line []byte) []byte {
		return bytes.ToUpper(line)
	})
	if err != nil {
		return nil, err
	}

	importJobFile, err := c.importJobFileRepository.InsertImportJobFile(ctx, path)
	if err != nil {
		return nil, err
	}
	wj, err := c.importWorkerRepository.InsertWorkerJob(ctx, importJobFile.JobID)
	if err != nil {
		return nil, err
	}
	fmt.Printf("worker job data %+v\n", wj)
	return importJobFile, nil
}
func TransformFileLineByLine(r io.Reader, w io.WriteCloser, transform func([]byte) []byte) error {
	scanner := bufio.NewScanner(r)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	writer := bufio.NewWriter(w)
	defer writer.Flush()
	for scanner.Scan() {
		line := scanner.Bytes()
		modified := transform(line)
		if _, err := writer.Write(modified); err != nil {
			return fmt.Errorf("failed writing: %w", err)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}
	return nil
}

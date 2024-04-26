package service

import (
	"bufio"
	"case-study-job-service/internal/core/models"
	"case-study-job-service/internal/core/port"
	"case-study-job-service/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"
)

type JobService struct {
	workerCount       int
	productRepository port.ProductRepository
	recordRepository  port.RecordRepository
}

func NewJobService(workerCount int, productRepository port.ProductRepository, recordRepository port.RecordRepository) *JobService {
	return &JobService{
		workerCount:       workerCount,
		productRepository: productRepository,
		recordRepository:  recordRepository,
	}
}

func (s JobService) RecordProduct(ctx context.Context) error {
	l := logger.FromCtx(ctx).Sugar()

	objs, err := s.recordRepository.ListObjectsFromBucket()
	if err != nil {
		l.Errorln(err)
		return err
	}

	for i := 0; i < len(objs); i++ {
		rawReq, err := s.recordRepository.GetObject(objs[i])
		if err != nil {
			l.Errorln(err)
			return err
		}

		go s.recordProduct(ctx, rawReq.Body)
	}

	return nil
}

func (s JobService) recordProduct(ctx context.Context, body io.ReadCloser) {
	jobs := make(chan string)
	wg := &sync.WaitGroup{}

	for w := 0; w < s.workerCount; w++ {
		wg.Add(1)
		go s.writeProcessing(ctx, jobs, wg)
	}

	go func() {
		scanner := bufio.NewScanner(body)
		for scanner.Scan() {
			st := scanner.Text()
			if json.Valid([]byte(st)) {
				jobs <- st
			}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
	}()
}

func (s JobService) writeProcessing(ctx context.Context, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	l := logger.FromCtx(ctx).Sugar()

	var productReq = &models.ProductRequest{}
	for j := range jobs {
		if err := json.Unmarshal([]byte(j), productReq); err != nil {
			l.Error(err)
		}

		key := fmt.Sprintf("%d", productReq.RecordID)

		ok, err := s.productRepository.GetBody(ctx, key, productReq)
		if err != nil {
			l.Error(err)
		}

		if !ok {
			if err := s.productRepository.Set(ctx, key, j, 0).Err(); err != nil {
				l.Error(err)
			}
		}
	}
}

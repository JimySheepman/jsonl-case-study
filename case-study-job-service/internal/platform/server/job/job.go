package server

import (
	"case-study-job-service/internal/core/port"
	"context"
)

type JobServer struct {
	jobHandlerClient port.JobHandlerClient
}

func NewJobServer(jobHandlerClient port.JobHandlerClient) *JobServer {
	return &JobServer{
		jobHandlerClient: jobHandlerClient,
	}
}

func (s JobServer) Run(ctx context.Context) error {
	return s.jobHandlerClient.Run(ctx)
}

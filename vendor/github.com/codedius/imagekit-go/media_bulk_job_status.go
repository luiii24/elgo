package imagekit

import (
	"context"
	"errors"
)

//
// RESPONSES
//

type BulkJobStatusResponse struct {
	// JobID you get in the response of bulk job API e.g. copy folder or move folder API.
	JobID string `json:"jobId"`
	// Type of operation, it could be either COPY_FOLDER or MOVE_FOLDER.
	Type string `json:"type"`
	// Status of the job.
	//
	// It can be either:
	// Pending - The job has been successfully submitted and is in progress.
	// Completed - The job has been completed.
	Status string `json:"status"`
}

//
// METHODS
//

// BulkJobStatus will copy one folder into another.
func (s *MediaService) BulkJobStatus(ctx context.Context, jid string) (*BulkJobStatusResponse, error) {
	if jid == "" {
		return nil, errors.New("file id is empty")
	}

	// Prepare request
	req, err := s.client.request("GET", "v1/bulkJobs/"+jid, nil, requestTypeAPI)
	if err != nil {
		return nil, err
	}

	// Submit the request
	res := new(BulkJobStatusResponse)

	err = s.client.do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

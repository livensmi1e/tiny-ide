package server

import (
	"context"

	pb "github.com/livensmi1e/tiny-ide/executor/proto"
)

const (
	PORT = "50001"
)

type Server struct {
	pb.UnimplementedExecutorServer
}

func (s *Server) Execute(ctx context.Context, req *pb.ExecuteRequest) (*pb.ExecuteResponse, error) {
	e := newEngine(req.SubmissionId, req.LanguageId, req.SourceCode)
	e.prepare()
	e.run()
	e.cleanup()
	if e.error() != nil {
		return nil, e.error()
	}
	return &pb.ExecuteResponse{
		SubmissionId:  req.SubmissionId,
		Stdout:        e.metadata.stdout,
		Stderr:        e.metadata.stderr,
		CompileOutput: "",
		Status:        "",
		Time:          "",
		Memory:        "",
	}, nil
}

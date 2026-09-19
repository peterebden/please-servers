package worker

import (
	"testing"

	pb "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	"github.com/stretchr/testify/assert"
)

func TestAllOutputSymlinks(t *testing.T) {
	file := &pb.OutputSymlink{Path: "file_link", Target: "file"}
	dir := &pb.OutputSymlink{Path: "dir_link", Target: "dir"}

	// No symlinks must stay nil so the serialised ActionResult doesn't change.
	assert.Nil(t, allOutputSymlinks(&pb.ActionResult{}))

	// Only the deprecated fields are populated
	assert.Equal(t, []*pb.OutputSymlink{file, dir}, allOutputSymlinks(&pb.ActionResult{
		OutputFileSymlinks:      []*pb.OutputSymlink{file},
		OutputDirectorySymlinks: []*pb.OutputSymlink{dir},
	}))

	// output_symlinks is already populated, which takes priority
	assert.Equal(t, []*pb.OutputSymlink{dir, file}, allOutputSymlinks(&pb.ActionResult{
		OutputSymlinks:          []*pb.OutputSymlink{dir, file},
		OutputFileSymlinks:      []*pb.OutputSymlink{file},
		OutputDirectorySymlinks: []*pb.OutputSymlink{dir},
	}))
}

package client_test

import (
	"context"
	"crypto/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/client"
	"github.com/data-preservation-programs/singularity/client/testutil"
	"github.com/data-preservation-programs/singularity/model"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/stretchr/testify/require"
)

func TestClients(t *testing.T) {
	ctx := context.Background()
	testutil.TestWithAllClients(ctx, t, func(t *testing.T, client client.Client) {
		// createDataset
		ds, err := client.CreateDataset(ctx, dataset.CreateRequest{
			Name:       "test",
			MaxSizeStr: "31.5GiB",
		})
		require.NoError(t, err)
		require.Equal(t, "test", ds.Name)

		// cannot create dataset with same name

		dupDataset, err := client.CreateDataset(ctx, dataset.CreateRequest{
			Name:       "test",
			MaxSizeStr: "31.5GiB",
		})
		var asConflict handler.DuplicateRecordError
		require.ErrorAs(t, err, &asConflict)
		require.Nil(t, dupDataset)

		// cannot create dataset with invalid parameter
		invalidDataset, err := client.CreateDataset(ctx, dataset.CreateRequest{})
		var asInvalidParameter handler.InvalidParameterError
		require.ErrorAs(t, err, &asInvalidParameter)
		require.Nil(t, invalidDataset)

		path := t.TempDir()
		// create datasource
		source, err := client.CreateLocalSource(ctx, "test", datasource.LocalRequest{
			SourcePath:     path,
			RescanInterval: "0h",
			ScanningState:  model.Ready,
		})
		require.NoError(t, err)
		require.Equal(t, "local", source.Type)
		require.Equal(t, ds.ID, source.DatasetID)
		require.Equal(t, path, source.Path)
		require.Equal(t, model.Ready, source.ScanningState)

		// list sources for dataset
		sources, err := client.ListSourcesByDataset(ctx, "test")
		require.NoError(t, err)
		require.Len(t, sources, 1)
		require.Equal(t, ds.ID, sources[0].DatasetID)
		require.Equal(t, path, sources[0].Path)
		require.Equal(t, model.Ready, sources[0].ScanningState)

		// create datasource when dataset not found
		notFoundSource, err := client.CreateLocalSource(ctx, "apples", datasource.LocalRequest{
			SourcePath:     path,
			RescanInterval: "0h",
		})
		var asNotFoundError handler.NotFoundError
		require.ErrorAs(t, err, &asNotFoundError)
		require.Nil(t, notFoundSource)

		// push item
		file, err := os.CreateTemp(path, "push-*")
		require.NoError(t, err)
		buf := make([]byte, 1000)
		_, _ = rand.Read(buf)
		file.Write(buf)
		name := file.Name()
		err = file.Close()
		require.NoError(t, err)
		item, err := client.PushItem(ctx, source.ID, datasource.ItemInfo{Path: filepath.Base(name)})
		require.NoError(t, err)
		require.Equal(t, filepath.Base(name), item.Path)

		// get item
		returnedItem, err := client.GetItem(ctx, item.ID)
		require.NoError(t, err)
		require.Equal(t, item.Path, returnedItem.Path)
	})
}

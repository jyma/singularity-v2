package datasetworker

import (
	"context"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
)

// scan scans the data source and inserts the chunking strategy back to database
// scanSource is true if the source will be actually scanned in addition to just picking up remaining ones
// resume is true if the scan will be resumed from the last scanned item, which is useful for resuming a failed scan
func (w *DatasetWorkerThread) scan(ctx context.Context, source model.Source, scanSource bool) error {
	dataset := *source.Dataset
	var remainingParts []model.ItemPart
	err := w.db.Joins("Item").
		Where("source_id = ? AND item_parts.chunk_id is null", source.ID).
		Order("item_parts.id asc").
		Find(&remainingParts).Error
	if err != nil {
		return err
	}

	if scanSource {
		sourceScanner, err := w.datasourceHandlerResolver.Resolve(ctx, source)
		if err != nil {
			return errors.Wrap(err, "failed to get source scanner")
		}
		entryChan := sourceScanner.Scan(ctx, "", source.LastScannedPath)
		for entry := range entryChan {
			if entry.Error != nil {
				w.logger.Errorw("failed to scan", "error", entry.Error)
				continue
			}

			item, itemParts, err := datasource.PushItem(ctx, w.db, entry.Info, source, dataset, w.directoryCache)
			if err != nil {
				return errors.Wrap(err, "failed to push item")
			}
			if item == nil {
				w.logger.Infow("item already exists", "path", entry.Info.Remote())
				continue
			}
			err = database.DoRetry(func() error {
				return w.db.Model(&model.Source{}).Where("id = ?", source.ID).
					Update("last_scanned_path", item.Path).Error
			})
			if err != nil {
				return errors.Wrap(err, "failed to update last scanned path")
			}
			remainingParts = append(remainingParts, itemParts...)
		}
	}
	w.logger.With("remaining", len(remainingParts)).Info("remaining items")
	_, err = datasource.ChunkItemParts(w.db, source.ID, dataset.MaxSize, remainingParts)
	if err != nil {
		return err
	}
	return datasource.SourceChunksReadyHandler(w.db, source.ID)
}

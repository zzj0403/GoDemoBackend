package impl

import (
	"GoDemoBackend/apps/host"
	"context"
	"fmt"
)

func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)

	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Error("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Error("commit error, %s", err)
			}
		}
	}()
	// 插入Resource数据
	rstmt, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()

	_, err = rstmt.ExecContext(ctx,
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		return err
	}
	// 插入Describe 数据
	dstmt, err := tx.PrepareContext(ctx, InsertDescribeSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()

	_, err = dstmt.ExecContext(ctx,
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)
	if err != nil {
		return err
	}
	return nil
}

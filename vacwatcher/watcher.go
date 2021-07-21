package vacwatcher

import (
	"context"
	"errors"

	"github.com/lixin9311/vac-bot/ent"
	"github.com/lixin9311/vac-bot/ent/vacdepartment"
	"github.com/lixin9311/vac-bot/tokyovacapi"
)

var (
	ErrPartitionNotExist = errors.New("partition does not exist")
)

type Watcher struct {
	client *tokyovacapi.Client
	db     *ent.Client
}

func (w *Watcher) ValidateAndInitPartition(ctx context.Context, partition string) error {
	_, err := w.db.VacDepartment.Query().Where(vacdepartment.Partition(partition)).FirstID(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			resp, err := w.client.GetArticles(ctx, tokyovacapi.Partition(partition))
			if err != nil {
				return err
			} else if len(resp.Articles) == 0 {
				return ErrPartitionNotExist
			}
			return w.putPartition(ctx, tokyovacapi.Partition(partition))
		}
		return err
	}
	return nil
}

func (w *Watcher) ValidateLogin(ctx context.Context, partition, username, password string) (resp *tokyovacapi.LoginResponse, err error) {
	resp, err = w.client.Login(ctx, tokyovacapi.Partition(partition), username, password)
	// TODO: update reservation
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *Watcher) putPartition(ctx context.Context, partition tokyovacapi.Partition) error {
	resp, err := w.client.GetDepartments(ctx, partition)
	if err != nil {
		return err
	}
	if len(resp.Departments) == 0 {
		return ErrPartitionNotExist
	}
	bulk := make([]*ent.VacDepartmentCreate, len(resp.Departments))
	for i, v := range resp.Departments {
		bulk[i] = w.db.VacDepartment.Create().
			SetDepartmentID(v.ID).
			SetPartition(string(partition)).
			SetData(v)
	}
	_, err = w.db.VacDepartment.CreateBulk(bulk...).Save(ctx)
	return err
}

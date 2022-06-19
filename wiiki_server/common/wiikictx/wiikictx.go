package wiikictx

import (
	"context"
	"sync"
	"time"
	"wiiki_server/common/wiikierr"

	"xorm.io/xorm"
)

type ctxKey string

const (
	TransactionKey    ctxKey = "wiiki-transaction"
	CommonKey         ctxKey = "wiiki-common"
	ErrorPresenterKey ctxKey = "wiiki-error-presenter"
)

type Transaction struct {
	TransactionDB *xorm.Session
	ReadOnlyDB    *xorm.Session
	IsTransaction bool // true: transaction, false: not transaction
}

func WithTransaction(ctx context.Context, transaction *Transaction) context.Context {
	return context.WithValue(ctx, TransactionKey, transaction)
}

func GetTransaction(ctx context.Context) (*Transaction, error) {
	tx, ok := ctx.Value(TransactionKey).(*Transaction)
	if !ok {
		return nil, wiikierr.New(wiikierr.FailedGetTransactionFromCtx, "")
	}
	return tx, nil
}

func GetDB(ctx context.Context) (*xorm.Session, error) {

	tx, err := GetTransaction(ctx)
	if err != nil {
		return nil, err
	}

	return tx.TransactionDB, nil
}

func GetReadOnlyDB(ctx context.Context) (*xorm.Session, error) {
	tx, err := GetTransaction(ctx)
	if err != nil {
		return nil, err
	}

	if !tx.IsTransaction || tx.ReadOnlyDB == nil {
		return tx.TransactionDB, nil
	}

	return tx.TransactionDB, nil
}

type Common struct {
	TxTime      time.Time
	AccessToken string
}

func WithCommon(ctx context.Context, common *Common) context.Context {
	return context.WithValue(ctx, CommonKey, common)
}

func GetCommon(ctx context.Context) (*Common, error) {
	common, ok := ctx.Value(CommonKey).(*Common)
	if !ok {
		return nil, wiikierr.New(wiikierr.FailedGetCommonFromCtx, "")
	}
	return common, nil
}

func GetTxTime(ctx context.Context) (time.Time, error) {
	common, err := GetCommon(ctx)
	if err != nil {
		return time.Time{}, err
	}
	return common.TxTime, nil
}

// Only use middleware and handler
type ErrorPresenter struct {
	errors    []error
	errorsMut sync.Mutex
}

func WithErrorPresenter(ctx context.Context) context.Context {
	return context.WithValue(ctx, ErrorPresenterKey, &ErrorPresenter{
		errors:    []error{},
		errorsMut: sync.Mutex{},
	})
}

func GetErrorPresenter(ctx context.Context) *ErrorPresenter {
	val, ok := ctx.Value(ErrorPresenterKey).(*ErrorPresenter)
	if !ok {
		panic(wiikierr.FailedGetErrorPresenterFromCtx)
	}
	return val
}

func AddError(ctx context.Context, err error) {

	if err == nil {
		return
	}

	errPresenter := GetErrorPresenter(ctx)
	errPresenter.errorsMut.Lock()
	defer errPresenter.errorsMut.Unlock()
	errPresenter.errors = append(errPresenter.errors, err)
}

func GetErrorList(ctx context.Context) []error {
	errPresenter := GetErrorPresenter(ctx)
	errPresenter.errorsMut.Lock()
	defer errPresenter.errorsMut.Unlock()
	if len(errPresenter.errors) == 0 {
		return nil
	}
	return errPresenter.errors
}

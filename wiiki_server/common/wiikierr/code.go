package wiikierr

const (
	MigrateFailed                  = "MigrateFailed"
	JsonMarshalFailed              = "JsonMarshalFailed"
	FailedCreateXormEngine         = "FailedCreateXormEngine"
	FailedPingXormEngine           = "FailedPingXormEngine"
	FailedGetTransactionFromCtx    = "FailedGetTransactionFromCtx"
	FailedGetCommonFromCtx         = "FailedGetCommonFromCtx"
	FailedGetErrorPresenterFromCtx = "FailedGetErrorPresenterFromCtx"
	FailedFindRepository           = "FailedFindRepository"
	FailedGetRepository            = "FailedGetRepository"
	FailedInsertRepository         = "FailedInsertRepository"
	FailedUpdateRepository         = "FailedUpdateRepository"
	FailedDeleteRepository         = "FailedDeleteRepository"
	FailedUnmarshalToml            = "FailedUnmarshalToml"
	FailedMarshalToml              = "FailedMarshalToml"
	FailedOpenFile                 = "FailedOpenFile"
	FailedReadFile                 = "FailedReadFile"
	FailedBeginTransaction         = "FailedBeginTransaction"
	FailedRollbackTransaction      = "FailedRollbackTransaction"
	FailedCommitTransaction        = "FailedCommitTransaction"
	FailedNewLogger                = "FailedNewLogger"
)

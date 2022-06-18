package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"wiiki_server/common/logger"
	"wiiki_server/common/wiikictx"
	"wiiki_server/common/wiikierr"
	"wiiki_server/infra/http/middleware"

	"github.com/99designs/gqlgen/graphql"
)

type errorHandlingImpl struct {
	logger logger.WiikiLogger
}

type errInfo struct {
	ErrCodeList []string `json:"err_codes"`
	LogList     []string `json:"logs"`
	OriginErr   error    `json:"origin_error"`
}

func NewErrorHandling(logger logger.WiikiLogger) middleware.ErrorHandling {
	return &errorHandlingImpl{
		logger: logger,
	}
}

func (impl *errorHandlingImpl) ErrorHandling() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			requestBody, err := io.ReadAll(r.Body)
			if err != nil {
				impl.logger.Warn("[WIIKI_WARN]", "failed read body")
				requestBody = []byte{}
			}
			r.Body = io.NopCloser(bytes.NewBuffer(requestBody))

			// info log
			impl.LogAccess(r, requestBody)

			r = r.WithContext(graphql.WithResponseContext(r.Context(), graphql.DefaultErrorPresenter, graphql.DefaultRecover))
			r = r.WithContext(wiikictx.WithErrorPresenter(r.Context()))
			next.ServeHTTP(w, r)

			graphErrList := graphql.GetErrors(r.Context())
			middlewareErrList := wiikictx.GetErrorList(r.Context())

			if len(graphErrList) == 0 && len(middlewareErrList) == 0 {
				return
			}

			var errInfoList []*errInfo
			for _, graphErr := range graphErrList {
				innerErr := graphErr.Unwrap()
				errInfo := impl.GetErrorInfo(innerErr)
				errInfoList = append(errInfoList, errInfo)
			}
			for _, middlewareErr := range middlewareErrList {
				errInfo := impl.GetErrorInfo(middlewareErr)
				errInfoList = append(errInfoList, errInfo)
			}

			// log
			impl.LogError(errInfoList, r, requestBody)
		})
	}
}

func (impl *errorHandlingImpl) LogAccess(r *http.Request, requestBody []byte) {
	var keysAndValues []interface{}
	keysAndValues = append(keysAndValues, "request_body", string(requestBody))
	keysAndValues = append(keysAndValues, "uri", r.RequestURI)
	impl.logger.Infow("[WIIKI_INFO]", keysAndValues...)
}

func (impl *errorHandlingImpl) LogError(errInfoList []*errInfo, r *http.Request, requestBody []byte) {
	var keysAndValues []interface{}

	keysAndValues = append(keysAndValues, "uri", r.RequestURI)
	keysAndValues = append(keysAndValues, "path", r.URL.Path)
	keysAndValues = append(keysAndValues, "method", r.Method)
	keysAndValues = append(keysAndValues, "request_body", string(requestBody))

	for i, errInfo := range errInfoList {

		keysAndValues = append(keysAndValues, fmt.Sprintf("error_codes.%d", i), errInfo.ErrCodeList)
		keysAndValues = append(keysAndValues, fmt.Sprintf("logs.%d", i), errInfo.LogList)
		keysAndValues = append(keysAndValues, fmt.Sprintf("origin_error.%d", i), errInfo.OriginErr.Error())
		keysAndValues = append(keysAndValues, fmt.Sprintf("stacktrace.%d", i))
	}
	impl.logger.Errorw("[WIIKI_ERROR]", keysAndValues...)
}

func (impl *errorHandlingImpl) GetErrorInfo(err error) *errInfo {
	return impl.getErrorInfo(err, &errInfo{})
}

func (impl *errorHandlingImpl) getErrorInfo(err error, errInfo *errInfo) *errInfo {

	switch e := err.(type) {
	case *wiikierr.Error:
		errInfo.ErrCodeList = append(errInfo.ErrCodeList, e.Code())
		errInfo.LogList = append(errInfo.LogList, e.Log())
		return impl.getErrorInfo(e.Cause(), errInfo)
	case error:
		errInfo.LogList = append(errInfo.LogList, fmt.Sprintf("[ORIGIN] %s", e.Error()))
		errInfo.OriginErr = e
		return errInfo
	default:
		impl.logger.Warn("到達よ的ではない独自エラーが渡されました")
		errInfo.LogList = append(errInfo.LogList, fmt.Sprintf("[UNKNOWN] %v", e))
		errInfo.OriginErr = e
		return errInfo
	}
}

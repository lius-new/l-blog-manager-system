package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/lius-new/blog-backend/api"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/analyzerclient"
)

type AnalyzerMiddleware struct {
	analyzer analyzerclient.Analyzer
}

func NewAnalyzerMiddleware(analyzer analyzerclient.Analyzer) *AnalyzerMiddleware {
	return &AnalyzerMiddleware{
		analyzer,
	}
}

func analyzerMiddlewareError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(err.Error()))
}

func (m *AnalyzerMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var (
			BLOCKIP             = errors.New("ip blocked")
			JUDGET_BLOCK_FAILED = errors.New("judge ip blocked failed: ")
			RECORDING_FAILED    = errors.New("recording failed: ")
		)

		defer func() {
			if catchErr := recover(); catchErr != nil {
				analyzerMiddlewareError(w, catchErr.(error))
			}
		}()

		hostSplit := strings.Split(r.Host, ":")
		if len(hostSplit) == 0 {
			panic(api.ErrHostNotFound)
		}

		ip := hostSplit[0]

		// 1. 先判断是否被封禁，如果被封禁那么
		// 判断是否包含端口，如果存在就删除端口
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		judgeResp, err := m.analyzer.JudgeBlockedByIP(ctx, &analyzer.JudgeBlockedByIPRequest{
			BlockIP: ip,
		})
		if err != nil {
			panic(JUDGET_BLOCK_FAILED.Error() + err.Error())
		}
		if judgeResp.Block {
			panic(BLOCKIP)
		}

		// 2. 记录日志
		// 这个记录过程会先往数据库中添加一条记录， 然后判断记录是否需要合并（指定时间段多次访问会合并数据）, 当如果需要合并那么就意味多次访问并将当前IP添加到禁止访问的数据集合(表)中。
		_, err = m.analyzer.CreateRecord(ctx, &analyzer.CreateRecordRequest{
			RequestIp:     ip,
			RequestMethod: r.Method,
			RequestPath:   r.URL.Path,
		})
		if err != nil {
			panic(RECORDING_FAILED)
		}

		next(w, r)
	}
}

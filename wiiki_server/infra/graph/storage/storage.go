package storage

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"wiiki_server/domain/model/repomodel"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/patrickmn/go-cache"
)

type ctxKey string

const loadersKey = ctxKey("wiiki-loader")

type Loaders struct {
	UserLoader *dataloader.Loader[string, *repomodel.User]
}

func NewLoaders(userReader UserReader) *Loaders {

	// userReader := &userReaderImpl{}

	// no cache この場合、keyと同じ長さを同じ順番に渡す必要がある、結構だるい
	// noCache := &dataloader.NoCache[string, *repomodel.User]{}

	// ttl cache これもどうなのって感じだけど
	userTTLCache := &TTLCache[string, *repomodel.User]{cache.New(time.Second/2, time.Second/2)}

	return &Loaders{
		// UserLoader: dataloader.NewBatchedLoader(userReader.GetUserList, dataloader.WithCache[string, *repomodel.User](noCache)),
		// UserLoader: dataloader.NewBatchedLoader(userReader.GetUserList),
		UserLoader: dataloader.NewBatchedLoader(userReader.GetUserList, dataloader.WithCache[string, *repomodel.User](userTTLCache)),
	}
}

func Middleware(loaders *Loaders) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
			r = r.WithContext(nextCtx)
			next.ServeHTTP(w, r)
		})
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

type TTLCache[K comparable, V any] struct {
	c *cache.Cache
}

func (c *TTLCache[K, V]) Get(_ context.Context, key K) (dataloader.Thunk[V], bool) {
	k := fmt.Sprintf("%v", key) // convert the key to string because the underlying library doesn't support Generics yet
	v, ok := c.c.Get(k)
	if ok {
		return v.(dataloader.Thunk[V]), ok
	}
	return nil, ok
}

func (c *TTLCache[K, V]) Set(_ context.Context, key K, value dataloader.Thunk[V]) {
	k := fmt.Sprintf("%v", key) // convert the key to string because the underlying library doesn't support Generics yet
	c.c.Set(k, value, 0)
}

func (c *TTLCache[K, V]) Delete(_ context.Context, key K) bool {
	k := fmt.Sprintf("%v", key) // convert the key to string because the underlying library doesn't support Generics yet
	if _, found := c.c.Get(k); found {
		c.c.Delete(k)
		return true
	}
	return false
}

func (c *TTLCache[K, V]) Clear() {
	c.c.Flush()
}

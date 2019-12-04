package controller

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

type Backend struct {
	URL   *url.URL
	Alive bool
	mux   sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct{
	backend []*Backend
	current uint64
}
func (s *ServerPool) NextIndex()int{
	return int(atomic.AddUint64(&s.current,uint64(1))%uint64(len(s.backend)))
}

func (b *Backend) SetAlive(alive bool){
	b.mux.Lock()
	b.Alive=alive
	b.mux.Unlock()
}

func (b *Backend) IsAlive()(alive bool){
	b.mux.RLock()
	alive=b.Alive
	b.mux.RUnlock()
	return
}
func (s *ServerPool) GetNextPeer()*Backend{
	next:=s.NextIndex()
	l:=len(s.backend)+next
	for i:=next;i<l;i++{
		idx:=i%len(s.backend)
		if s.backend[idx].IsAlive(){
			if i!=next{
				atomic.StoreUint64(&s.current,uint64(idx))
			}
			return s.backend[idx]
		}
	}
	return nil
}

func Lb(w http.ResponseWriter,r *http.Request){
	peer:=ServerPool.GetNextPeer()
	if peer!=nil{
		peer.ReverseProxy.ServeHTTP(w,r)
		return
	}
	http.Error(w,"service not avalable",http.StatusServiceUnavailable)
}


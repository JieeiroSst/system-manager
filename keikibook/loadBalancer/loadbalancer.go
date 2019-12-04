package loadBalancer

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type Server struct {
	Route string `json:"route"`
	Alvie bool `json:"alvie"`
	ReverseProxy *httputil.ReverseProxy
}

type ServerList struct {
	Servers []Server `json:"servers"`
	Lastest int      `json:"lastest"`
}
func (server *Server) IsAlive()bool{
	timeout:=time.Duration(1*time.Second)

	log.Println("started health check for:",server.Route)
	_,err:=net.DialTimeout("tcp",server.Route,timeout)
	if err!=nil{
		log.Println(server.Route,"is dead")
		log.Println("heath check error:",err)
		server.Alvie=false

		return false
	}

	log.Println(server.Route,"is alive")
	server.Alvie=true
	return true
}

func (serverList * ServerList) init(serverRouter []string){
	log.Println("creating list for router:",serverRouter)

	for _,serverRoute:=range serverRouter{
		var localserver Server

		localserver.Route=serverRoute
		localserver.Alvie=localserver.IsAlive()

		origin,_:=url.Parse("http://"+serverRoute)
		director:=func(r *http.Request){
			r.Header.Add("x-forward-host",r.Host)
			r.Header.Add("x-origin-host",origin.Host)
			r.URL.Scheme="http"
			r.URL.Host=origin.Host
		}

		localserver.ReverseProxy=&httputil.ReverseProxy{Director:director}

		log.Println("server",localserver,"added to server list")
		serverList.Servers=append(serverList.Servers,localserver)
	}

	serverList.Lastest=-1
	log.Println("successfully created server list:",serverList)
}

func (serverList *ServerList) Nextserver()int{
	return (serverList.Lastest+1)%len(serverList.Servers)
}

func (serverList *ServerList) LoadBalance(w http.ResponseWriter,r *http.Request){
	if len(serverList.Servers)>0{
		serverCount:=0
		for index:=serverList.Nextserver();serverCount<len(serverList.Servers);index=serverList.Nextserver(){
			if serverList.Servers[index].IsAlive(){
				log.Println("routing request",r.URL,"to",serverList.Servers[index].Route)
				serverList.Servers[index].ReverseProxy.ServeHTTP(w,r)

				serverList.Lastest=index
				log.Println("update latest server to:",serverList.Lastest)

				return
			}
			serverCount++
			serverList.Lastest=serverList.Nextserver()
		}
	}
	log.Println("no server availabe")

	http.Error(w,"no server available ",http.StatusServiceUnavailable)
}
package webserver

import (
	"fmt"
	"log"
	"net/http"
)

func SetupWebServer(port int) {
	addr := fmt.Sprintf(":%d", port)
	log.Println("fred now serving at http://localhost" + addr)
	http.ListenAndServe(addr, &HotReloadFileServer{})
}

type HotReloadFileServer struct{}

func (handler *HotReloadFileServer) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		writer.Write([]byte(clientWebsocketCode))
	}

	fs := http.FileServer(http.Dir("."))
	fs.ServeHTTP(writer, req)
}

const clientWebsocketCode = `
<script>
(() => {
    const createSocket = () => {
        const socketUrl = 'ws://localhost:8090';
        let socket = new WebSocket(socketUrl);

        socket.onmessage = (e) => {
            console.log(e)
            location.reload()
        }

        socket.onclose = () => {
            setTimeout(createSocket, 1000)
        }
    }

    createSocket()
})();
</script>
`

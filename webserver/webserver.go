package webserver

import (
	"log"
	"net/http"
)

func SetupWebServer() {
	log.Println("fred now serving at http://localhost:8000")
	http.ListenAndServe(":8000", &HotReloadFileServer{})
}

type HotReloadFileServer struct{}

func (handler *HotReloadFileServer) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte(clientWebsocketCode))
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

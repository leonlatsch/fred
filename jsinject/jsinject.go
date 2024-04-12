package jsinject

import "net/http"

const js = `
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

type WsInjectorMiddleware struct {
	Next http.Handler
}

func (handler *WsInjectorMiddleware) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte(js))
	handler.Next.ServeHTTP(writer, req)
}

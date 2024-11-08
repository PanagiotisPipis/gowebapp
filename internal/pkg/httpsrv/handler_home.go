package httpsrv

import (
	"net/http"
    "github.com/gorilla/csrf"
)

func (s *Server) handlerHome(w http.ResponseWriter, r *http.Request) {
    tmplData := map[string]interface{}{
        "WsUrl":"ws://"+r.Host+"/goapp/ws",
        csrf.TemplateTag: csrf.TemplateField(r),
    }
    s.templates.ExecuteTemplate(w, "home", tmplData)
}

func homeTemplate() string {
    template := `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };
    async function GetRestricted() {
        try {
            const response = await fetch("http://localhost:8080/restricted", {
            method: "POST",
            // Set the FormData instance as the request body
            });
            console.log(await response.json());
        } catch (e) {
            console.error(e);
        }
    }
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.WsUrl}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
    document.getElementById("restrictedButton").onclick = function(evt) {
        GetRestricted()
    }
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="{}">
<button id="send">Reset</button>
</form>
<p>
<form action="/goapp/restricted" method="post">
<button id="restrictedButton">Restricted</button>
{{ .csrfField }}
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`
return template
}
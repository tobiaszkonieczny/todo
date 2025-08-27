const ws = new WebSocket("ws://localhost:8081/ws");

ws.onmessage = (event) => {
    console.log("Received:", event.data);
};
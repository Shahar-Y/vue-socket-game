const express = require("express");
const http = require("http").Server(express);
const socketio = require("socket.io")(http);

const position = {
    x: 100,
    y: 100
};

socketio.on("connection", socket => {
    socket.emit("position", position);
    socket.on("move", data => {
        switch(data) {
            case "a":
                position.x -= 5;
                // socketio.emit("position", position);
                break;
            case "d":
                position.x += 5;
                // socketio.emit("position", position);
                break;
            case "w":
                position.y -= 5;
                // socketio.emit("position", position);
                break;
            case "s":
                position.y += 5;
                break;
        }
        socketio.emit("position", position);
    });
});


var server = http.listen(3000, () => {
    console.log("Listening at port 3000");
    // setTimeout(function() {
    //     server.close();
    //     console.log("Stopped listening after timeout");
    // }, 5000);
})
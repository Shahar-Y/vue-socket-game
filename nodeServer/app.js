// const express = require("express");
// const http = require("http").Server(express);
// const socketio = require("socket.io")(http, { 'path': '/ws' });

const position = {
    X: 100,
    Y: 100
};

// list of currently connected clients (users)
var clients = [ ];

var WebSocketServer = require('websocket').server;
var http = require('http');

var server = http.createServer(function (request, response) {
    // process HTTP request. Since we're writing just WebSockets
    // server we don't have to implement anything.
});
server.listen(3000, function () { });

// create the server
wsServer = new WebSocketServer({
    httpServer: server,
    path: '/ws'
});

// WebSocket server
wsServer.on('request', function (request) {
    console.log("request!!");
    var connection = request.accept(null, request.origin);

    // This is the most important callback for us, we'll handle
    // all messages from users here.
    connection.on('message', function (message) {
        console.log("message!!");
        if (message.type === 'utf8') {
            console.log('utf8');
            // process WebSocket message
        }
        console.log(message);
        switch (message.utf8Data) {
            case "a":
                position.X -= 5;
                // socketio.emit("position", position);
                break;
            case "d":
                position.X += 5;
                // socketio.emit("position", position);
                break;
            case "w":
                position.Y -= 5;
                // socketio.emit("position", position);
                break;
            case "s":
                position.Y += 5;
                break;
        }
        // socketio.emit("position", position);
    });

    connection.on('close', function (connection) {
        console.log("close!!");
        // close user connection
    });
});




// socketio.on("connection", socket => {
//     socket.emit("position", position);
//     socket.on("move", data => {
//         switch (data) {
//             case "a":
//                 position.X -= 5;
//                 // socketio.emit("position", position);
//                 break;
//             case "d":
//                 position.X += 5;
//                 // socketio.emit("position", position);
//                 break;
//             case "w":
//                 position.Y -= 5;
//                 // socketio.emit("position", position);
//                 break;
//             case "s":
//                 position.Y += 5;
//                 break;
//         }
//         socketio.emit("position", position);
//     });
// });



// var server = http.listen(3000, () => {
//     console.log("Listening at port 3000");
//     // setTimeout(function() {
//     //     server.close();
//     //     console.log("Stopped listening after timeout");
//     // }, 5000);
// })



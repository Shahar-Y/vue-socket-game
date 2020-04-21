const express = require("express");
const http = require("http").Server(express);
const socketio = require("socket.io")(http);

http.listen(3000, () => {
    console.log("Listening at port 3000");
})
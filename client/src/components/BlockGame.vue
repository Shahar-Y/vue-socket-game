<template>
  <div @keyup="keyPress">
    <!-- <input v-on:keyup="keyPress"> -->
    <canvas ref="game" width="640" height="640" class="myCanvas"></canvas>
  </div>
</template>

<script>
// import io from "socket.io-client";

export default {
  name: "BlockGame",
  data() {
    return {
      socket: {},
      ws: {},
      context: {},
      position: {
        x: 0,
        y: 0
      }
    };
  },
  created() {
    // this.socket = new WebSocket("ws://localhost:3000");
    // this.socket = io("http://localhost:3000");
    console.log("creating socket ..." + "ws://" + window.location.host + "/ws");
    this.ws = new WebSocket("ws://localhost:3000/ws");
    this.ws.onopen = () => {
      console.log("onopen works!!");
      this.ws.send("{ message:}");
    };

    this.ws.onmessage = (msg) => {
      console.log("onmessage RECEIVED!!");
      console.log(msg);
      // this.ws.send(JSON({ message: "HELLO!!" }));
    };

    this.ws.onclose = event => {
      console.log("closing connection: ", event);
    };

    this.ws.onerror = error => {
      console.log("Socket error: ", error);
    };
    // console.log("opened socket: " + "ws://" + window.location.host + "/ws");
    this.ws.addEventListener("message", function(e) {
      console.log("received message!");
      // var msg = JSON.parse(e.data);
      console.log(e.data);
      // var element = document.getElementById("chat-messages");
      // element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
    });
  },
  mounted() {
    document.addEventListener("keydown", this.keyPress);
    this.context = this.$refs.game.getContext("2d");
    // this.socket.on("position", data => {
    //   this.position = data;
    //   this.context.clearRect(0, 0, 640, 640);
    //   this.context.fillRect(this.position.x, this.position.y, 20, 20);
    // });
  },
  methods: {
    send: function() {
      console.log("sending...");
      this.ws.send({ messages: "hello from client!" });
      this.newMsg = ""; // Reset newMsg
    },
    move(direction) {
      console.log(direction);
      // this.send();
      // this.socket.emit("move", direction);
    },
    keyPress(event) {
      console.log(event);
      this.send();
      // this.socket.emit("move", event.key);
    }
  }
};
</script>

<style scoped>
.myCanvas {
  border: 1px solid black;
}
</style>

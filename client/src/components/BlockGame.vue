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
    // this.socket = io("http://localhost:3000");
    console.log("creating socket ..." + "ws://localhost:3000/ws");
    this.ws = new WebSocket("ws://localhost:3000/ws");
    this.ws.onopen = () => {
      console.log("onopen works!!");
      this.ws.send("xyz");
    };



    this.ws.onclose = event => {
      console.log("closing connection: ", event);
    };

    this.ws.onerror = error => {
      console.log("Socket error: ", error);
    };
    // this.ws.addEventListener("message", function(e) {
    //   console.log("received message!");
    //   // console.log(e.data);
    // });
  },
  mounted() {
    document.addEventListener("keydown", this.keyPress);
    this.context = this.$refs.game.getContext("2d");
    this.ws.onmessage = msg => {
      let data = JSON.parse(msg.data)
      console.log(data);
      this.position.x = data.X;
      this.position.y = data.Y;
      this.context.clearRect(0, 0, 640, 640);
      this.context.fillRect(this.position.x, this.position.y, 20, 20);
      console.log(this.position);
      // this.ws.send(JSON({ message: "HELLO!!" }));
    };

    // this.socket.on("position", data => {
    //   this.position = data;
    //   this.context.clearRect(0, 0, 640, 640);
    //   this.context.fillRect(this.position.x, this.position.y, 20, 20);
    // });
  },
  methods: {
    send: function(key) {
      console.log("sending...");
      this.ws.send(key);
      this.newMsg = ""; // Reset newMsg
    },
    move(direction) {
      console.log(direction);
      // this.send();
      // this.socket.emit("move", direction);
    },
    keyPress(event) {
      console.log(event);
      this.send(event.key);
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

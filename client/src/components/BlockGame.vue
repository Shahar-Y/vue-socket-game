<template>
  <div @keyup="keyPress">
    <!-- <input v-on:keyup="keyPress"> -->
    <canvas ref="game" width="640" height="640" class="myCanvas"></canvas>
  </div>
</template>

<script>
import io from "socket.io-client";

export default {
  name: "BlockGame",
  data() {
    return {
      socket: {},
      context: {},
      position: {
        x: 0,
        y: 0
      }
    };
  },
  created() {
    this.socket = io("http://localhost:3000");
  },
  mounted() {
    document.addEventListener("keydown", this.keyPress);
    this.context = this.$refs.game.getContext("2d");
    this.socket.on("position", data => {
      this.position = data;
      this.context.clearRect(
        0,
        0,
        640,
        640
      );
      this.context.fillRect(this.position.x, this.position.y, 20, 20);
    });
  },
  methods: {
    move(direction) {
      this.socket.emit("move", direction);
    },
    keyPress(event) {
      this.socket.emit("move", event.key);
    }
  }
};
</script>

<style scoped>
.myCanvas {
  border: 1px solid black;
}
</style>

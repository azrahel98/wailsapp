<template>
  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png" />
  <HelloWorld />
</template>

<script lang="ts" setup>
import { onMounted } from "vue";
import { Login } from "../wailsjs/go/services/LoginService";
import HelloWorld from "./components/HelloWorld.vue";

onMounted(async () => {
  const x = await Login("adf", "afa");
  console.log(x);
});
</script>

<style lang="css">
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>

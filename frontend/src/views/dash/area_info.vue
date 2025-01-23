<template>
  <div class="mains">
    <div class="col text-start">
      <div class="page-pretitle fw-medium">Informacion del Area</div>
      <h2 class="page-title fs-4">{{ router.currentRoute.value.params.area }}</h2>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { router } from '@router/router'
import { onMounted } from 'vue'

onMounted(() => {
  console.log(router.currentRoute.value.params.area)
})
</script>

<style lang="scss" scoped>
.mains {
  display: grid;
  grid-template-rows: min-content min-content 1fr;
  row-gap: 1vh;
  grid-template-columns: 1fr;
  width: 100%;
  height: 100vh;
  overflow-y: auto;
  padding-top: 2vh;
  padding-bottom: 1vh;
}
</style>

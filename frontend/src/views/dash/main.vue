<template>
  <div class="main">
    <div class="text-start pt-2">
      <div class="page-pretitle">Overview</div>
      <h2 class="page-title">Dashboard</h2>
    </div>
    <div class="grid-items">
      <C_regimen_res />
      <c_sexo_res />
    </div>
  </div>
</template>
<script setup lang="ts">
import C_regimen_res from '@comp/dashboard/c_regimen_res.vue'
import c_sexo_res from '@comp/dashboard/c_sexo_res.vue'
import mapa from '@comp/dashboard/mapa.vue'
</script>

<style lang="scss" scoped>
.main {
  height: 100vh;
  display: grid;
  row-gap: 4px;
  grid-template-rows: min-content 1fr;
  .grid-items {
    max-height: 100%;
    display: grid;
    grid-template-columns: minmax(20vw, 40vw) 1fr;
    column-gap: 2vh;
    .card {
      height: min-content;
    }
  }
}
</style>

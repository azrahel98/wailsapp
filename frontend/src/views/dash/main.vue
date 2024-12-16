<template>
  <div class="main">
    <div class="text-start pt-2">
      <div class="page-pretitle">Overview</div>
      <h2 class="page-title">Dashboard</h2>
    </div>
    <div class="row row-deck row-cards p-0 m-0">
      <div class="col-sm-12 col-lg-6">
        <C_regimen_res />
      </div>
      <div class="col-sm-12 col-lg-6">
        <mapa></mapa>
      </div>
      <div class="col-md-4">
        <Lista class="lista" />
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import C_regimen_res from '@comp/dashboard/c_regimen_res.vue'
import Lista from '@comp/dashboard/lista.vue'
import mapa from '@comp/dashboard/mapa.vue'
</script>

<style lang="scss" scoped>
.main {
  height: 100vh;
  display: grid;
  row-gap: 4px;
  grid-template-rows: min-content 1fr;
  .lista {
    height: min-content;
    max-height: 50vh;
  }
}
</style>

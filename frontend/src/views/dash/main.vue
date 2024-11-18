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
    </div>
    <div class="cal">
      <Calendario />
    </div>
  </div>
</template>
<script setup lang="ts">
import C_regimen_res from '@comp/dashboard/c_regimen_res.vue'
import Calendario from '@comp/dashboard/calendar/calendario.vue'
import mapa from '@comp/dashboard/mapa.vue'
</script>

<style lang="scss" scoped>
.main {
  height: 100vh;
  display: grid;
  row-gap: 4px;
  grid-template-rows: 6vh auto auto;
}
</style>

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
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    grid-template-rows: min-content;
    gap: 1.5rem;

    .card {
      display: flex;
      width: max-content;
      height: min-content;
      align-items: center;
      background: #f9f9f9;
      border: 1px solid #ddd;
      border-radius: 8px;
      padding: 1rem;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }
  }
}
</style>

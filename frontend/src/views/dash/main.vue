<template>
  <div class="main">
    <div class="text-start pt-2">
      <div class="page-pretitle">Overview</div>
      <h2 class="page-title">Dashboard</h2>
    </div>
    <div class="grid-items">
      <C_regimen_res />
      <c_sexo_res />
      <Areas />
    </div>
  </div>
</template>
<script setup lang="ts">
import Areas from '@comp/dashboard/areas.vue'
import C_regimen_res from '@comp/dashboard/c_regimen_res.vue'
import c_sexo_res from '@comp/dashboard/c_sexo_res.vue'
</script>

<style lang="scss" scoped>
.main {
  height: 100vh;
  display: grid;
  max-width: 100vw;
  grid-template-rows: min-content 1fr;
  gap: 1rem;
  padding: 1rem;

  .header {
    text-align: start;
    .page-pretitle {
      font-size: 0.9rem;
      color: #6c757d;
      text-transform: uppercase;
      margin-bottom: 0.5rem;
    }

    .page-title {
      font-size: 1.5rem;
      font-weight: 600;
      color: #343a40;
    }
  }

  .grid-items {
    max-width: 100%;
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(20rem, 1fr));
    align-content: start;
    gap: 1rem;
    overflow-x: hidden;

    .card {
      display: flex;
      flex-direction: column;
      background: #ffffff;
      border: 1px solid #e9ecef;
      border-radius: 0.5rem;
      padding: 1rem;
      box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
      transition: box-shadow 0.3s;
      &:hover {
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
      }
      width: 100%;
      height: min-content;
    }
  }
}
</style>

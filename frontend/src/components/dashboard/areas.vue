<template>
  <div class="card">
    <div class="card-header p-2">
      <h3 class="card-title fs-4 text-secondary">Trabajadores por Area</h3>
    </div>
    <div class="card-table table-responsive">
      <table class="table table-vcenter">
        <thead>
          <tr>
            <th>Area</th>
            <th>Activos</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="x in area" class="fs-5">
            <td class="text-start">
              {{ x.Nombre }}
            </td>
            <td>{{ x.Cantidad }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Trabajadore_Activos_Area } from '@wails/services/DashboardService'

const area = ref(<any>[])

onMounted(async () => {
  try {
    area.value = await Trabajadore_Activos_Area()
  } catch (error) {}
})
</script>

<style lang="scss" scoped>
.card {
  display: flex !important;
  justify-content: center !important;
  align-items: start !important;
  .card-table {
    width: 100% !important;
    max-height: 40vh;
    tbody {
      overflow-y: auto;
    }
  }
}
</style>

<template>
  <div class="page">
    <div class="page-header d-print-none mb-2">
      <div class="container-xl">
        <div class="row g-2 align-items-center">
          <div class="col">
            <h2 class="page-title">Dashboard</h2>
          </div>
        </div>
      </div>
    </div>

    <div class="page-body">
      <div class="container-xl">
        <div class="row row-deck row-cards">
          <div class="col-sm-6 col-lg-4">
            <c_regimen_res :regim="regimenes" />
          </div>
          <div class="col-sm-6 col-lg-4">
            <c_sexo_res :regimes="sexos" title="Personal por Sexo" />
          </div>
          <div class="col-sm-6 col-lg-6">
            <c_areas_res :columns="columns" :rows="areas" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import c_regimen_res from '@comp/dashboard/c_regimen_res.vue'
import c_sexo_res from '@comp/dashboard/c_sexo_res.vue'
import c_areas_res from '@comp/dashboard/areas.vue'
import { onMounted, ref } from 'vue'
import {
  Regimenes_resumen,
  Resumen_Dashboard,
  Trabajadore_Activos_Area,
  Sexo_cantidad
} from '@wails/services/DashboardService'

interface Regimen {
  Cantidad: number
  Nombre: string
}

const regimenes = ref<Array<Regimen>>([])
const resumen = ref<any>()
const sexos = ref<any>([])
const areas = ref(<any>[])

onMounted(async () => {
  try {
    regimenes.value = await Regimenes_resumen()
    resumen.value = await Resumen_Dashboard()
    areas.value = await Trabajadore_Activos_Area()
    sexos.value = await Sexo_cantidad()
  } catch (error) {
    console.error('Error al cargar los datos:', error)
  }
})

const columns = [
  { field: 'Nombre', title: 'Area' },
  { field: 'Cantidad', title: 'Trabajadores' }
]
</script>

<style lang="scss" scoped>
.page {
  display: grid !important;
  grid-template-rows: min-content min-content min-content;
  height: 100vh;
}

.chart-lg {
  height: 180px;
}

.card-sm {
  margin-bottom: 0.75rem;
}

.card-sm .card-body {
  padding: 0.75rem;
}

.subheader {
  font-size: 0.625rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.04em;
  line-height: 1.6;
  color: #656d77;
}

.table td,
.table th {
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
}

.chart-sparkline {
  width: 4rem;
  height: 1.5rem;
}
</style>

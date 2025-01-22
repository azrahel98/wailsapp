<template>
  <div class="page">
    <!-- Page header -->
    <div class="page-header d-print-none mb-2">
      <div class="container-xl">
        <div class="row g-2 align-items-center">
          <div class="col">
            <h2 class="page-title">Dashboard</h2>
          </div>
        </div>
      </div>
    </div>

    <!-- Page body -->
    <div class="page-body">
      <div class="container-xl">
        <div class="row row-deck row-cards">
          <div class="col-sm-6 col-lg-3">
            <div class="card card-sm">
              <div class="card-body">
                <div class="d-flex align-items-center">
                  <div class="subheader">Total Personal Registrado</div>
                </div>
                <div class="h1 mb-3">{{ regimenes.reduce((s, i) => s + i.Cantidad, 0) }}</div>
                <div class="d-flex mb-2">
                  <div>Activos</div>
                  <div class="ms-auto">
                    <span class="text-green d-inline-flex align-items-center lh-1">
                      {{ ((resumen[0].Activos / resumen[0].Cantidad) * 100).toFixed() }} %
                    </span>
                  </div>
                </div>
                <div class="progress progress-sm">
                  <div
                    class="progress-bar bg-primary"
                    :style="{
                      width:
                        resumen[0] && resumen[0].Cantidad > 0
                          ? ((resumen[0].Activos / resumen[0].Cantidad) * 100).toFixed() + '%'
                          : '0%'
                    }"
                    role="progressbar"
                    :aria-valuenow="
                      resumen[0] && resumen[0].Cantidad > 0
                        ? ((resumen[0].Activos / resumen[0].Cantidad) * 100).toFixed()
                        : 0
                    "
                    aria-valuemin="0"
                    aria-valuemax="100"
                    aria-label="Progress"
                  >
                    <span class="visually-hidden">{{
                      ((resumen[0].Activos / resumen[0].Cantidad) * 100).toFixed()
                    }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-sm-6 col-lg-3">
            <div class="card card-sm">
              <div class="card-body">
                <div class="d-flex align-items-center">
                  <div class="subheader">Sindicatos</div>
                </div>
                <div class="d-flex align-items-baseline">
                  <div class="h1 mb-0 me-2">2</div>
                  <div class="me-auto">
                    <span class="text-yellow d-inline-flex align-items-center lh-1"> 0% </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-sm-6 col-lg-3">
            <div class="card card-sm">
              <div class="card-body">
                <div class="d-flex align-items-center">
                  <div class="subheader">Renuncias (Último mes)</div>
                </div>
                <div class="d-flex align-items-baseline">
                  <div class="h1 mb-0 me-2">{{ resumen[1].Cantidad }}</div>
                  <div class="me-auto">
                    <span class="text-red d-inline-flex align-items-center lh-1"> +3 </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-sm-6 col-lg-3">
            <div class="card card-sm">
              <div class="card-body">
                <div class="d-flex align-items-center">
                  <div class="subheader">Nuevas Contrataciones</div>
                </div>
                <div class="d-flex align-items-baseline">
                  <div class="h1 mb-0 me-2">25</div>
                  <div class="me-auto">
                    <span class="text-green d-inline-flex align-items-center lh-1"> +8 </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-sm-6 col-lg-4">
            <c_regimen_res :regim="regimenes" />
          </div>
          <div class="col-sm-6 col-lg-4">
            <c_sexo_res />
          </div>
          <div class="col-12 areas">
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
  Trabajadore_Activos_Area
} from '@wails/services/DashboardService'

interface Regimen {
  Cantidad: number
  Nombre: string
}

const regimenes = ref<Array<Regimen>>([])
const resumen = ref<any>()
const areas = ref(<any>[])

onMounted(async () => {
  try {
    regimenes.value = await Regimenes_resumen()
    resumen.value = await Resumen_Dashboard()
    areas.value = await Trabajadore_Activos_Area()
  } catch (error) {
    console.error('Error al cargar los datos:', error)
  }
})

const columns = [
  { field: 'Nombre', title: 'Area' },
  { field: 'Cantidad', title: 'Trabajadores' }
]
</script>

<style scoped>
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

.areas {
  max-height: 35vh;
}
</style>

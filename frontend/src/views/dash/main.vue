<template>
  <div class="container-xl">
    <div class="px-4 pt-2">
      <div class="max-w-7xl mx-auto">
        <div class="space-y-1">
          <span class="text-sm font-medium"> Resumen de los indicadores </span>
          <h2 class="font-bold text-gray tracking-tight">Dashboard</h2>
        </div>
      </div>
    </div>
    <div class="row row-deck row-cards">
      <div class="col-12">
        <div class="row row-cards">
          <div class="col-sm-6 col-lg-3">
            <card_info
              :title="`${resumen.personalregistrado?.Cantidad} Personas`"
              :cantidad="resumen.personalregistrado?.Activos"
              descripcion=" activos"
            >
              <span class="text-white avatar bg-primary">
                <IconUsersGroup stroke="1.1" class="icon" />
              </span>
            </card_info>
          </div>
          <div class="col-sm-6 col-lg-3">
            <card_info
              title="Renuncias del mes"
              :cantidad="resumen.renunciasmes?.[0].Cantidad"
              descripcion=" renuncias"
            >
              <span class="text-white avatar bg-warning">
                <IconUserOff stroke="1.1" class="icon" />
              </span>
            </card_info>
          </div>
          <div class="col-sm-6 col-lg-3">
            <card_info title="Puestos CAP - 2015" :cantidad="175">
              <span class="text-white avatar bg-facebook">
                <IconBriefcase2 stroke="1.1" class="icon" />
              </span>
            </card_info>
          </div>
          <div class="col-sm-6 col-lg-3">
            <card_info
              :cantidad="(resumen.regimenescantidad?.reduce((x:any, i:any) => x + i.Cantidad, 0) + 175) * 0.05"
              descripcion=" disponibles"
              :title="`${resumen.regimenescantidad?.reduce((x:any, i:any) => x + i.Cantidad, 0)} puestos CAS`"
            >
              <span class="text-white avatar bg-secondary">
                <IconCalculator stroke="1.1" class="icon" />
              </span>
            </card_info>
          </div>
        </div>
      </div>
      <div class="col-md-6 col-lg-4">
        <regimenesMedia :reg="mediaregimenes" />
      </div>
      <div class="col-md-12 col-lg-8"><cumpleañosCard :info="cumpleaños" /></div>
      <div class="col-12">
        <areasresumen :columns="columns" :rows="areas" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import card_info from '@comp/dashboard/card_info.vue'
import regimenesMedia from '@comp/dashboard/regimenes-media.vue'
import areasresumen from '@comp/dashboard/areas.vue'
import cumpleañosCard from '@comp/dashboard/cumpleaños.vue'
import {
  Resumen_Dashboard,
  Trabajadore_Activos_Area,
  Resumen_Regiemenes,
  Cumpleaños
} from '@wails/services/DashboardService'
import { IconBriefcase2, IconCalculator, IconUserOff, IconUsersGroup } from '@tabler/icons-vue'

const resumen = ref<any>([])
const areas = ref(<any>[])
const mediaregimenes = ref(<any>[])
const cumpleaños = ref<Array<any>>([])

onMounted(async () => {
  try {
    resumen.value = await Resumen_Dashboard()
    areas.value = await Trabajadore_Activos_Area()
    mediaregimenes.value = await Resumen_Regiemenes()
    cumpleaños.value = await Cumpleaños(new Date().getMonth() + 1)
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
  width: 100%;
  max-width: 100%;
  grid-template-rows: min-content min-content min-content;
  height: 100vh;
}
</style>

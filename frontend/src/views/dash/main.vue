<template>
  <div class="container-xl main-page">
    <div class="px-4 pt-2">
      <div class="max-w-7xl mx-auto">
        <div class="space-y-1">
          <span class="text-sm font-medium"> Resumen de los indicadores </span>
          <h2 class="font-bold text-gray tracking-tight">Dashboard</h2>
        </div>
      </div>
    </div>
    <div class="items">
      <div class="uno">
        <div class="row row-cards">
          <div class="col-sm-6 col-lg-3">
            <card_info
              :title="`${resumen.personalregistrado?.Cantidad} Personas`"
              :cantidad="resumen.personalregistrado?.Activos"
              :descarga="true"
              :funcion="exportar_activos"
              descripcion=" activos"
            >
              <span class="text-white avatar bg-primary">
                <IconUsersGroup stroke="1.1" class="icon" />
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
          <div class="col-sm-4 col-lg-3">
            <card_info
              :cantidad="resumen.sindicatos?.[0].Cantidad"
              descripcion="afiliados"
              title="SUTRAMUVES"
              :funcion="() => exportar_sindicato(2)"
              :descarga="true"
            >
              <span class="text-white avatar bg-success">
                <IconBrandMinecraft stroke="1.1" class="icon" />
              </span>
            </card_info>
          </div>
          <div class="col-sm-4 col-lg-3">
            <card_info
              :cantidad="resumen.sindicatos?.[1].Cantidad"
              descripcion="afiliados"
              :funcion="() => exportar_sindicato(1)"
              :descarga="true"
              title="SOMUVES"
            >
              <span class="text-white avatar bg-dribbble">
                <IconBuilding stroke="1.1" class="icon" />
              </span>
            </card_info>
          </div>
        </div>
      </div>
      <div class="dos">
        <div class="row row-cards">
          <div class="col-md-6 col-lg-4">
            <regimenesMedia :reg="mediaregimenes" />
          </div>
          <div class="col-md-12 col-lg-8"><cumpleañosCard :info="cumpleaños" /></div>
        </div>
      </div>
      <div class="row row-cards">
        <areasresumen :columns="columns" :rows="areas" />
      </div>
    </div>
    <Toast ref="toastRef" />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import Toast from '@comp/toast.vue'
import card_info from '@comp/dashboard/card_info.vue'
import regimenesMedia from '@comp/dashboard/regimenes-media.vue'
import areasresumen from '@comp/dashboard/areas.vue'
import cumpleañosCard from '@comp/dashboard/cumpleaños.vue'
import {
  Resumen_Dashboard,
  Trabajadore_Activos_Area,
  Resumen_Regiemenes,
  Cumpleaños,
  Reporte_Activos,
  Reporte_Sindicato
} from '@wails/services/DashboardService'
import { IconBrandMinecraft, IconBuilding, IconCalculator, IconUsersGroup } from '@tabler/icons-vue'
import * as XLSX from 'xlsx'

const resumen = ref<any>([])
const areas = ref(<any>[])
const mediaregimenes = ref(<any>[])
const cumpleaños = ref<Array<any>>([])

const toastRef = ref<InstanceType<typeof Toast> | null>(null)
onMounted(async () => {
  try {
    resumen.value = await Resumen_Dashboard()
    areas.value = await Trabajadore_Activos_Area()
    mediaregimenes.value = await Resumen_Regiemenes()
    cumpleaños.value = await Cumpleaños(new Date().getMonth() + 1)
  } catch (error) {
    toastRef.value?.showToast(error as string)
  }
})

const exportar_activos = async () => {
  try {
    const data = await Reporte_Activos()
    const worksheet = XLSX.utils.json_to_sheet(data)
    const workbook = XLSX.utils.book_new()
    XLSX.utils.book_append_sheet(workbook, worksheet, 'Datos')
    const excelBuffer = XLSX.write(workbook, { bookType: 'xlsx', type: 'array' })
    const blob = new Blob([excelBuffer], {
      type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
    })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = `trabajadores_activos.xlsx`
    document.body.appendChild(a)
    a.click()
    URL.revokeObjectURL(url)
    document.body.removeChild(a)
  } catch (error) {
    toastRef.value?.showToast(error as string)
  }
}

const exportar_sindicato = async (sindicato: number) => {
  try {
    const data = await Reporte_Sindicato(sindicato)
    const worksheet = XLSX.utils.json_to_sheet(data)
    const workbook = XLSX.utils.book_new()
    XLSX.utils.book_append_sheet(workbook, worksheet, 'Datos')
    const excelBuffer = XLSX.write(workbook, { bookType: 'xlsx', type: 'array' })
    const blob = new Blob([excelBuffer], {
      type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
    })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = `trabajadores_activos.xlsx`
    document.body.appendChild(a)
    a.click()
    URL.revokeObjectURL(url)
    document.body.removeChild(a)
  } catch (error) {
    toastRef.value?.showToast(error as string)
  }
}

const columns = [
  { field: 'Nombre', title: 'Area' },
  { field: 'Cantidad', title: 'Trabajadores' }
]
</script>

<style lang="scss" scoped>
.main-page {
  height: 100vh;
  display: grid;
  width: 100%;
  grid-template-rows: min-content auto;
  grid-template-columns: 1fr;
  .items {
    display: grid;
    grid-template-rows: min-content min-content auto;
    grid-template-columns: 1fr;
    row-gap: 2vh;
    max-height: 100%;
    overflow-y: auto;
    overflow-x: hidden;
    width: 100%;
    .uno,
    .dos {
      width: 100%;
    }
    .dos {
      height: 100%;
      max-height: 100%;
    }
  }
}
</style>

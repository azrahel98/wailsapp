<template>
  <div class="mains">
    <div class="col text-start mb-4">
      <div class="page-pretitle fw-medium">Información del Área</div>
      <h3 class="">{{ router.currentRoute.value.params.area }}</h3>
    </div>

    <div class="container-xl">
      <div class="row row-card" style="row-gap: 0.2vh">
        <div class="col-md-4">
          <RegimenesMedia :reg="resumenregimenes" :header="false" />
        </div>
        <div class="col-md-4">
          <div class="row row-card justify-content-between gap-2">
            <div class="col-sm-12 col-lg-12">
              <Card_info title="Trabajadores" :cantidad="trabajadores.length" descripcion="activos">
                <span class="text-white avatar bg-primary">
                  <IconUserCheck stroke="1.1" class="icon" /> </span
              ></Card_info>
            </div>
            <div class="col-sm-12 col-lg-12 text-center">
              <button class="btn btn-md btn-green" @click="exportar">
                <span class="h-50">
                  <IconFileTypeXls stroke="2" class="icon text-white icon-md" />
                </span>
                <h4 class="m-0 fw-normal small">Descargar</h4>
              </button>
            </div>
          </div>
        </div>
        <div class="col-sm-6 col-lg-4">
          <div class="card">
            <div class="card-body">
              <div class="input-icon">
                <span class="input-icon-addon">
                  <IconSearch class="icon" />
                </span>
                <input
                  v-model="searchQuery"
                  type="text"
                  class="form-control form-control-rounded"
                  placeholder="Busca..."
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="puestos">
      <div class="card card-sm" v-for="x in filteredTrabajadores" :key="x.dni">
        <div class="card-body d-flex flex-column justify-content-between text-center py-4">
          <div class="">
            <span class="avatar avatar-lg mb-3 rounded">{{
              obtenerInicialesApellidos(x.nombre)
            }}</span>
            <RouterLink
              class="text-dark mb-2 cursor-pointer pt-0 mt-0"
              :to="{ name: 'perfil', params: { dni: x.dni } }"
            >
              <h4 class="fw-bold mb-1">{{ x.nombre }}</h4>
            </RouterLink>
            <div class="text-secondary pt-0 mt-0">{{ x.area }}</div>
          </div>
          <div class="mt-3">
            <span
              class="badge text-white fs-6 py-1 align-self-start"
              :style="{
                backgroundColor:
                  resumenregimenes.find((e) => e.nombre === x.reg)?.color || '#6c757d'
              }"
            >
              {{ x.reg }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { Trabajadores_por_area } from '@wails/services/DashboardService'
import { IconUserCheck, IconSearch, IconFileTypeXls } from '@tabler/icons-vue'
import * as XLSX from 'xlsx'
import Card_info from '@comp/dashboard/card_info.vue'
import RegimenesMedia from '@comp/dashboard/regimenes-media.vue'

const regresum = () => {
  ;([...new Set(trabajadores.value.map((e) => e.reg))] as Array<any>).forEach((e) => {
    var cantidad = trabajadores.value.filter((x) => x.reg == e).length
    resumenregimenes.value.push({
      nombre: e,
      cantidad: cantidad,
      porcentaje: (100 * cantidad) / trabajadores.value.length,
      color: `#00${Math.floor(Math.random() * 256)
        .toString(16)
        .padStart(2, '0')}FF`
    })
  })
}

const router = useRouter()
const trabajadores = ref<any[]>([])
const resumenregimenes = ref<any[]>([])

const searchQuery = ref('')

const filteredTrabajadores = computed(() => {
  const query = searchQuery.value.toLowerCase()
  return trabajadores.value.filter(
    (trabajador: any) =>
      trabajador.nombre.toLowerCase().includes(query) ||
      trabajador.area.toLowerCase().includes(query) ||
      trabajador.reg.toLowerCase().includes(query)
  )
})

onMounted(async () => {
  try {
    trabajadores.value = await Trabajadores_por_area(
      router.currentRoute.value.params.area.toString()
    )
    regresum()
  } catch (error) {
    console.error(error)
  }
})

watch(searchQuery, () => {})

const obtenerInicialesApellidos = (nombreCompleto: string) => {
  const palabras = nombreCompleto.trim().split(/\s+/)
  if (palabras.length < 2) return ''
  const apellido1 = palabras[palabras.length - 2][0] || ''
  const apellido2 = palabras[palabras.length - 1][0] || ''

  return apellido1.toUpperCase() + apellido2.toUpperCase()
}

const exportar = () => {
  const worksheet = XLSX.utils.json_to_sheet(trabajadores.value)
  const workbook = XLSX.utils.book_new()
  XLSX.utils.book_append_sheet(workbook, worksheet, 'Datos')
  const excelBuffer = XLSX.write(workbook, { bookType: 'xlsx', type: 'array' })
  const blob = new Blob([excelBuffer], {
    type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
  })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `${router.currentRoute.value.params.area.toString()}.xlsx`
  document.body.appendChild(a)
  a.click()
  URL.revokeObjectURL(url)
  document.body.removeChild(a)
}
</script>

<style lang="scss" scoped>
.mains {
  display: grid;
  height: 100vh;
  grid-template-rows: min-content max-content;
  gap: 1rem;
  padding: 1rem;

  .puestos {
    align-self: flex-start;
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
    row-gap: 2vh;
    column-gap: 1vh;
    height: 100%;
    align-content: start;
    justify-content: center;
    align-items: center;
    justify-items: center;
    overflow-y: auto;
    .card {
      max-width: 200px;
      min-height: 200px;
      max-height: 220px;
    }
  }
}
</style>

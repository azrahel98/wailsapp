<template>
  <div class="mains">
    <div class="col text-start mb-4">
      <div class="page-pretitle fw-medium">Información del Área</div>
      <h2 class="page-title fs-4">{{ router.currentRoute.value.params.area }}</h2>
    </div>

    <div class="graficos pb-3">
      <div class="row">
        <div class="col-md-4 col-sm-6">
          <div class="chart">
            <c_regimen_res :regim="regCount" :responsive="true" />
          </div>
        </div>
        <div class="col-md-8 col-sm-6">
          <div class="info">
            <div class="card card-sm mb-3">
              <div class="card-body py-1 px-2">
                <div class="row align-items-center">
                  <div class="col-auto">
                    <span class="bg-primary text-white avatar avatar-sm">
                      <IconUserBolt />
                    </span>
                  </div>
                  <div class="col">
                    <div class="font-weight-medium">{{ trabajadores.length }}</div>
                    <div class="text-secondary fs-5">Trabajadores</div>
                  </div>
                </div>
              </div>
            </div>

            <div class="card card-sm mb-3">
              <div class="card-body py-1 px-2">
                <div class="row align-items-center">
                  <div class="col-auto">
                    <span class="bg-instagram text-white avatar">
                      <IconUserBolt class="icon" />
                    </span>
                  </div>
                  <div class="col">
                    <div class="font-weight-medium">
                      {{ trabajadores.filter((e: any) => e.reg === 'CAS-D').length }}
                    </div>
                    <div class="text-secondary fs-5">Transitorios</div>
                  </div>
                </div>
              </div>
            </div>

            <div class="card card-sm mb-3">
              <div class="card-body py-1 px-2">
                <div class="row align-items-center">
                  <div class="col-auto">
                    <span class="bg-primary text-white avatar">
                      <IconUserBolt class="icon" />
                    </span>
                  </div>
                  <div class="col">
                    <div class="font-weight-medium">
                      {{
                        trabajadores.length -
                        trabajadores.filter((e: any) => e.reg === 'CAS-D').length
                      }}
                    </div>
                    <div class="text-secondary fs-5">Indeterminado</div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="acciones d-flex justify-content-between mt-3">
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

            <button class="btn btn-green rounded-5 btn-icon" @click="exportar">
              <IconDownload />
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="puestos">
      <div class="card card-link p-0 m-0 mb-3" v-for="x in filteredTrabajadores" :key="x.dni">
        <div class="card-body p-3">
          <RouterLink class="text-dark" :to="{ name: 'perfil', params: { dni: x.dni } }">
            <h5 class="m-0 mb-2 fw-bold">{{ x.nombre }}</h5>
          </RouterLink>
          <div class="text-muted small mb-2">{{ x.dni }}</div>
          <div class="text-muted mb-2 fs-5">{{ x.area }}</div>
          <span class="badge bg-primary text-white fs-6 py-1">
            {{ x.reg }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { Trabajadores_por_area } from '@wails/services/DashboardService'
import { IconDownload, IconSearch, IconUserBolt } from '@tabler/icons-vue'
import * as XLSX from 'xlsx'
import c_regimen_res from '@comp/dashboard/c_regimen_res.vue'

const router = useRouter()
const trabajadores = ref<any[]>([])
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
  } catch (error) {
    console.error(error)
  }
})

watch(searchQuery, () => {})

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

interface Regimen {
  Cantidad: number
  Nombre: string
}
const regCount = computed<Regimen[]>(() => {
  const counts = trabajadores.value.reduce((acc: Record<string, number>, obj: any) => {
    acc[obj.reg] = (acc[obj.reg] || 0) + 1
    return acc
  }, {})
  return Object.entries(counts).map(([Nombre, Cantidad]) => ({
    Nombre,
    Cantidad: Number(Cantidad)
  }))
})
</script>

<style lang="scss" scoped>
.mains {
  display: grid;
  height: 100vh;
  grid-template-rows: min-content auto;
  gap: 1rem;
  padding: 1rem;

  .graficos {
    .chart {
      max-width: 100%;
    }

    .info {
      display: flex;
      flex-wrap: wrap;
      justify-content: end;
      gap: 1rem;

      .card {
        width: 18vw;
        height: max-content;
      }
    }

    .acciones {
      gap: 1rem;
    }
  }

  .puestos {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    gap: 1rem;
    height: 100%;
    overflow-y: auto;
    background-color: red;

    .card {
      width: 100%;
    }
  }
}
</style>

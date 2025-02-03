<template>
  <div class="mains">
    <div class="col text-start">
      <div class="page-pretitle fw-medium">Informacion del Area</div>
      <h2 class="page-title fs-4">{{ router.currentRoute.value.params.area }}</h2>
    </div>
    <div class="graficos pb-3">
      <div class="chart">
        <c_regimen_res :regim="regCount" :responsive="true" />
      </div>
      <div class="controles">
        <div class="info">
          <div class="card card-sm">
            <div class="card-body">
              <div class="row align-items-center">
                <div class="col-auto">
                  <span class="bg-primary text-white avatar">
                    <IconUserBolt class="icon" />
                  </span>
                </div>
                <div class="col">
                  <div class="font-weight-medium">{{ trabajadores.length }}</div>
                  <div class="text-secondary">Trabajadores</div>
                </div>
                <div class="col-auto"></div>
              </div>
            </div>
          </div>

          <div class="card card-sm">
            <div class="card-body">
              <div class="row align-items-center">
                <div class="col-auto">
                  <span class="bg-instagram text-white avatar">
                    <IconUserBolt class="icon" />
                  </span>
                </div>
                <div class="col">
                  <div class="font-weight-medium">
                    {{ trabajadores.filter((e: any) => e.reg == 'CAS-D').length }}
                  </div>
                  <div class="text-secondary">Transitorios</div>
                </div>
                <div class="col-auto"></div>
              </div>
            </div>
          </div>
          <div class="card card-sm">
            <div class="card-body">
              <div class="row align-items-center">
                <div class="col-auto">
                  <span class="bg-primary text-white avatar">
                    <IconUserBolt class="icon" />
                  </span>
                </div>
                <div class="col">
                  <div class="font-weight-medium">
                    {{
                      trabajadores.length - trabajadores.filter((e: any) => e.reg == 'CAS-D').length
                    }}
                  </div>
                  <div class="text-secondary">Indeterminado</div>
                </div>
                <div class="col-auto"></div>
              </div>
            </div>
          </div>
        </div>
        <div class="acciones">
          <div class="d-flex">
            <div class="ms-auto">
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
          <button class="btn btn-green rounded-5 btn-icon" @click="exportar">
            <IconDownload />
          </button>
        </div>
      </div>
    </div>

    <div class="puestos">
      <div class="card card-link p-0 m-0" v-for="x in filteredTrabajadores" :key="x.dni">
        <div class="card-body p-3">
          <RouterLink class="text-dark" :to="{ name: 'perfil', params: { dni: x.dni } }">
            <h5 class="m-0 mb-2 fw-bold">{{ x.nombre }}</h5>
          </RouterLink>

          <div class="text-muted small mb-2">
            {{ x.dni }}
          </div>

          <div class="text-muted mb-2 fs-5">
            {{ x.area }}
          </div>
          <span class="badge bg-primary text-white fs-6 py-1 p">
            {{ x.reg }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { router } from '@router/router'
import c_regimen_res from '@comp/area/regimenes.vue'
import { computed, onMounted, ref, watch } from 'vue'
import { Trabajadores_por_area } from '@wails/services/DashboardService'
import { IconDownload, IconSearch, IconUserBolt } from '@tabler/icons-vue'
import * as XLSX from 'xlsx'

const trabajadores = ref(<any>[])
const searchQuery = ref('')
const filteredTrabajadores = ref(<any>[])

onMounted(async () => {
  try {
    trabajadores.value = await Trabajadores_por_area(
      router.currentRoute.value.params.area.toString()
    )
    filteredTrabajadores.value = trabajadores.value
  } catch (error) {
    console.log(error)
  }
})

watch(searchQuery, (newQuery) => {
  const query = newQuery.toLowerCase()
  filteredTrabajadores.value = trabajadores.value.filter(
    (trabajador: any) =>
      trabajador.nombre.toLowerCase().includes(query) ||
      trabajador.area.toLowerCase().includes(query) ||
      trabajador.reg.toLowerCase().includes(query)
  )
})

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
  grid-template-rows: min-content min-content auto;
  grid-template-columns: 1fr;
  row-gap: 1vh;
  width: 100%;
  height: 100vh;
  padding-top: 2vh;
  padding-bottom: 1vh;

  .graficos {
    width: 100%;
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
    grid-template-rows: auto;
    justify-items: center;

    .chart {
      max-width: 280px;
      height: min-content;
    }
    .controles {
      display: grid;
      grid-template-rows: auto auto;
      justify-items: center;
      .info {
        width: 100%;
        display: flex;
        flex-wrap: wrap;
        justify-content: end;
        row-gap: 2px;
        column-gap: 5px;
        .card {
          max-width: 25ch;
          height: max-content;
          .card-body {
            padding: 5px 0px 5px 5px;
          }
        }
      }
      .acciones {
        display: flex;
        justify-content: space-around;
        height: min-content;
        width: 100%;
      }
    }
  }

  .puestos {
    height: 100%;
    overflow-y: auto;
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
    row-gap: 2px;
    column-gap: 4px;
    justify-content: start;
    align-content: start;
    justify-items: center;
    width: 100%;

    .card {
      width: 100%;
      max-width: 30vw;
      height: min-content;
      .card-body {
        display: grid;
        grid-template-rows: auto auto auto auto;
        justify-items: center;
        span {
          height: min-content;
        }
        .mt-1 {
          height: min-content;
          width: 100%;

          span {
            font-size: 0.7rem;
            line-break: break-all;
            height: min-content;
            white-space: wrap;
            line-break: loose;
            line-height: normal;
          }
        }
      }
    }
  }
}
</style>

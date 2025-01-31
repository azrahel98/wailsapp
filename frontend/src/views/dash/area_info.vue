<template>
  <div class="mains">
    <div class="col text-start">
      <div class="page-pretitle fw-medium">Informacion del Area</div>
      <h2 class="page-title fs-4">{{ router.currentRoute.value.params.area }}</h2>
    </div>
    <div class="chart">
      <c_regimen_res :regim="regCount" />
    </div>
    <div class="puestos">
      <div class="d-flex">
        <div class="ms-auto">
          <div class="input-icon">
            <span class="input-icon-addon">
              <IconSearch class="icon" />
            </span>
            <input
              v-model="searchQuery"
              type="text"
              class="form-control form-control-sm form-control-rounded"
              placeholder="Busca..."
            />
          </div>
        </div>
      </div>
      <div class="card card-link p-0 m-0" v-for="x in filteredTrabajadores" :key="x.dni">
        <div class="card-body">
          <RouterLink class="text-black" :to="{ name: 'perfil', params: { dni: x.dni } }">
            <h5 class="m-0 mb-1">{{ x.nombre }}</h5>
          </RouterLink>
          <div class="text-secondary fs-5">{{ x.dni }}</div>
          <div class="mt-1">
            <span class="text-secondary fw-bold fs-5">{{ x.area }} </span>
          </div>
          <span class="badge fs-6 bg-success text-white fs-6">{{ x.reg }} </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { router } from '@router/router'
import c_regimen_res from '@comp/dashboard/c_regimen_res.vue'
import { computed, onMounted, ref, watch } from 'vue'
import { Trabajadores_por_area } from '@wails/services/DashboardService'
import { IconSearch } from '@tabler/icons-vue'

const trabajadores = ref(<any>[])
const searchQuery = ref('')
const filteredTrabajadores = ref(<any>[])

onMounted(async () => {
  try {
    trabajadores.value = await Trabajadores_por_area(
      router.currentRoute.value.params.area.toString()
    )
    filteredTrabajadores.value = trabajadores.value // Inicializar lista filtrada
    console.log(trabajadores.value)
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
  grid-template-rows: min-content min-content 1fr;
  grid-template-columns: 1fr;
  row-gap: 1vh;
  width: 100%;
  height: 100vh;
  padding-top: 2vh;
  padding-bottom: 1vh;

  .chart {
    justify-self: center;
    max-width: max-content;
  }

  .puestos {
    height: 100%;
    overflow-y: auto;
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    row-gap: 2px;
    column-gap: 4px;
    justify-content: start;
    align-content: start;
    width: 100%;

    .card {
      height: 100%;
      width: 100%;
    }
  }
}
</style>

<template>
  <div class="mains">
    <div class="px-4">
      <div class="max-w-7xl mx-auto">
        <div class="space-y-1">
          <span class="text-sm font-medium"> Buscar Trabajadores </span>
          <h2 class="font-bold text-gray tracking-tight">Buscar</h2>
        </div>
      </div>
    </div>
    <div class="search">
      <div class="input-icon w-33">
        <span class="input-icon-addon">
          <IconSearch class="icon" />
        </span>
        <input
          type="text"
          v-model="busqueda"
          @keyup.enter="realizarBusqueda"
          class="form-control text-center"
          placeholder="Search…"
          aria-label="Search in website"
        />
      </div>
    </div>
    <div class="resultados">
      <div class="card placeholder-glow" v-for="_ in 5" v-if="loading">
        <div class="ratio ratio-21x9 card-img-top placeholder"></div>
        <div class="card-body">
          <div class="placeholder col-9 mb-3"></div>
          <div class="placeholder placeholder-xs col-10"></div>
          <div class="placeholder placeholder-xs col-11"></div>
        </div>
      </div>

      <div v-else class="card" v-for="x in trabajadores">
        <div class="card-body pt-2 text-center">
          <span class="avatar avatar-xl mb-3 rounded">
            <img :src="`${x.Foto.String}`" class="border-1 border-secondary" v-if="x.Foto.Valid" />
            <img
              src="../../assets/images/man.svg"
              class="border-1 border-secondary"
              v-else-if="x.Sexo == 'M'" />
            <img src="../../assets/images/mujer.svg" class="border-1 border-secondary" v-else
          /></span>

          <RouterLink
            class="text-black bg-azure"
            :to="{ name: 'perfil', params: { dni: x.Dni.toString() } }"
            @click="console.log(`Navigating to /perfil/${x.Dni}`)"
          >
            <h4 class="m-0 mb-1">{{ x.Nombre }}</h4>
          </RouterLink>
          <div class="text-secondary">{{ x.Dni }}</div>
          <div class="mt-1">
            <span
              class="badge fs-5"
              :class="[x.Estado == 'activo' ? 'bg-primary text-white' : 'bg-secondary text-white']"
              >{{ x.Estado == 'activo' ? 'Activo' : 'Inactivo' }}</span
            >
          </div>
        </div>
      </div>
    </div>
  </div>
  <Toast ref="toastRef" />
</template>
<script setup lang="ts">
import Toast from '@comp/toast.vue'
import { IconSearch } from '@tabler/icons-vue'
import { Buscar_trabajador } from '@wails/services/DashboardService'
import { ref } from 'vue'

const toastRef = ref<InstanceType<typeof Toast> | null>(null)
const trabajadores = ref<Array<any>>([])
const busqueda = ref('')
const loading = ref(false)

const realizarBusqueda = async () => {
  try {
    trabajadores.value = []
    if (busqueda.value.length > 3) {
      loading.value = true
      const res: any = await Buscar_trabajador(busqueda.value)
      trabajadores.value = res
      if (trabajadores.value.length < 1) {
        throw 'No hay trabajadores'
      }
    }
  } catch (error) {
    toastRef.value?.showToast('No se encontraron trabajadores')
    trabajadores.value = []
  } finally {
    loading.value = false
  }
}
</script>

<style lang="scss" scoped>
.mains {
  display: grid;
  grid-template-rows: min-content min-content 1fr;
  row-gap: 1vh;
  grid-template-columns: 1fr;
  width: 100%;
  height: 100vh;
  overflow-y: auto;
  padding-top: 2vh;
  padding-bottom: 1vh;

  .info {
    color: #707eae;
    padding-top: 2vh;
    font-family: 'DM Sans';
    font-size: 0.74;
  }
  .search {
    justify-self: center;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
  }

  .resultados {
    overflow-y: scroll;
    padding-top: 3vh;
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    grid-auto-flow: row;
    column-gap: 2vh;
    row-gap: 3vh;
    justify-content: center;
    justify-items: center;
    height: 100%;
    width: 100%;
    .card {
      height: min-content;
      width: 200px;
      max-height: min-content;
    }
  }
}
</style>

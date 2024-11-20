<template>
  <div class="mains">
    <div class="col text-start">
      <div class="page-pretitle fw-medium">Buscar por nombres</div>
      <h2 class="page-title">Buscar</h2>
    </div>
    <div class="search">
      <input
        type="text"
        v-model="busqueda"
        @keyup.enter="realizarBusqueda"
        class="form-control rounded-2 text-center fs-6 w-100"
      />
    </div>
    <div class="resultados">
      <div class="card" v-for="x in trabajadores">
        <div
          class="card-status-top"
          :class="[x.Estado == 'activo' ? 'bg-success' : 'bg-danger']"
        ></div>
        <div class="card-body p-4 text-center">
          <span class="avatar avatar-xl mb-3 rounded">
            <img :src="`${x.Foto.String}`" class="border-1 border-secondary" v-if="x.Foto.Valid" />
            <img src="" class="border-1 border-secondary" v-else-if="!x.Sexo" />
            <img src="" class="border-1 border-secondary" v-else
          /></span>

          <RouterLink
            class="text-black"
            :to="{ name:'perfil',params:{dni:x.Dni.toString()} }"
            @click="console.log(`Navigating to /perfil/${x.Dni}`)"
          >
            <h4 class="m-0 mb-1">{{ x.Nombre }}</h4>
          </RouterLink>
          <div class="text-secondary">{{ x.Dni }}</div>
          <div class="mt-1">
            <span
              class="badge"
              :class="[x.Estado == 'activo' ? 'bg-success text-white' : 'bg-danger text-white']"
              >{{ x.Estado ? 'Activo' : 'Inactivo' }}</span
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { Buscar_trabajador } from '@wails/services/DashboardService'
import { ref } from 'vue'

const trabajadores = ref<Array<any>>([])

const busqueda = ref('')

const realizarBusqueda = async () => {
  try {
    trabajadores.value = []
    if (busqueda.value.length > 3) {
      //   const res: any = await invoke('buscar_trabajadores', { nombre: busqueda.value })
      //   console.log(res)
      const res: any = await Buscar_trabajador(busqueda.value)
      trabajadores.value = res
      console.log(res)
    }
  } catch (error) {
    trabajadores.value = []
    console.log(error)
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
  padding-top: 3vh;
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
    width: 30vw;
    align-items: center;
    justify-content: center;
    .form-control {
      justify-self: center;
      font-size: 0.84rem !important;
      max-width: 25vw;
    }
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
      max-height: 40vh;
    }
  }
}
</style>

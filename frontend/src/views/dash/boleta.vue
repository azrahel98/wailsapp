<template>
  <div class="organigrama">
    <div class="px-4">
      <div class="max-w-7xl mx-auto">
        <div class="space-y-1">
          <span class="text-sm font-medium"> Organigrama </span>
          <h2 class="font-bold text-gray tracking-tight">Organos / Unidades</h2>
        </div>
      </div>
    </div>
    <div class="container">
      <div class="row row-cards">
        <div class="col-12">
          <div class="card">
            <div class="card-header">
              <h3 class="card-title">{{ gerenciaPrincipal?.Area }}</h3>
            </div>
            <div class="card-body">
              <div class="row g-3">
                <div
                  class="col-md-6 col-sm-12"
                  v-for="x in gerenciaPrincipal?.Subgerencias"
                  :key="x.Id"
                >
                  <div
                    class="row g-3 align-items-center cursor-pointer"
                    :class="selected === x.Id ? 'text-primary fw-bold' : ''"
                    @click="selected = x.Id"
                  >
                    <a class="col-auto">
                      <span class="avatar avatar-2">
                        <span class="badge" :class="x.Jefe ? 'bg-success' : 'bg-red'" />
                        {{ x.Subgerencias?.length || 0 }}
                      </span>
                    </a>
                    <div class="col text-truncate">
                      <a class="text-reset d-block text-truncate">{{ x.Jefe }}</a>
                      <div class="text-secondary text-truncate mt-n1">{{ x.Area }}</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="col-10 offset-1" v-if="selected !== 0">
          <div class="card">
            <div class="card-header d-flex justify-content-center flex-column">
              <RouterLink
                v-if="subgerenciaSeleccionada?.Dni"
                class="text-reset d-block text-truncate"
                :to="{ name: 'perfil', params: { dni: subgerenciaSeleccionada.Dni.toString() } }"
                ><span class="px-0 py-0 m-0">{{ subgerenciaSeleccionada?.Jefe }}</span>
              </RouterLink>

              <h4 class="px-0 py-0 m-0">{{ subgerenciaSeleccionada?.Area }}</h4>
            </div>
            <div class="card-body">
              <div class="row g-3">
                <div
                  class="col-md-6"
                  v-for="x in subgerenciaSeleccionada?.Subgerencias"
                  :key="x.Id"
                >
                  <div class="row g-3 align-items-center">
                    <a class="col-2">
                      <span class="avatar avatar-2">
                        <span class="badge" :class="x.Jefe ? 'bg-success' : 'bg-red'" />
                        {{ x.Subgerencias?.length || 0 }}
                      </span>
                    </a>
                    <div class="col-10 text-truncate">
                      <RouterLink
                        v-if="x.Dni"
                        class="text-reset d-block text-truncate"
                        :to="{ name: 'perfil', params: { dni: x.Dni.toString() } }"
                      >
                        {{ x.Jefe }}
                      </RouterLink>
                      <a v-else class="text-reset d-block text-truncate">{{ x.Jefe }}</a>
                      <div class="text-secondary text-truncate mt-n1">{{ x.Area }}</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Reporte_Organigrama } from '@wails/services/DashboardService'

const gerencias = ref([])
const selected = ref(0)

onMounted(async () => {
  try {
    gerencias.value = await Reporte_Organigrama()
  } catch (error) {
    console.error(error)
  }
})

const gerenciaPrincipal = computed(() => gerencias.value.find((e) => e.Subgerencias.length > 0))
const subgerenciaSeleccionada = computed(() =>
  gerenciaPrincipal.value?.Subgerencias.find((e) => e.Id === selected.value)
)
</script>

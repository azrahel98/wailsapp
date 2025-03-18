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
              <h3 class="card-title">
                {{ gerencias.filter((e) => e.Subgerencias != [])[0]?.Area }}
              </h3>
            </div>
            <div class="card-body">
              <div class="row g-3">
                <div
                  class="col-md-6 col-sm-12"
                  v-for="x in gerencias.filter((e) => e.Subgerencias != [])[0]?.Subgerencias"
                >
                  <div
                    class="row g-3 align-items-center cursor-pointer"
                    :class="selected == x.Id ? 'text-primary fw-bold' : ''"
                    @click="() => (selected = x.Id)"
                  >
                    <a class="col-auto">
                      <span class="avatar avatar-2">
                        <span class="badge bg-success" v-if="x.Jefe" />
                        <span class="badge bg-red" v-else />
                        {{ x.Subgerencias?.length }}
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
        <div class="col-10 offset-1" v-if="selected != 0">
          <div class="card">
            <div
              class="card-header d-flex flex-column justify-content-center align-items-center text-center px-0 py-2"
            >
              <h4 class="fw-bold px-0 py-0 m-0">
                {{
                  gerencias
                    .filter((e) => e.Subgerencias != [])[0]
                    ?.Subgerencias?.filter((e) => e.Id == selected)[0]?.Area
                }}
              </h4>
              <span class="px-0 py-0 m-0">{{
                gerencias
                  .filter((e) => e.Subgerencias != [])[0]
                  ?.Subgerencias?.filter((e) => e.Id == selected)[0]?.Jefe
              }}</span>
            </div>
            <div class="card-body">
              <div class="row g-3">
                <div
                  class="col-md-6"
                  v-for="x in gerencias
                    .filter((e) => e.Subgerencias != [])[0]
                    ?.Subgerencias?.filter((e) => e.Id == selected)[0]?.Subgerencias"
                >
                  <div class="row g-3 align-items-center">
                    <a class="col-2">
                      <span class="avatar avatar-2">
                        <span class="badge bg-success" v-if="x.Jefe" />
                        <span class="badge bg-red" v-else />
                        {{ x.Subgerencias?.length }}
                      </span>
                    </a>
                    <div class="col-10 text-truncate">
                      <a class="text-reset d-block text-truncate">{{ x.Jefe }}</a>
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
import { ref, onMounted } from 'vue'
import { Reporte_Organigrama } from '@wails/services/DashboardService'

const gerencias = ref([])

const selected = ref(0)

onMounted(async () => {
  try {
    const data = await Reporte_Organigrama()
    gerencias.value = data
    console.log(data)
  } catch (error) {
    console.log(error)
  }
})
</script>

<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Cumpleaños del Mes</h3>
    </div>
    <div class="list-group list-group-flush overflow-auto" style="max-height: 35rem">
      <div class="list-group-item" v-for="x in cumples">
        <div class="row">
          <div class="col-auto">
            <a>
              <span class="avatar"
                >{{ (x.Nombre as string).split(' ')[0][0]
                }}{{ (x.Nombre as string).split(' ')[1][0] }}</span
              >
            </a>
          </div>
          <div class="col text-truncate">
            <a class="text-body text-start d-block">{{ x.Nombre }}</a>
            <div class="text-secondary text-truncate mt-n1">
              {{ format(parseISO(x.Nacimiento), 'dd-MM-yyyy') }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Cumpleaños } from '@wails/services/DashboardService'
import { onMounted, ref } from 'vue'
import { getYear, format, parseISO } from 'date-fns'

const cumples = ref<Array<any>>([])
const ahora = ref<Date>(new Date())

onMounted(async () => {
  await consulta()
})

const consulta = async () => {
  try {
    cumples.value = []

    cumples.value = await Cumpleaños(ahora.value.getMonth() + 1)
  } catch (error) {
    console.log(error)
  }
}
</script>

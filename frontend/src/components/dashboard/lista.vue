<template>
  <div class="d-flex align-items-center justify-content-center">
    <div class="card shadow" style="width: 320px">
      <div class="card-body p-4">
        <h5 class="card-title mb-4">
          {{ cumples.filter((e) => getDay(formatISO(e.Nacimiento)) == ahora.getDay()).length }}
          Cumpleaños Hoy
        </h5>

        <button class="btn btn-primary w-100 mt-3">Revisar Todos</button>
      </div>
      <div class="card-footer">
        <div class="row align-items-center">
          <div class="col-auto ms-auto">
            <div class="avatar-list">
              <span class="avatar avatar-sm rounded" v-for="nomb in cumples.slice(0, 5)"
                >{{ (nomb.Nombre as string).split(' ')[0][0]
                }}{{ (nomb.Nombre as string).split(' ')[1][0] }}</span
              >

              <span class="avatar avatar-sm rounded">+{{ cumples.length - 5 }}</span>
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
import { format, formatISO, getDay, parseISO } from 'date-fns'

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

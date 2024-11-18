<template>
  <div class="calendario card">
    <div class="header">
      <div class="botones">
        <button
          class="btn btn-sm p-2 btn-ghost-facebook"
          @click="() => (ahora = subMonths(ahora, 1))"
        >
          <IconArrowLeft size="17" />
        </button>
        <button class="btn btn-sm p-2 btn-outline" @click="() => (ahora = new Date())">Hoy</button>
        <button
          class="btn btn-sm p-2 btn-ghost-facebook"
          @click="() => (ahora = addMonths(ahora, 1))"
        >
          <IconArrowRight size="17" />
        </button>
      </div>
      <span class="page-title fs-1">{{ meses[ahora.getMonth()] }} {{ ahora.getFullYear() }}</span>
    </div>
    <div class="semana text-secondary">
      <span>Lun</span>
      <span>Mar</span>
      <span>Mie</span>
      <span>Jue</span>
      <span>Vie</span>
      <span>Sab</span>
      <span>Dom</span>
    </div>
    <div class="cuerpo">
      <div v-for="_x in semana()" class="card bg-white" />
      <div class="card dia fs-4 fw-bold hoydia" v-for="x in getDaysInMonth(ahora)">
        <h4 class="m-0 p-0 py-1">{{ x }}</h4>
        <div class="d-flex flex-wrap gap-1 justify-content-center">
          <span
            class="w-100"
            v-for="d in cumples.filter((e: any) => getDate(parseISO(e.Nacimiento)) == x)"
          >
            <span class="badge text-bg-flickr fs-6">{{ d.Nombre }}</span>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { IconArrowLeft, IconArrowRight } from '@tabler/icons-vue'
import { Cumpleaños } from '@wails/services/DashboardService'
import { addMonths, getDate, getDay, getDaysInMonth, parseISO, subMonths } from 'date-fns'
import { onMounted, ref, watch } from 'vue'
const meses = [
  'Enero',
  'Febrero',
  'Marzo',
  'Abril',
  'Mayo',
  'Junio',
  'Julio',
  'Agosto',
  'Septiembre',
  'Octubre',
  'Noviembre',
  'Diciembre'
]
const ahora = ref<Date>(new Date())
const cumples = ref<Array<any>>([])

onMounted(async () => {
  await consulta()
})

watch(ahora, async (_x, _y) => {
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

const semana = () => {
  let dia = getDay(new Date(ahora.value.getFullYear(), ahora.value.getMonth(), 1)) - 1
  if (dia < 0) return 6
  return dia
}
</script>

<style lang="scss" scoped>
.calendario {
  display: grid;
  grid-template-rows: min-content min-content auto;
  padding: 0.4vh 1vw 1vh 1vw;
  align-items: start;
  height: 100%;
  overflow: hidden;
  .semana,
  .cuerpo {
    display: grid;
    row-gap: 0.25vh;
    align-items: start;
    align-content: start;
    vertical-align: middle;
    column-gap: 0.2vw;
    text-align: center;
    grid-template-columns: repeat(7, 1fr);
  }

  .semana {
    padding-top: 1.3vh;
    height: min-content;
  }

  .cuerpo {
    height: 100%;
    grid-auto-rows: min-content;

    .card {
      display: flex;
      flex-direction: column;
      height: 100%;
      padding: 0.5em;
      border: 1px solid #ddd;
      border-radius: 10px;
      background-color: #f9f9f9;
      overflow-y: auto;
      max-height: 15vh;
      span {
        height: min-content;
        min-height: 2vh;
        white-space: wrap;
        width: 100%;
      }
    }
  }
}
</style>

<template>
  <div
    class="modal fade"
    id="searchcalendar"
    data-bs-backdrop="static"
    data-bs-keyboard="false"
    tabindex="-1"
    aria-labelledby="staticBackdropLabel"
    aria-hidden="true"
  >
    <div class="modal-dialog modal-xl">
      <div class="modal-content">
        <div class="modal-header d-flex">
          <div class="d-flex justify-content-around align-items-center pt-3 mb-3 w-100">
            <button class="btn btn-icon" @click="prevMonth" :disabled="isLoading">
              <PanelLeftCloseIcon class="icon" />
            </button>

            <h3 class="text-center m-0 text-secondary">
              {{ monthNames[currentMonth] }} {{ currentYear }}
            </h3>

            <button class="btn btn-icon" @click="nextMonth" :disabled="isLoading">
              <PanelRightCloseIcon class="icon" />
            </button>
          </div>
          <button
            type="button"
            class="btn-close"
            data-bs-dismiss="modal"
            @click="cerrar_modal"
            aria-label="Close"
          ></button>
        </div>
        <div class="modal-body">
          <div class="calendar">
            <div class="calendar-header">
              <div class="day">Dom</div>
              <div class="day">Lun</div>
              <div class="day">Mar</div>
              <div class="day">Mié</div>
              <div class="day">Jue</div>
              <div class="day">Vie</div>
              <div class="day">Sáb</div>
            </div>
            <div class="calendar-body" v-if="isLoading">
              <div class="calendar-day" v-for="_ in 30">
                <div class="card placeholder-glow">
                  <div class="ratio ratio-21x9 card-img-top placeholder"></div>
                  <div class="card-body">
                    <div class="placeholder col-9 mb-3"></div>
                    <div class="placeholder placeholder-xs col-10"></div>
                    <div class="placeholder placeholder-xs col-11"></div>
                  </div>
                </div>
              </div>
            </div>
            <div class="calendar-body" v-else>
              <div v-for="x in fristnumday" class="calendar-day"></div>
              <div
                class="calendar-day"
                v-for="day in totaldiasMes"
                :class="{ 'has-data': search_find(day).length > 0 }"
              >
                <span class="date">{{ day }}</span>

                <ul v-for="x in search_find(day)" class="attendance-list">
                  <li v-if="x['r1']"><IconClock class="icon icon-sm" /> {{ x['r1'] }}</li>
                  <li v-if="x['r2']"><IconClock class="icon icon-sm" /> {{ x['r2'] }}</li>
                  <li v-if="x['r3']"><IconClock class="icon icon-sm" /> {{ x['r3'] }}</li>
                  <li v-if="x['r4']"><IconClock class="icon icon-sm" /> {{ x['r4'] }}</li>
                  <li v-if="x['r5']"><IconClock class="icon icon-sm" /> {{ x['r5'] }}</li>
                  <li v-if="x['r6']"><IconClock class="icon icon-sm" /> {{ x['r6'] }}</li>
                  <li v-if="x['r7']"><IconClock class="icon icon-sm" /> {{ x['r7'] }}</li>
                  <li v-if="x['r8']"><IconClock class="icon icon-sm" /> {{ x['r8'] }}</li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <button
            type="button"
            @click="cerrar_modal"
            class="btn btn-secondary"
            data-bs-dismiss="modal"
          >
            Cerrar
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { PanelLeftCloseIcon, PanelRightCloseIcon } from 'lucide-vue-next'
import { Buscar_Asistencia } from '@wails/services/PersonalService'
import { format, getDate, parseISO, addDays } from 'date-fns'
import { IconClock } from '@tabler/icons-vue'

const prop = defineProps({
  openmodal: { type: Boolean, required: true },
  dni: { type: String, required: true },
  nombre: { type: String, required: true }
})

const isopen = computed(() => prop.openmodal)

watch(isopen, async (newValue) => {
  try {
    if (newValue) {
      changeLoad()
    }
  } catch (error) {
    console.log(error)
    asistencia.value = []
  }
})

const changeLoad = async () => {
  try {
    isLoading.value = true
    totaldiasMes.value = new Date(currentYear.value, currentMonth.value + 1, 0).getDate()
    asistencia.value = await Buscar_Asistencia(prop.dni, currentMonth.value + 1, currentYear.value)
    fristnumday.value = new Date(currentYear.value, currentMonth.value, 1).getDay()
    isLoading.value = false
  } catch (error) {
    console.log(error)
  }
}

const cerrar_modal = () => {
  emit('close', false)
}

const emit = defineEmits(['close'])
const asistencia = ref<Array<any>>([])
const totaldiasMes = ref(30)
const fristnumday = ref()

const search_find = (day: number) => {
  if (asistencia.value) {
    return asistencia.value.filter((e) => {
      return getDate(format(addDays(parseISO(e.fecha), 1), 'yyyy/MM/dd')) == day
    })
  }
  return []
}
const monthNames = [
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

const currentMonth = ref(new Date().getMonth())
const currentYear = ref(new Date().getFullYear())
const isLoading = ref(false)

const prevMonth = () => {
  if (currentMonth.value === 0) {
    currentMonth.value = 11
    currentYear.value--
  } else {
    currentMonth.value--
  }
  changeLoad()
}

const nextMonth = () => {
  if (currentMonth.value === 11) {
    currentMonth.value = 0
    currentYear.value++
  } else {
    currentMonth.value++
  }
  changeLoad()
}
</script>

<style scoped>
.calendar {
  display: grid;
  grid-template-rows: auto 1fr;
  gap: 10px;
}

.calendar-header {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  text-align: center;
  font-weight: bold;
}

.calendar-body {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 5px;
}

.calendar-day {
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 5px;
  min-height: 100px;
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.calendar-day.has-data {
  background-color: #f8f9fa;
}

.calendar-day .date {
  font-size: 14px;
  font-weight: bold;
}

.attendance-list {
  margin: 0;
  padding: 0;
  list-style: none;
  font-size: 0.75rem;
  text-align: center;
  vertical-align: middle;
  width: 100%;
  margin-top: 5px;
}
</style>

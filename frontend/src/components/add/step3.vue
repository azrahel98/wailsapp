<template>
  <div class="container-fluid p-3">
    <div class="card card-sm">
      <div class="card-body">
        <form class="row g-3" @submit.prevent="siguiente">
          <div class="col-md-6 col-lg-6">
            <label for="dni" class="form-label">Area</label>
            <div class="input-group input-group-flat">
              <input
                type="text"
                v-model="area.nombre"
                class="form-control form-control-sm"
                placeholder="7024xxxx"
                autocomplete="off"
                required
              />
              <span
                class="input-group-text cursor-pointer"
                data-bs-toggle="modal"
                data-bs-target="#buscararea"
              >
                <IconSearch class="icon" />
              </span>
              <Modal
                id="buscararea"
                title="Buscar Areas disponibles"
                :cargo="false"
                @select="buscar"
              />
            </div>
          </div>
          <div class="col-md-6 col-lg-6">
            <label for="dni" class="form-label">Cargo</label>
            <div class="input-group input-group-flat">
              <input
                type="text"
                v-model="cargo.nombre"
                class="form-control form-control-sm"
                placeholder="7024xxxx"
                autocomplete="off"
                required
              />
              <span
                class="input-group-text cursor-pointer"
                data-bs-toggle="modal"
                data-bs-target="#buscarcargo"
                @click=""
              >
                <IconSearch class="icon" />
              </span>
              <Modal
                id="buscarcargo"
                title="Buscar cargos disponibles"
                :cargo="true"
                @select="buscar"
              />
            </div>
          </div>
          <div class="col-md-6 col-lg-6">
            <label for="dni" class="form-label">Regimen</label>
            <select class="form-select" required v-model="regimen">
              <option :value="x.id" v-for="x in regimenes">{{ x.name }}</option>
            </select>
          </div>
          <div class="mt-3 d-flex gap-2 justify-content-center">
            <button
              @click="() => emit('prevStep', true)"
              class="btn btn-outline-secondary btn-md p-1 fs-5"
            >
              Anterior
            </button>
            <button type="submit" class="btn btn-outline-facebook btn-md p-1 fs-5">
              Siguiente
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { IconSearch } from '@tabler/icons-vue'
import { ref } from 'vue'
import Modal from './modal.vue'

const area = ref({
  id: 0,
  nombre: ''
})
const cargo = ref({
  id: 0,
  nombre: ''
})

const emit = defineEmits(['nextStep', 'prevStep'])

const regimen = ref(1)
const regimenes = [
  { id: 1, name: 'CAS D.L 1057' },
  { id: 2, name: 'CAS D.U 083-2021' },
  { id: 3, name: 'CAS D.U 034-2021' },
  { id: 4, name: 'CAS LIMP' },
  { id: 5, name: '728' },
  { id: 6, name: '276-PER' },
  { id: 7, name: '276' },
  { id: 8, name: '276-OBRE' },
  { id: 9, name: 'CAS - F' },
  { id: 10, name: 'MODALIDAD FORMATIVA' },
  { id: 11, name: 'CAS-D' }
]

const siguiente = () => {
  emit('nextStep', {
    area: area.value.id,
    cargo: cargo.value.id,
    regimen: regimen.value
  })
}

const buscar = (res: any, iscargo: boolean) => {
  if (iscargo) {
    cargo.value.nombre = res.nombre
    cargo.value.id = res.id
  } else {
    area.value.nombre = res.nombre
    area.value.id = res.id
  }
}
</script>

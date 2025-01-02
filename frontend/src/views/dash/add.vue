<template>
  <div class="container">
    <div class="col text-start mt-4">
      <div class="page-pretitle fw-medium">Registrar a un nuevo</div>
      <h2 class="page-title">Trabajador</h2>
    </div>
    <div class="row">
      <div class="col-sm-12 col-md-auto px-3 steps-app">
        <ol class="steps steps-counter steps-vertical" aria-label="breadcrumbs">
          <li
            v-for="(step, index) in steps"
            :key="index"
            class="step-item"
            :class="{ active: currentStep === index + 1, disabled: currentStep < index + 1 }"
          >
            <a @click.prevent="goToStep(index + 1)">{{ index + 1 }}. {{ step.title }}</a>
          </li>
        </ol>
      </div>
      <div class="col-sm-12 col-md-8 card">
        <div class="card-body px-0 mx-0">
          <form @submit.prevent="submitForm">
            <Step1 v-if="currentStep === 1" @next-step="handle1" />

            <Step2 v-if="currentStep === 2" />

            <!-- Paso 3: Información Laboral -->
            <div v-if="currentStep === 3">
              <div class="mb-3">
                <label class="form-label">Cargo</label>
                <input v-model="formData.cargo" type="text" class="form-control" required />
              </div>
              <div class="mb-3">
                <label class="form-label">Área</label>
                <input v-model="formData.area" type="text" class="form-control" required />
              </div>
            </div>

            <!-- Paso 4: Resumen -->
            <div v-if="currentStep === 4">
              <h4 class="mb-3">Resumen de la Información</h4>
              <div v-for="(value, key) in formData" :key="key" class="mb-2">
                <strong>{{ key.charAt(0).toUpperCase() + key.slice(1) }}:</strong> {{ value }}
              </div>
            </div>

            <!-- <div class="d-flex justify-content-center gap-5">
              <button
                v-if="currentStep > 1"
                @click.prevent="prevStep"
                class="btn btn-sm btn-secondary"
              >
                Anterior
              </button>
              <button
                v-if="currentStep < 4"
                @click.prevent="nextStep"
                class="btn btn-sm btn-primary"
              >
                Siguiente
              </button>
              <button v-if="currentStep === 4" type="submit" class="btn btn-sm btn-success">
                Guardar
              </button>
            </div> -->
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import Step1 from '@comp/add//step1.vue'
import Step2 from '@comp/add/step2.vue'
import { ref, reactive } from 'vue'

const currentStep = ref(2)
const steps = [
  { title: 'Información Personal' },
  { title: 'Documento' },
  { title: 'Información Laboral' },
  { title: 'Resumen' }
]

const handle1 = (data) => {
  formDataF.persona = data
  nextStep()
}

const formData = reactive({
  dni: '',
  nombres: '',
  correo: '',
  direccion: '',
  telefono: '',
  tipoDocumento: '',
  numeroDocumento: '',
  fechaDocumento: '',
  descripcion: '',
  sueldo: '',
  cargo: '',
  area: ''
})

const formDataF = reactive({
  persona: {
    dni: '',
    aparterno: '',
    amaterno: '',
    nombres: '',
    direccion: '',
    telf1: null,
    nacimiento: '',
    sexo: 'M',
    email: ''
  },
  documento: {
    tipo: '',
    numero: '',
    año: '',
    fecha: null,
    fecha_valida: '',
    descripcion: '',
    sueldo: ''
  },
  vinculo: {
    dni: '',
    doc_ingreso_id: '',
    estado: '',
    area: '',
    cargo: '',
    regimen: ''
  }
})

const nextStep = () => {
  if (currentStep.value < 4) {
    currentStep.value++
  }
}

const prevStep = () => {
  if (currentStep.value > 1) {
    currentStep.value--
  }
}

const goToStep = (step) => {
  if (step <= currentStep.value) {
    currentStep.value = step
  }
}

const submitForm = () => {
  console.log('Datos del formulario:', formData)
  alert('Formulario guardado con éxito!')
  // Aquí puedes agregar la lógica para enviar los datos a tu backend
}
</script>

<style scoped lang="scss">
.container {
  display: grid;
  grid-template-rows: min-content auto;
  height: 100vh;
  .row {
    padding-top: 5vh;
    column-gap: 5vh;
    justify-items: center;
    justify-content: center;
    .steps-app {
      vertical-align: middle;
    }
    .card {
      height: min-content;
    }
  }
}
</style>

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
            <span @click.prevent="goToStep(index + 1)">{{ index + 1 }}. {{ step.title }}</span>
          </li>
        </ol>
      </div>
      <div class="col-sm-12 col-md-8 card">
        <div class="card-body px-0 mx-0">
          <div>
            <Step1 v-if="currentStep === 1" @next-step="handle1" />

            <Step2 v-if="currentStep === 2" @prev-step="prevStep" @next-step="handle2" />

            <Step3 v-if="currentStep === 3" @prev-step="prevStep" @next-step="handle3" />

            <div v-if="currentStep === 4">
              <h4 class="mb-3">Resumen de la Información</h4>
              <div v-for="(value, key) in formDataF" :key="key" class="mb-2">
                <strong>{{ key.charAt(0).toUpperCase() + key.slice(1) }}:</strong> {{ value }}
              </div>
            </div>
            <button type="button" @click="submitForm">click</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import Step1 from '@comp/add//step1.vue'
import Step2 from '@comp/add/step2.vue'
import Step3 from '@comp/add/step3.vue'
import { ref, reactive } from 'vue'
import { Crear_nuevoTrabajador } from '@wails/services/PersonalService'
import { router } from '@router/router'

const existe = ref(false)
const formDataF = reactive({
  persona: {},
  documento: {},
  vinculo: {
    dni: '',
    doc_ingreso_id: '',
    estado: 'activo',
    area: '',
    cargo: '',
    regimen: ''
  }
})
const currentStep = ref(1)
const steps = [
  { title: 'Información Personal' },
  { title: 'Documento' },
  { title: 'Información Laboral' },
  { title: 'Resumen' }
]

const handle1 = (data, x) => {
  formDataF.persona = data
  existe.value = x

  nextStep()
}

const handle2 = (data) => {
  formDataF.documento = data
  nextStep()
}

const handle3 = (data) => {
  formDataF.vinculo.area = data.area
  formDataF.vinculo.cargo = data.cargo
  formDataF.vinculo.regimen = data.regimen
  nextStep()
}

const prevStep = () => {
  if (currentStep.value > 1) {
    currentStep.value--
  }
}

const nextStep = () => {
  if (currentStep.value < 4) {
    currentStep.value++
  }
}

const goToStep = (step) => {
  if (step <= currentStep.value) {
    currentStep.value = step
  }
}

const submitForm = async () => {
  try {
    const x = await Crear_nuevoTrabajador(formDataF, existe.value)

    await router.push(`/perfil/${x}`)
  } catch (error) {
    console.log(error)
  }
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

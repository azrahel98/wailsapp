<template>
  <div
    class="modal modal-md fade"
    ref="modal"
    :id="`sindicato-${prop.id}`"
    tabindex="-1"
    aria-hidden="true"
  >
    +
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h1 class="modal-title fs-5">Afiliación a Sindicato</h1>
          <button
            type="button"
            class="btn-close"
            data-bs-dismiss="modal"
            aria-label="Close"
          ></button>
        </div>

        <div class="modal-body">
          <Step2 :sueldo="false" @next-step="handle1" v-if="step === 1" />

          <div v-if="step === 2">
            <h5>Paso 2: Selección del Sindicato</h5>
            <div class="mb-3">
              <label class="form-label">Seleccione un sindicato</label>
              <select v-model="sindicato" class="form-control">
                <option value="">Seleccione...</option>
                <option value="1">SOMUVES</option>
                <option value="2">SUTRAMUVES</option>
              </select>
            </div>
          </div>

          <div v-if="step === 3">
            <h5>Paso 3: Confirmación</h5>
            <p><strong>Tipo de Documento:</strong> {{ documento }}</p>
            <p><strong>Sindicato Seleccionado:</strong> {{ sindicato }}</p>
          </div>
        </div>

        <div class="modal-footer">
          <button v-if="step > 1" class="btn btn-secondary" @click="step--">Atrás</button>
          <button v-if="step < 3 && step !== 1" class="btn btn-primary" @click="nextStep">
            Siguiente
          </button>
          <button v-if="step === 3" class="btn btn-success" @click="submit(id)">Registrar</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Step2 from '@comp/add/step2.vue'
import { router } from '@router/router'
import { ref } from 'vue'
import { Crear_vinculoSindical } from '@wails/services/PersonalService'
import { models } from '@wails/models'

const prop = defineProps({
  id: { type: Number, required: true }
})

const step = ref(1)
const documento = ref<models.Documento>({
  tipoDocumento: '',
  fecha: '',
  descripcion: ''
})
const sindicato = ref(1)

const handle1 = (data: any) => {
  documento.value = data
  step.value++
}

const nextStep = () => {
  if (step.value === 2 && !sindicato.value) {
    alert('Por favor, seleccione un sindicato.')
    return
  }
  step.value++
}

const submit = async (id: number) => {
  try {
    await Crear_vinculoSindical(documento.value, id, parseInt(sindicato.value.toString()))
    router.go(0)
  } catch (error) {
    alert(error)
  }
}
</script>

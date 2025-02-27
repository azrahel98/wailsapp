<template>
  <div class="container-fluid p-3">
    <div class="card card-sm">
      <div class="card-body">
        <div class="row g-3">
          <div class="col-md-6">
            <label class="form-label">Tipo de Doc</label>
            <select
              v-if="sueldo"
              v-model="formData.tipoDocumento"
              :class="errors?.tipoDocumento ? 'is-invalid' : ''"
              class="form-select"
              required
            >
              <option value="contrato">Contrato</option>
              <option value="Resolucion">Resolucion</option>
              <option value="Acta">Acta</option>
            </select>
            <select
              v-else
              v-model="formData.tipoDocumento"
              :class="errors?.tipoDocumento ? 'is-invalid' : ''"
              class="form-select"
              required
            >
              <option value="Carta">Carta</option>
              <option value="Doc.Adm">Otros</option>
              <option value="Memorando">Memorando</option>
            </select>

            <span v-if="errors?.tipoDocumento">
              <p class="text-danger fs-6" v-for="x in errors.tipoDocumento._errors">{{ x }}</p>
            </span>
          </div>
          <div class="col-md-3">
            <label class="form-label">Número</label>
            <input
              v-model="formData.numeroDocumento"
              :class="errors?.numeroDocumento ? 'is-invalid' : ''"
              type="number"
              class="form-control"
              required
            />
            <span v-if="errors?.numeroDocumento">
              <p class="text-danger fs-6" v-for="x in errors.numeroDocumento._errors">{{ x }}</p>
            </span>
          </div>
          <div class="col-md-3">
            <label class="form-label">Año</label>
            <input
              v-model="formData.añoDocumento"
              :class="errors?.añoDocumento ? 'is-invalid' : ''"
              type="number"
              class="form-control"
              required
            />
            <span v-if="errors?.añoDocumento">
              <p class="text-danger fs-6" v-for="x in errors.añoDocumento._errors">{{ x }}</p>
            </span>
          </div>
          <div class="col-md-4">
            <label class="form-label">Fecha del Documento</label>
            <input
              v-model="formData.fechaValida"
              type="date"
              :class="errors?.fechaValida ? 'is-invalid' : ''"
              class="form-control"
              required
            />
            <span v-if="errors?.fechaValida">
              <p class="text-danger fs-6" v-for="x in errors.fechaValida._errors">{{ x }}</p>
            </span>
          </div>
          <div class="col-md-4">
            <label class="form-label">Fecha del Documento</label>
            <input
              v-model="formData.fecha"
              :class="errors?.fecha ? 'is-invalid' : ''"
              type="date"
              class="form-control"
              required
            />
            <span v-if="errors?.fecha">
              <p class="text-danger fs-6" v-for="x in errors.fecha._errors">{{ x }}</p>
            </span>
          </div>
          <div class="col-md-4" v-if="sueldo">
            <label class="form-label">Sueldo</label>
            <input
              v-model="formData.sueldo"
              type="number"
              step="0.01"
              :class="errors?.sueldo ? 'is-invalid' : ''"
              class="form-control"
              required
            />
            <span v-if="errors?.sueldo">
              <p class="text-danger fs-6" v-for="x in errors.sueldo._errors">{{ x }}</p>
            </span>
          </div>
          <div class="mb-3">
            <label class="form-label">Descripción</label>
            <textarea
              v-model="formData.descripcion"
              :class="errors?.descripcion ? 'is-invalid' : ''"
              class="form-control"
              rows="3"
              required
            />
            <span v-if="errors?.descripcion">
              <p class="text-danger fs-6" v-for="x in errors.descripcion._errors">{{ x }}</p>
            </span>
          </div>

          <div class="mt-3 d-flex gap-2 justify-content-center">
            <button @click="prevent" class="btn btn-outline-secondary btn-md p-1 fs-5">
              Anterior
            </button>
            <button type="button" @click="submit" class="btn btn-outline-facebook btn-md p-1 fs-5">
              Siguiente
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { getYear } from 'date-fns'
import { reactive, ref } from 'vue'
import { z } from 'zod'

defineProps({
  sueldo: { type: Boolean, required: false, default: true }
})

const emit = defineEmits(['nextStep', 'prevStep'])

const schema_validate = z.object({
  tipoDocumento: z.string().nonempty('Se debe seleccionar un tipo de documento'),
  numeroDocumento: z.number().int().positive('El numero debe de ser positivo'),
  añoDocumento: z.number().int().positive().min(2000, 'El año debe ser mayor al 2000'),
  fecha: z.string().date('La fecha no es válida'),
  fechaValida: z.string().date('La fecha no es válida'),
  descripcion: z.string().min(5, 'La descripción debe tener al menos 5 caracteres'),
  sueldo: z
    .number()
    .positive('El sueldo debe ser mayor a 0')
    .max(15600, 'El sueldo no puede ser mayor a 15600')
    .nullable()
})
type schema_validateType = z.infer<typeof schema_validate>
const errors = ref<z.ZodFormattedError<schema_validateType> | null>(null)

const formData = reactive({
  tipoDocumento: '',
  añoDocumento: '',
  numeroDocumento: '',
  descripcion: '',
  fechaValida: '',
  sueldo: null,
  fecha: ''
})

const submit = () => {
  errors.value = null
  const success = schema_validate.safeParse(formData)
  if (!success.success) {
    errors.value = success.error.format()
    console.log(errors.value)
  } else {
    if (
      getYear(new Date(formData.añoDocumento)) > getYear(new Date(formData.fecha)) ||
      getYear(new Date(formData.añoDocumento)) > getYear(new Date(formData.fechaValida))
    ) {
      errors.value = {
        añoDocumento: {
          _errors: ['El año debe ser menor o igual al año de la fecha del documento']
        },
        fechaValida: {
          _errors: ['El año de la fecha válida debe ser mayor o igual al año del documento']
        },
        _errors: []
      }
    } else {
      emit('nextStep', formData)
    }
  }
}

const prevent = () => {
  emit('prevStep')
}
</script>

<style lang="scss" scoped>
label,
input {
  font-size: 0.75rem;
}
</style>

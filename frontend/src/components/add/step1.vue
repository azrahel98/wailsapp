<template>
  <div class="container-fluid p-3">
    <div class="card card-sm">
      <div class="card-body">
        <form @submit.prevent="submitForm">
          <div class="row g-3">
            <div class="col-md-6 col-lg-3">
              <label for="dni" class="form-label">DNI</label>
              <div class="input-group input-group-flat">
                <input
                  type="text"
                  v-model="formData.dni"
                  class="form-control"
                  :class="errors?.dni ? 'is-invalid' : ''"
                  placeholder="7024xxxx"
                  autocomplete="off"
                />
                <span
                  class="input-group-text cursor-pointer"
                  :class="errors?.dni ? 'border-red text-danger' : ''"
                  @click="buscar_dni"
                >
                  <IconSearch class="icon" />
                </span>
              </div>
              <span v-if="errors?.dni">
                <p class="text-danger fs-6" v-for="x in errors.dni._errors">{{ x }}</p>
              </span>
            </div>
            <div class="col-md-6 col-lg-3">
              <label for="nombres" class="form-label">Nombres</label>
              <input
                v-model="formData.nombres"
                type="text"
                class="form-control"
                id="nombres"
                :class="errors?.nombres ? 'is-invalid' : ''"
                required
              />
              <span v-if="errors?.nombres">
                <p class="text-danger fs-6" v-for="x in errors.nombres._errors">{{ x }}</p>
              </span>
            </div>
            <div class="col-md-6 col-lg-3">
              <label for="apaterno" class="form-label">Apellido Paterno</label>
              <input
                v-model="formData.apaterno"
                type="text"
                :class="errors?.apaterno ? 'is-invalid' : ''"
                class="form-control"
                id="apaterno"
                required
              />
              <span v-if="errors?.apaterno">
                <p class="text-danger fs-6" v-for="x in errors.apaterno._errors">{{ x }}</p>
              </span>
            </div>
            <div class="col-md-6 col-lg-3">
              <label for="amaterno" class="form-label">Apellido Materno</label>
              <input
                v-model="formData.amaterno"
                type="text"
                class="form-control"
                id="amaterno"
                required
                :class="errors?.amaterno ? 'is-invalid' : ''"
              />
              <span v-if="errors?.amaterno">
                <p class="text-danger fs-6" v-for="x in errors.amaterno._errors">{{ x }}</p>
              </span>
            </div>
            <div class="col-md-4">
              <label for="email" class="form-label">Correo Electrónico</label>
              <input
                v-model="formData.email"
                type="email"
                class="form-control"
                :class="errors?.email ? 'is-invalid' : ''"
                id="email"
                :disabled="formData.existe"
                required
              />
              <span v-if="errors?.email">
                <p class="text-danger fs-6" v-for="x in errors.email._errors">{{ x }}</p>
              </span>
            </div>
            <div class="col-md-4">
              <label for="telf1" class="form-label">Teléfono</label>
              <input
                v-model="formData.telf1"
                type="tel"
                class="form-control"
                :class="errors?.telf1 ? 'is-invalid' : ''"
                id="telf1"
                :disabled="formData.existe"
                required
              />
              <span v-if="errors?.telf1">
                <p class="text-danger fs-6" v-for="x in errors.telf1._errors">{{ x }}</p>
              </span>
            </div>
            <div class="col-md-4">
              <label for="nacimiento" class="form-label">Fecha de Nacimiento</label>
              <input
                v-model="formData.nacimiento"
                type="date"
                class="form-control"
                :disabled="formData.existe"
                id="nacimiento"
                required
              />
            </div>
            <div class="col-8">
              <label for="direccion" class="form-label">Dirección</label>
              <input
                v-model="formData.direccion"
                type="text"
                class="form-control"
                id="direccion"
                :class="errors?.direccion ? 'is-invalid' : ''"
                required
                :disabled="formData.existe"
              />
              <span v-if="errors?.direccion">
                <p class="text-danger fs-6" v-for="x in errors.direccion._errors">{{ x }}</p>
              </span>
            </div>

            <div class="col-md-4">
              <label class="form-label d-block">Sexo</label>
              <div class="form-selectgroup">
                <label class="form-selectgroup-item">
                  <input
                    v-model="formData.sexo"
                    type="radio"
                    name="sexo"
                    :disabled="formData.existe"
                    value="M"
                    class="form-selectgroup-input"
                  />
                  <span class="form-selectgroup-label fs-6 p-1">Masculino</span>
                </label>
                <label class="form-selectgroup-item">
                  <input
                    v-model="formData.sexo"
                    type="radio"
                    name="sexo"
                    :disabled="formData.existe"
                    value="F"
                    class="form-selectgroup-input"
                  />
                  <span class="form-selectgroup-label fs-6 p-1">Femenino</span>
                </label>
              </div>
            </div>
          </div>
          <div class="mt-3">
            <button type="submit" class="btn btn-outline-facebook btn-md p-1 fs-5">
              Siguiente
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { IconSearch } from '@tabler/icons-vue'
import { reactive, ref } from 'vue'
import { Buscar_persona_by_dni, Buscar_dni_onlinne } from '@wails/services/PersonalService'
import { z } from 'zod'
import { parseISO } from 'date-fns'

interface FormData {
  dni: string
  apaterno: string
  amaterno: string
  nombres: string
  direccion: string | undefined
  telf1: string | undefined
  nacimiento: string
  sexo: string | undefined
  email: string | undefined
  existe?: boolean
}
const phoneRegex = /^\d{9}$/
const schema_validate = z.object({
  dni: z.string().length(8, 'El DNI debe tener 8 caracteres'),
  apaterno: z.string().nonempty(),
  amaterno: z.string().nonempty(),
  nombres: z.string().nonempty(),
  direccion: z.string().nonempty(),
  telf1: z.string().regex(phoneRegex, { message: 'El número de teléfono no es válido' }),
  email: z.string().email('El correo electrónico no es válido'),
  nacimiento: z.string().date('La fecha no es válida')
})
type schema_validateType = z.infer<typeof schema_validate>
const errors = ref<z.ZodFormattedError<schema_validateType> | null>(null)

const formData: FormData = reactive({
  dni: '',
  apaterno: '',
  amaterno: '',
  nombres: '',
  direccion: '',
  telf1: '',
  nacimiento: '',
  sexo: 'M',
  email: '',
  existe: false
})

const buscar_dni = async () => {
  try {
    errors.value = null
    if (formData.dni.length != 8) {
      errors.value = {
        dni: {
          _errors: ['El DNI debe tener 8 caracteres']
        },
        _errors: []
      }
    } else {
      const user = await Buscar_persona_by_dni(formData.dni)
      formData.apaterno = user.Aparterno
      formData.amaterno = user.Amaterno
      formData.nombres = user.Nombre
      formData.telf1 = user.Telf1
      formData.direccion = user.Direccion
      formData.nacimiento = parseISO(user.Nacimiento).toISOString().split('T')[0]
      formData.sexo = user.Sexo
      formData.email = user.Email
      formData.existe = true
    }
  } catch (error) {
    try {
      const res = await Buscar_dni_onlinne(formData.dni)
      formData.apaterno = res.apellidoPaterno
      formData.amaterno = res.apellidoMaterno
      formData.nombres = res.nombres
    } catch (secondError) {
      console.log(secondError)
    }
  }
}

const emit = defineEmits(['nextStep'])

const submitForm = () => {
  errors.value = null
  const result = schema_validate.safeParse(formData)
  if (result.success || formData.existe) {
    emit('nextStep', formData)
  } else {
    errors.value = result.error.format()
  }
}
</script>

<style scoped>
.form-selectgroup {
  display: inline-flex;
}

.form-selectgroup-item {
  margin-right: 0.5rem;
}
label,
input {
  font-size: 0.75rem;
}

@media (max-width: 768px) {
  .card-body {
    padding: 1rem;
  }
}
</style>

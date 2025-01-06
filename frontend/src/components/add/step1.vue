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
                  v-model="formData.Dni"
                  class="form-control"
                  :class="errors?.Dni ? 'is-invalid' : ''"
                  placeholder="7024xxxx"
                  autocomplete="off"
                />
                <span
                  class="input-group-text cursor-pointer"
                  :class="errors?.Dni ? 'border-red text-danger' : ''"
                  @click="buscar_dni"
                >
                  <IconSearch class="icon" />
                </span>
              </div>
              <span v-if="errors?.Dni">
                <p class="text-danger fs-6" v-for="x in errors.Dni._errors">{{ x }}</p>
              </span>
            </div>
            <div class="col-md-6 col-lg-3">
              <label for="nombres" class="form-label">Nombres</label>
              <input
                v-model="formData.Nombre"
                type="text"
                class="form-control"
                id="nombres"
                :class="errors?.Nombre ? 'is-invalid' : ''"
                required
              />
              <span v-if="errors?.Nombre">
                <p class="text-danger fs-6" v-for="x in errors.Nombre._errors">{{ x }}</p>
              </span>
            </div>
            <div class="col-md-6 col-lg-3">
              <label for="apaterno" class="form-label">Apellido Paterno</label>
              <input
                v-model="formData.Apaterno"
                type="text"
                :class="errors?.Apaterno ? 'is-invalid' : ''"
                class="form-control"
                id="apaterno"
                required
              />
              <span v-if="errors?.Apaterno">
                <p class="text-danger fs-6" v-for="x in errors.Apaterno._errors">{{ x }}</p>
              </span>
            </div>
            <div class="col-md-6 col-lg-3">
              <label for="amaterno" class="form-label">Apellido Materno</label>
              <input
                v-model="formData.Amaterno"
                type="text"
                class="form-control"
                id="Amaterno"
                required
                :class="errors?.Amaterno ? 'is-invalid' : ''"
              />
              <span v-if="errors?.Amaterno">
                <p class="text-danger fs-6" v-for="x in errors.Amaterno._errors">{{ x }}</p>
              </span>
            </div>
            <div class="col-md-4">
              <label for="email" class="form-label">Correo Electrónico</label>
              <input
                v-model="formData.Email"
                type="email"
                class="form-control"
                :class="errors?.Email ? 'is-invalid' : ''"
                id="email"
                :disabled="existe"
                required
              />
              <span v-if="errors?.Email">
                <p class="text-danger fs-6" v-for="x in errors.Email._errors">{{ x }}</p>
              </span>
            </div>
            <div class="col-md-4">
              <label for="telf1" class="form-label">Teléfono</label>
              <input
                v-model="formData.Telf1"
                type="tel"
                class="form-control"
                :class="errors?.Telf1 ? 'is-invalid' : ''"
                id="telf1"
                :disabled="existe"
                required
              />
              <span v-if="errors?.Telf1">
                <p class="text-danger fs-6" v-for="x in errors.Telf1._errors">{{ x }}</p>
              </span>
            </div>
            <div class="col-md-4">
              <label for="nacimiento" class="form-label">Fecha de Nacimiento</label>
              <input
                v-model="formData.Nacimiento"
                type="date"
                class="form-control"
                :disabled="existe"
                id="nacimiento"
                required
              />
            </div>
            <div class="col-12">
              <label for="direccion" class="form-label">Dirección</label>
              <input
                v-model="formData.Direccion"
                type="text"
                class="form-control"
                id="direccion"
                :class="errors?.Direccion ? 'is-invalid' : ''"
                required
                :disabled="existe"
              />
              <span v-if="errors?.Direccion">
                <p class="text-danger fs-6" v-for="x in errors.Direccion._errors">{{ x }}</p>
              </span>
            </div>

            <div class="col-md-6">
              <label class="form-label d-block">Sexo</label>
              <div class="form-selectgroup">
                <label class="form-selectgroup-item">
                  <input
                    v-model="formData.Sexo"
                    type="radio"
                    name="sexo"
                    :disabled="existe"
                    value="M"
                    class="form-selectgroup-input"
                  />
                  <span class="form-selectgroup-label fs-6 p-1">Masculino</span>
                </label>
                <label class="form-selectgroup-item">
                  <input
                    v-model="formData.Sexo"
                    type="radio"
                    name="sexo"
                    :disabled="existe"
                    value="F"
                    class="form-selectgroup-input"
                  />
                  <span class="form-selectgroup-label fs-6 p-1">Femenino</span>
                </label>
              </div>
            </div>
            <div class="col-6">
              <label for="direccion" class="form-label">Ruc</label>
              <input
                v-model="formData.Ruc"
                type="text"
                class="form-control"
                id="Ruc"
                :class="errors?.Ruc ? 'is-invalid' : ''"
                required
                :disabled="existe"
              />
              <span v-if="errors?.Ruc">
                <p class="text-danger fs-6" v-for="x in errors.Ruc._errors">{{ x }}</p>
              </span>
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
import { models } from '@wails/models'

const phoneRegex = /^\d{9}$/
const schema_validate = z.object({
  Dni: z.string().length(8, 'El DNI debe tener 8 caracteres'),
  Apaterno: z.string().nonempty(),
  Amaterno: z.string().nonempty(),
  Nombre: z.string().nonempty(),
  Direccion: z.string().nonempty(),
  Ruc: z.string().min(11).max(11),
  Telf1: z.string().regex(phoneRegex, { message: 'El número de teléfono no es válido' }),
  Email: z.string().email('El correo electrónico no es válido'),
  Nacimiento: z.string().date('La fecha no es válida')
})
type schema_validateType = z.infer<typeof schema_validate>
const errors = ref<z.ZodFormattedError<schema_validateType> | null>(null)

const formData: models.Perfil = reactive({
  Dni: '',
  Apaterno: '',
  Amaterno: '',
  Nombre: '',
  Ruc: '',
  Direccion: '',
  Telf1: '',
  Email: '',
  Nacimiento: '',
  Sexo: 'M'
})

const existe = ref(false)

const buscar_dni = async () => {
  try {
    errors.value = null
    if (formData.Dni.length != 8) {
      errors.value = {
        Dni: {
          _errors: ['El DNI debe tener 8 caracteres']
        },
        _errors: []
      }
    } else {
      const user = await Buscar_persona_by_dni(formData.Dni)
      formData.Apaterno = user.Apaterno
      formData.Amaterno = user.Amaterno
      formData.Nombre = user.Nombre
      formData.Telf1 = user.Telf1
      formData.Direccion = user.Direccion
      formData.Nacimiento = parseISO(user.Nacimiento).toISOString().split('T')[0]
      formData.Sexo = user.Sexo
      formData.Email = user.Email
      existe.value = true
    }
  } catch (error) {
    try {
      const res = await Buscar_dni_onlinne(formData.Dni)
      formData.Apaterno = res.apellidoPaterno
      formData.Amaterno = res.apellidoMaterno
      formData.Nombre = res.nombres
    } catch (secondError) {
      console.log(secondError)
    }
  }
}

const emit = defineEmits(['nextStep'])

const submitForm = () => {
  errors.value = null
  const result = schema_validate.safeParse(formData)
  if (result.success || existe.value) {
    emit('nextStep', formData, existe.value)
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

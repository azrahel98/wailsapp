<template>
  <div
    class="modal fade"
    id="editmodal"
    tabindex="-1"
    aria-labelledby="editmodal"
    aria-hidden="true"
  >
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h1 class="modal-title fs-5" id="exampleModalLabel">Editar Informacion</h1>
          <button
            type="button"
            class="btn-close"
            data-bs-dismiss="modal"
            aria-label="Close"
          ></button>
        </div>
        <div class="modal-body">
          <form @submit.prevent="guardar(user)">
            <div class="row">
              <div class="col-md-12 mb-3">
                <label for="nombre" class="form-label">Nombre</label>
                <input
                  type="text"
                  id="nombre"
                  class="form-control"
                  v-model="user.Nombre"
                  disabled
                  aria-disabled="true"
                />
              </div>

              <div class="col-md-6 mb-3">
                <label for="telefono" class="form-label">Teléfono</label>
                <input
                  type="tel"
                  id="telefono"
                  :class="user.Telf1 ? '' : 'border-warning border-1'"
                  class="form-control"
                  v-model="user.Telf1"
                  placeholder="Ingrese su teléfono"
                  required
                />
                <span
                  v-if="errors?.Telf1"
                  class="invalid-feedback fs-6"
                  v-for="x in errors.Telf1._errors"
                  >{{ x }}</span
                >
              </div>

              <div class="col-md-6 mb-3">
                <label for="direccion" class="form-label">Dirección</label>
                <input
                  type="text"
                  id="direccion"
                  class="form-control"
                  :class="user.Direccion ? '' : 'border-warning border-1'"
                  v-model="user.Direccion"
                  required
                  placeholder="Ingrese su dirección"
                />
                <span
                  v-if="errors?.Direccion"
                  class="invalid-feedback fs-6"
                  v-for="x in errors.Direccion._errors"
                  >{{ x }}</span
                >
              </div>

              <div class="col-md-6 mb-3">
                <label for="cuenta" class="form-label">Ruc</label>
                <input
                  type="text"
                  id="cuenta"
                  class="form-control"
                  v-model="user.Ruc"
                  :class="user.Ruc ? '' : 'border-warning border-1'"
                  placeholder="Ingrese su numero de RUC"
                />
              </div>

              <div class="col-md-6 mb-3">
                <label for="email" class="form-label">Email</label>
                <input
                  type="email"
                  id="email"
                  class="form-control"
                  :class="user.Email ? '' : 'border-warning border-1'"
                  v-model="user.Email"
                  placeholder="Ingrese su email"
                  required
                />
                <span
                  v-if="errors?.Email"
                  class="text-danger fs-5"
                  v-for="x in errors.Email._errors"
                  >{{ x }}</span
                >
              </div>
            </div>
            <div class="modal-footer bg-transparent">
              <button type="button" class="btn" data-bs-dismiss="modal" aria-label="Close">
                Cancelar
              </button>
              <button type="submit" class="btn btn-primary">Guardar</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { router } from '@router/router'
import { EditByDni } from '@wails/services/PersonalService'
import { ref } from 'vue'
import { z } from 'zod'
const phoneRegex = /^\d{9}$/
const schema_validate = z.object({
  Telf1: z.string().regex(phoneRegex, { message: 'El número de teléfono no es válido' }),
  Direccion: z.string().min(5),
  Email: z.string().email('El correo es invalido'),
  Ruc: z.string().min(11).max(11)
})
type schema_validateType = z.infer<typeof schema_validate>
const errors = ref<z.ZodFormattedError<schema_validateType> | null>(null)

const guardar = (user: any) => {
  try {
    console.log('click aquii')
    errors.value = null
    const valid = schema_validate.safeParse(user)
    if (!valid.success) {
      errors.value = valid.error.format()
    }
    if (valid.success) {
      EditByDni(user.Telf1, user.Telf2, user.Direccion, user.Email, user.Dni, user.Ruc)
      router.go(0)
    }
  } catch (error) {
    console.log(error)
  }
}

defineProps({
  user: { type: Object, required: true }
})
</script>

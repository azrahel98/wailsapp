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
          <form @submit.prevent="">
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
                />
              </div>

              <div class="col-md-6 mb-3">
                <label for="direccion" class="form-label">Dirección</label>
                <input
                  type="text"
                  id="direccion"
                  class="form-control"
                  :class="user.Direccion ? '' : 'border-warning border-1'"
                  v-model="user.Direccion"
                  placeholder="Ingrese su dirección"
                />
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
                />
              </div>
            </div>
          </form>
        </div>

        <div class="modal-footer">
          <button type="button" class="btn" data-bs-dismiss="modal" aria-label="Close">
            Cancelar
          </button>
          <button type="submit" class="btn btn-primary" @click="guardar(user)">Guardar</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { EditByDni } from '@wails/services/PersonalService'

const guardar = (user: any) => {
  try {
    const x = EditByDni(user.Telf1, user.Telf2, user.Direccion, user.Email, user.Dni)
    console.log(x)
  } catch (error) {
    console.log(error)
  }
}

defineProps({
  user: { type: Object, required: true }
})
</script>

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
          <h1 class="modal-title fs-5" id="exampleModalLabel">Modal title</h1>
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
              <!-- Nombre -->
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

              <!-- Teléfono -->
              <div class="col-md-6 mb-3">
                <label for="telefono" class="form-label">Teléfono</label>
                <input
                  type="tel"
                  id="telefono"
                  class="form-control"
                  v-model="user.Telf1"
                  placeholder="Ingrese su teléfono"
                />
              </div>

              <!-- Dirección -->
              <div class="col-md-6 mb-3">
                <label for="direccion" class="form-label">Dirección</label>
                <input
                  type="text"
                  id="direccion"
                  class="form-control"
                  v-model="user.Direccion"
                  placeholder="Ingrese su dirección"
                />
              </div>

              <!-- Cuenta Bancaria -->
              <div class="col-md-6 mb-3">
                <label for="cuenta" class="form-label">Cuenta Bancaria</label>
                <input
                  type="text"
                  id="cuenta"
                  class="form-control"
                  v-model="user.Ruc"
                  placeholder="Ingrese su cuenta bancaria"
                />
              </div>

              <div class="col-md-6 mb-3">
                <label for="email" class="form-label">Email</label>
                <input
                  type="email"
                  id="email"
                  class="form-control"
                  v-model="user.Email"
                  placeholder="Ingrese su email"
                />
              </div>
            </div>
          </form>
        </div>

        <div class="modal-footer">
          <button type="button" class="btn btn-secondary me-2" @click="">Cancelar</button>
          <button type="submit" class="btn btn-primary">Guardar</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { models } from '@wails/models'

defineProps({
  user: { type: models.Perfil, required: true }
})
</script>

<template>
  <div
    class="modal fade"
    id="salidmodal"
    tabindex="-1"
    aria-labelledby="salidmodal"
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
                <input type="text" id="nombre" class="form-control" disabled aria-disabled="true" />
              </div>

              <div class="col-md-6 mb-3">
                <label for="telefono" class="form-label">Teléfono</label>
                <input
                  type="tel"
                  id="telefono"
                  class="form-control"
                  placeholder="Ingrese su teléfono"
                />
              </div>

              <div class="col-md-6 mb-3">
                <label for="direccion" class="form-label">Dirección</label>
                <input
                  type="text"
                  id="direccion"
                  class="form-control"
                  placeholder="Ingrese su dirección"
                />
              </div>

              <div class="col-md-6 mb-3">
                <label for="cuenta" class="form-label">Ruc</label>
                <input
                  type="text"
                  id="cuenta"
                  class="form-control"
                  placeholder="Ingrese su numero de RUC"
                />
              </div>

              <div class="col-md-6 mb-3">
                <label for="email" class="form-label">Email</label>
                <input
                  type="email"
                  id="email"
                  class="form-control"
                  placeholder="Ingrese su email"
                />
              </div>
            </div>
          </form>
        </div>

        <div class="modal-footer">
          <button type="button" class="btn btn-secondary me-2" @click="">Cancelar</button>
          <button type="submit" class="btn btn-primary" @click="">Guardar</button>
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
</script>

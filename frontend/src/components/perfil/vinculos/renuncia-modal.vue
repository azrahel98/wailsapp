<template>
  <div class="modal fade" :id="`${prop.id}`" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h1 class="modal-title fs-5" id="exampleModalLabel">Registrar Renuncia</h1>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" />
        </div>
        <div class="modal-body">
          <form @submit.prevent="">
            <div class="row">
              <div class="col-md-6 mb-3">
                <label for="nombre" class="form-label">Tipo de Documento</label>
                <select class="form-select">
                  <option value="Carta">Carta</option>
                  <option value="Resolucion">Resolucion</option>
                  <option value="Acta">Acta</option>
                  <option value="Doc.Adm">Doc.Adm</option>
                </select>
              </div>

              <div class="col-md-3 mb-3">
                <label for="telefono" class="form-label">Numero</label>
                <input type="number" id="telefono" class="form-control" placeholder="##" />
              </div>

              <div class="col-md-3 mb-3">
                <label for="direccion" class="form-label">Año</label>
                <input type="number" id="direccion" class="form-control" placeholder="202$" />
              </div>

              <div class="col-md-6 mb-3">
                <label for="cuenta" class="form-label">Fecha del Documento</label>
                <input
                  type="date"
                  id="date"
                  class="form-control"
                  placeholder="Ingrese su numero de RUC"
                />
              </div>

              <div class="col-md-6 mb-3">
                <label for="email" class="form-label">Fecha Valida</label>
                <input type="date" id="date" class="form-control" placeholder="Ingrese su email" />
              </div>
              <div class="col-md-12 mb-3">
                <label for="email" class="form-label">Descripcion</label>
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
          <button type="submit" class="btn btn-primary" @click="">Guardar</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { EditByDni } from '@wails/services/PersonalService'

const prop = defineProps({
  id: { type: Number, reuired: true }
})
</script>

<style lang="scss" scoped>
label {
  font-size: 0.82rem;
}

input,
select {
  font-size: 0.75rem;
  text-align: center;
}
</style>

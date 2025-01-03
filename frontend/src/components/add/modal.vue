<template>
  <div class="modal fade" :id="id" tabindex="-1" :aria-labelledby="id" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h1 class="modal-title fs-5" id="exampleModalLabel">{{ title }}</h1>
          <button
            type="button"
            class="btn-close"
            data-bs-dismiss="modal"
            aria-label="Close"
          ></button>
        </div>
        <div class="modal-body">
          <form @submit.prevent="">
            <div class="row g-3 justify-content-center">
              <div class="col-md-6 px-3">
                <label for="nombre" class="form-label">Nombre</label>
                <input
                  type="text"
                  id="nombre"
                  class="form-control border"
                  autocomplete="off"
                  v-on:keyup.enter="buscar(cargo)"
                  v-model="search"
                />
              </div>

              <div class="col-md-12 mb-3">
                <div class="card">
                  <div class="table-responsive">
                    <table class="table table-vcenter table-mobile-md card-table">
                      <thead>
                        <tr>
                          <th>Nombre</th>
                          <th>Cantidad</th>
                          <th class="w-1"></th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="item in resultado">
                          <td data-label="Name">{{ item.nombre }}</td>
                          <td data-label="Title">
                            {{ item.cantidad }}
                          </td>
                          <td>
                            <div class="btn-list flex-nowrap">
                              <a @click="select(id, item, cargo)" class="btn"> Edit </a>
                            </div>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
            </div>
            <div class="modal-footer">
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

<script setup lang="ts">
import { ref } from 'vue'
import { Buscar_Areas, Buscar_Cargos } from '@wails/services/PersonalService'
import { defineEmits } from 'vue'

defineProps({
  title: {
    type: String,
    required: true
  },
  id: {
    type: String,
    required: true
  },
  cargo: { type: Boolean, default: true }
})

const search = ref('')
const resultado = ref<Array<any>>([])

const emit = defineEmits(['select'])

const buscar = async (cargo: boolean) => {
  try {
    if (cargo) {
      resultado.value = await Buscar_Cargos(search.value)
    } else {
      resultado.value = await Buscar_Areas(search.value)
    }
  } catch (error) {
    console.log(error)
  }
}

const select = (id: string, res: any, iscargo: boolean) => {
  emit('select', res, iscargo)
  console.log(res, iscargo)
  const modalElement = document.getElementById(id)
  if (modalElement) {
    modalElement.classList.remove('show')
    modalElement.setAttribute('aria-hidden', 'true')
    modalElement.style.display = 'none'
    document.body.classList.remove('modal-open')
    const backdrop = document.querySelector('.modal-backdrop')
    if (backdrop) {
      backdrop.remove()
    }
  }
}
</script>

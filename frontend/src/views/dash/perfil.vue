<template>
  <div class="main">
    <div class="text-start pt-2">
      <div class="page-pretitle">perfil</div>
      <h2 class="page-title">{{ perfil.Nombre }}</h2>
    </div>

    <div class="row row-cards pt-2 m-0">
      <div
        class="col-md-4 col-sm-12 d-flex flex-column gap-4 justify-content-center align-items-center"
      >
        <div class="avatar" style="height: 100%; width: 150px">
          <img
            src="https://encrypted-tbn3.gstatic.com/licensed-image?q=tbn:ANd9GcR6pHfNWfIcZ9h948cUbNqKWt1qJvJvzFwf6JPr8cgDa7SFoWf5CNkwS_AgnTMkD_nxvY-y6eVyS9r6YWw"
            class="rounded-5"
          />
        </div>
      </div>
      <div class="col-md-2">
        <div class="d-flex gap-3 pt-3">
          <button class="btn btn-outline-facebook btn-md btn-icon">
            <IconCamera class="icon" />
          </button>
          <button class="btn btn-md btn-outline-warning btn-icon"><IconEdit class="icon" /></button>
        </div>
      </div>

      <div class="col-md-6 col-sm-12">
        <div class="card rounded-3 card-info shadow-sm">
          <div class="card-body infogrid">
            <div class="d-flex align-items-center mb-2">
              <span class="badge bg-facebook">DNI:</span>
              <span class="mx-2 page-subtitle align-middle text-center">{{ perfil.Dni }}</span>
            </div>
            <div class="d-flex align-items-center mb-2">
              <span class="badge" :class="perfil.Email ? 'bg-facebook-lt' : 'bg-warning-lt'"
                >CORREO:</span
              >
              <span class="mx-2 page-subtitle align-middle text-center" v-if="perfil.Email">{{
                perfil.Email
              }}</span>
              <span class="mx-3 status-dot bg-warning status-dot-animated" v-else />
            </div>
            <div class="d-flex align-items-center mb-2">
              <span class="badge" :class="perfil.Telf1 ? 'bg-facebook' : 'bg-youtube'">TELF:</span>
              <span class="mx-2 page-subtitle align-middle text-center" v-if="perfil.Telf1">{{
                perfil.Telf1
              }}</span>
              <span class="mx-3 status-dot bg-youtube status-dot-animated" v-else />
            </div>
            <div class="d-flex align-items-center mb-2">
              <span class="badge" :class="perfil.Nacimiento ? 'bg-facebook-lt' : 'bg-warning-lt'"
                >CUMPLEAÑOS:</span
              >
              <span class="mx-2 page-subtitle align-middle text-center" v-if="perfil.Nacimiento">{{
                perfil.Nacimiento.split('T')[0]
              }}</span>
              <span class="mx-3 status-dot bg-warning status-dot-animated" v-else />
            </div>
            <div class="d-flex align-items-center mb-2">
              <span class="badge" :class="perfil.Ruc ? 'bg-facebook-lt' : 'bg-warning-lt'"
                >RUC:</span
              >
              <span class="mx-2 page-subtitle align-middle text-center" v-if="perfil.Ruc">{{
                perfil.Ruc
              }}</span>
              <span class="mx-3 status-dot bg-warning status-dot-animated" v-else />
            </div>
            <div class="d-flex align-items-center mb-2">
              <span class="badge" :class="perfil.Sexo ? 'bg-facebook-lt' : 'bg-warning-lt'"
                >SEXO:</span
              >
              <span
                class="mx-2 page-subtitle align-middle text-center fw-bold text-secondary"
                v-if="perfil.Sexo"
                >{{ perfil.Sexo }}</span
              >
              <span class="mx-3 status-dot bg-warning status-dot-animated" v-else />
            </div>
          </div>
        </div>
      </div>

      <Contratos />
    </div>
  </div>

  <Modalinfo :user="perfil!" />
  <!-- Modal para editar/agregar trabajo -->
  <!-- <div
    class="modal modal-blur fade"
    :class="{ show: isEditWorkModalOpen }"
    tabindex="-1"
    style="display: block"
    v-if="isEditWorkModalOpen"
  >
    <div class="modal-dialog modal-lg" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">{{ editingWork.id ? 'Editar Trabajo' : 'Agregar Trabajo' }}</h5>
          <button type="button" class="btn-close" @click="closeEditWorkModal"></button>
        </div>
        <div class="modal-body">
          <div class="mb-3">
            <label class="form-label">Empresa</label>
            <input type="text" class="form-control" v-model="editingWork.company" />
          </div>
          <div class="mb-3">
            <label class="form-label">Posición</label>
            <input type="text" class="form-control" v-model="editingWork.position" />
          </div>
          <div class="mb-3">
            <label class="form-label">Fecha de Inicio</label>
            <input type="date" class="form-control" v-model="editingWork.startDate" />
          </div>
          <div class="mb-3">
            <label class="form-label">Fecha de Fin</label>
            <input type="date" class="form-control" v-model="editingWork.endDate" />
          </div>
          <div class="mb-3">
            <label class="form-label">Régimen</label>
            <input type="text" class="form-control" v-model="editingWork.regime" />
          </div>
          <div class="mb-3">
            <label class="form-label">Salario</label>
            <input type="text" class="form-control" v-model="editingWork.salary" />
          </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-link link-secondary" @click="closeEditWorkModal">
            Cancelar
          </button>
          <button type="button" class="btn btn-primary ms-auto" @click="saveWorkInfo">
            {{ editingWork.id ? 'Guardar cambios' : 'Agregar trabajo' }}
          </button>
        </div>
      </div>
    </div>
  </div> -->
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { Buscar_persona_by_dni } from '@wails/services/PersonalService'
import { models } from '@wails/models'
import {
  IconEdit,
  IconPhone,
  IconMapPin,
  IconCreditCard,
  IconCamera,
  IconEditCircle,
  IconBandageFilled,
  IconMoneybag,
  IconCards
} from '@tabler/icons-vue'
import { router } from '@router/router'
import Modalinfo from '@comp/perfil/avatar/modal_info.vue'
import Contratos from '@comp/perfil/vinculos/contratos.vue'
import { BanknoteIcon } from 'lucide-vue-next'
import { format, formatDate } from 'date-fns'

const perfil = ref<models.Perfil>({
  Dni: '',
  Nombre: '',
  Nacimiento: ''
})

onMounted(async () => {
  try {
    perfil.value = await Buscar_persona_by_dni(router.currentRoute.value.params.dni.toString())
    console.log(perfil.value)
  } catch (error) {
    console.log(error)
  }
})
</script>
<style lang="scss" scoped>
.main {
  height: 100vh;
  display: grid;
  row-gap: 4px;
  grid-template-rows: min-content min-content;
  .card-info {
    justify-content: center;
    align-items: start;
    justify-items: start;
  }

  .infogrid {
    width: 100%;
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(50%, 1fr));
    justify-content: start;
    justify-items: start;
    .d-flex {
      width: 100%;
      text-align: end;
    }
  }
}
</style>

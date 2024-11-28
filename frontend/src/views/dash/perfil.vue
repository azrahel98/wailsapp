<template>
  <div class="main">
    <div class="text-start pt-2">
      <div class="page-pretitle">Overview</div>
      <h2 class="page-title">Dashboard</h2>
    </div>

    <div class="row row-cards p-0 m-0">
      <div class="col-md-6 col-lg-4">
        <div class="card p-0 m-0">
          <div class="card-body p-0 m-0 text-center">
            <div class="grid-avatar">
              <span class="avatar avatar-xl mb-3 avatar-rounded">
                <img :src="`${perfil.Foto}`" class="border-1 border-secondary" />
              </span>
              <div class="d-flex flex-wrap justify-content-center align-content-center gap-2">
                <input
                  type="file"
                  id="avatar-upload"
                  class="d-none"
                  accept="image/*"
                  @change="handleAvatarChange"
                />
                <label for="avatar-upload" class="btn btn-outline-primary btn-sm mt-2">
                  Cambiar foto
                </label>
                <button
                  type="button"
                  class="card-btn btn btn-sm p-0 m-0"
                  data-bs-toggle="modal"
                  data-bs-target="#editmodal"
                >
                  <IconEdit />
                  Editar Información
                </button>
              </div>
            </div>
            <h3 class="m-0 mb-1">{{ perfil?.Nombre }}</h3>
            <div class="text-muted">{{ perfil?.Email }}</div>
          </div>
        </div>
      </div>
      <div class="col-md-6 col-lg-8">
        <div class="card">
          <div class="card-header">
            <h3 class="card-title">Información Personal</h3>
          </div>
          <div class="card-body">
            <div class="datagrid">
              <div class="datagrid-item">
                <div class="datagrid-title">Teléfono</div>
                <div class="datagrid-content">
                  <IconPhone class="icon-inline text-muted me-1" />
                  {{ perfil?.Telf1 }}
                </div>
              </div>
              <div class="datagrid-item">
                <div class="datagrid-title">Teléfono</div>
                <div class="datagrid-content">
                  <IconPhone class="icon-inline text-muted me-1" />
                  {{ perfil?.Dni }}
                </div>
              </div>
              <div class="datagrid-item">
                <div class="datagrid-title">Teléfono</div>
                <div class="datagrid-content">
                  <IconPhone class="icon-inline text-muted me-1" />
                  {{ perfil?.Telf2 }}
                </div>
              </div>
              <div class="datagrid-item">
                <div class="datagrid-title">Dirección</div>
                <div class="datagrid-content">
                  <IconMapPin class="icon-inline text-muted me-1" />
                  {{ perfil?.Direccion }}
                </div>
              </div>
              <div class="datagrid-item">
                <div class="datagrid-title">Cuenta Bancaria</div>
                <div class="datagrid-content">
                  <IconCreditCard class="icon-inline text-muted me-1" />
                  {{ perfil?.Ruc }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Contratos />
    </div>
  </div>

  <Modalinfo :user="perfil!" />
  <!-- Modal para editar/agregar trabajo -->
  <div
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
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { Buscar_persona_by_dni } from '@wails/services/PersonalService'
import { models } from '@wails/models'
import { IconEdit, IconPhone, IconMapPin, IconCreditCard } from '@tabler/icons-vue'
import { router } from '@router/router'
import Modalinfo from '@comp/perfil/avatar/modal_info.vue'
import Contratos from '@comp/perfil/vinculos/contratos.vue'

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

interface User {
  name: string
  avatar: string
  phone: string
  address: string
  bankAccount: string
  email: string
}

interface WorkHistory {
  id: number
  company: string
  startDate: string
  endDate: string
  regime: string
  salary: string
  position: string
}

const user = reactive<User>({
  name: 'María García',
  avatar: '/placeholder.svg?height=200&width=200',
  phone: '+34 123 456 789',
  address: 'Calle Principal 123, Madrid, España',
  bankAccount: 'ES91 2100 0418 4502 0005 1332',
  email: 'maria.garcia@ejemplo.com'
})

const workHistory = ref<WorkHistory[]>([
  {
    id: 1,
    company: 'Tech Innovators S.L.',
    startDate: '2020-03-15',
    endDate: '2022-06-30',
    regime: 'Tiempo completo',
    salary: '35.000€',
    position: 'Desarrollador Frontend'
  },
  {
    id: 2,
    company: 'Digital Solutions Inc.',
    startDate: '2018-01-10',
    endDate: '2020-02-28',
    regime: 'Tiempo parcial',
    salary: '25.000€',
    position: 'Diseñador UX/UI'
  },
  {
    id: 3,
    company: 'StartUp Ventures',
    startDate: '2022-08-01',
    endDate: 'Presente',
    regime: 'Tiempo completo',
    salary: '45.000€',
    position: 'Ingeniero de Software Senior'
  }
])

const isEditUserModalOpen = ref(false)
const isEditWorkModalOpen = ref(false)
const editingUser = reactive<User>({ ...user })
const editingWork = reactive<WorkHistory>({
  id: 0,
  company: '',
  startDate: '',
  endDate: '',
  regime: '',
  salary: '',
  position: ''
})

const openEditUserModal = () => {
  Object.assign(editingUser, user)
  isEditUserModalOpen.value = true
}

const closeEditUserModal = () => {
  isEditUserModalOpen.value = false
}

const saveUserInfo = () => {
  Object.assign(user, editingUser)
  closeEditUserModal()
}

const openEditWorkModal = (work: WorkHistory) => {
  Object.assign(editingWork, work)
  isEditWorkModalOpen.value = true
}

const openAddWorkModal = () => {
  Object.assign(editingWork, {
    id: 0,
    company: '',
    startDate: '',
    endDate: '',
    regime: '',
    salary: '',
    position: ''
  })
  isEditWorkModalOpen.value = true
}

const closeEditWorkModal = () => {
  isEditWorkModalOpen.value = false
}

const saveWorkInfo = () => {
  if (editingWork.id) {
    const index = workHistory.value.findIndex((w) => w.id === editingWork.id)
    if (index !== -1) {
      workHistory.value[index] = { ...editingWork }
    }
  } else {
    workHistory.value.push({ ...editingWork, id: Date.now() })
  }
  closeEditWorkModal()
}

const handleAvatarChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files[0]) {
    const file = target.files[0]
    const reader = new FileReader()
    reader.onload = (e) => {
      if (e.target?.result) {
        user.avatar = e.target.result as string
      }
    }
    reader.readAsDataURL(file)
  }
}
</script>
<style lang="scss" scoped>
.main {
  height: 100vh;
  display: grid;
  row-gap: 4px;
  grid-template-rows: 6vh auto auto;
}
.grid-avatar {
  width: 100%;
  height: 100%;
  display: grid;
  grid-template-columns: min-content min-content;
  .btn-sm {
    width: min-content;
    word-wrap: break-word;
    white-space: wrap;
  }
}
</style>

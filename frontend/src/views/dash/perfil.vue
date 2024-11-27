<template>
  <div class="page">
    <div class="page-wrapper">
      <div class="container-xl">
        <div class="page-header d-print-none">
          <div class="row align-items-center">
            <div class="col">
              <h2 class="page-title">Dashboard de Usuario</h2>
            </div>
          </div>
        </div>
      </div>
      <div class="page-body">
        <div class="container-xl">
          <div class="row row-cards">
            <div class="col-md-6 col-lg-4">
              <div class="card">
                <div class="card-body p-4 text-center">
                  <span
                    class="avatar avatar-xl mb-3 avatar-rounded"
                    :style="{ backgroundImage: `url(${user.avatar})` }"
                  ></span>
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
                  <h3 class="m-0 mb-1">{{ perfil?.Nombre }}</h3>
                  <div class="text-muted">{{ perfil?.Email }}</div>
                  <div class="mt-3">
                    <span class="badge bg-purple-lt">Usuario Premium</span>
                  </div>
                </div>
                <div class="d-flex">
                  <a href="#" class="card-btn" @click.prevent="openEditUserModal">
                    <IconEdit />
                    Editar Información
                  </a>
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
            <div class="col-12">
              <div class="card">
                <div class="card-header">
                  <h3 class="card-title">Historial Laboral</h3>
                  <div class="card-actions">
                    <a href="#" class="btn btn-primary" @click.prevent="openAddWorkModal">
                      <IconPlus class="icon-inline me-1" />
                      Agregar Trabajo
                    </a>
                  </div>
                </div>
                <div class="card-body">
                  <div class="table-responsive">
                    <table class="table table-vcenter card-table">
                      <thead>
                        <tr>
                          <th>Empresa</th>
                          <th>Posición</th>
                          <th>Período</th>
                          <th>Régimen</th>
                          <th>Salario</th>
                          <th>Acciones</th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="job in workHistory" :key="job.id">
                          <td>{{ job.company }}</td>
                          <td class="text-muted">
                            <IconBriefcase class="icon-inline me-1" />
                            {{ job.position }}
                          </td>
                          <td class="text-muted">
                            <IconCalendar class="icon-inline me-1" />
                            {{ job.startDate }} - {{ job.endDate }}
                          </td>
                          <td class="text-muted">{{ job.regime }}</td>
                          <td>
                            <span class="badge bg-success">{{ job.salary }}</span>
                          </td>
                          <td>
                            <a
                              href="#"
                              class="btn btn-sm btn-outline-secondary"
                              @click.prevent="openEditWorkModal(job)"
                            >
                              <IconEdit class="icon-inline me-1" />
                              Editar
                            </a>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal para editar información personal -->
    <div
      class="modal modal-blur fade"
      :class="{ show: isEditUserModalOpen }"
      tabindex="-1"
      style="display: block"
      v-if="isEditUserModalOpen"
    >
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Editar Información Personal</h5>
            <button type="button" class="btn-close" @click="closeEditUserModal"></button>
          </div>
          <div class="modal-body">
            <div class="mb-3">
              <label class="form-label">Nombre</label>
              <input type="text" class="form-control" v-model="editingUser.name" />
            </div>
            <div class="mb-3">
              <label class="form-label">Teléfono</label>
              <input type="tel" class="form-control" v-model="editingUser.phone" />
            </div>
            <div class="mb-3">
              <label class="form-label">Dirección</label>
              <input type="text" class="form-control" v-model="editingUser.address" />
            </div>
            <div class="mb-3">
              <label class="form-label">Cuenta Bancaria</label>
              <input type="text" class="form-control" v-model="editingUser.bankAccount" />
            </div>
            <div class="mb-3">
              <label class="form-label">Email</label>
              <input type="email" class="form-control" v-model="editingUser.email" />
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-link link-secondary" @click="closeEditUserModal">
              Cancelar
            </button>
            <button type="button" class="btn btn-primary ms-auto" @click="saveUserInfo">
              Guardar cambios
            </button>
          </div>
        </div>
      </div>
    </div>

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
  </div>
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
  IconPlus,
  IconBriefcase,
  IconCalendar
} from '@tabler/icons-vue'
import { router } from '@router/router'

const perfil = ref<models.Perfil>()

onMounted(async () => {
  try {
    perfil.value = await Buscar_persona_by_dni(router.currentRoute.value.params.dni.toString())
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

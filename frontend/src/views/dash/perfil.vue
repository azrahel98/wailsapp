<template>
  <div class="page-wrapper">
    <div class="page-header">
      <div class="container">
        <div class="col text-start">
          <div class="page-pretitle fw-medium">Perfil e Informacion Personal</div>
        </div>
        <div class="row align-items-center text-start">
          <div class="col-auto">
            <span class="avatar avatar-lg rounded">
              <img :src="perfil.Foto" v-if="perfil.Foto" />
              <img src="../../assets/images/man.svg" v-else-if="perfil.Sexo == 'M'" />
              <img src="../../assets/images/mujer.svg" v-else />
            </span>
          </div>
          <div class="col">
            <h1 class="fw-bold">{{ perfil.Nombre }}</h1>
            <div class="list-inline list-inline-dots text-secondary">
              <div class="list-inline-item">
                <IconCardboards class="icon" /> Dni:
                {{ perfil.Dni }}
              </div>
            </div>
          </div>
          <div class="col-auto ms-auto">
            <div class="btn-list">
              <button class="btn btn-icon btn-outline-facebook px-2">
                <IconEdit class="icon p-0 m-0" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="page-body">
      <div class="container-xl">
        <div class="row g-3">
          <div class="col vinculos">
            <ul class="timeline">
              <li class="timeline-event text-start" v-for="x in vinculos">
                <div
                  class="timeline-event-icon"
                  :class="x.Estado !== 'activo' ? 'bg-danger-lt' : 'bg-facebook-lt'"
                >
                  <IconBriefcase
                    class="icon text-secondary"
                    :class="x.Estado !== 'activo' ? 'text-danger' : 'text-facebook'"
                  />
                </div>
                <div class="card timeline-event-card">
                  <div class="card-body">
                    <div class="text-secondary float-end">
                      {{
                        x.Fecha_ingreso
                          ? format(parseISO(x.Fecha_ingreso), 'yyyy/MM/dd')
                          : 'Fecha no disponible'
                      }}
                    </div>
                    <div class="page-subtitle">
                      <h4 class="p-0 m-0">{{ x.Cargo }}</h4>
                      <p class="text-secondary p-0 m-0">
                        {{ x.Area }}
                      </p>
                    </div>
                    <div class="info">
                      <div class="accordion" id="accordion-example">
                        <div class="accordion-item border-0 p-0 m-0">
                          <h2
                            class="accordion-header text-center d-flex justify-content-between pt-2 align-content-center align-items-center"
                            id="heading-1"
                          >
                            <button
                              :id="'collapse#' + x.Id"
                              class="btn btn-action"
                              data-bs-toggle="collapse"
                              :data-bs-target="'#collapse' + x.Id"
                              aria-expanded
                              @click="click_collapse(x.Id)"
                            >
                              <IconEyePlus class="icon text-facebook fw-bold me-1" />
                              <IconEyeX
                                id="abierto"
                                class="icon me-1 text-warning fw-bold d-none"
                              />
                            </button>
                            <div class="d-flex justify-content-center align-items-center">
                              <button class="btn-action">
                                <IconTrashX class="icon text-youtube" />
                              </button>
                              <button class="btn-action">
                                <IconEdit class="icon text-success" />
                              </button>
                              <span
                                class="badge fs-6 text-uppercase"
                                :class="x.Estado !== 'activo' ? 'bg-danger' : 'bg-green'"
                              >
                                {{ x.Estado }}
                              </span>
                            </div>
                          </h2>
                          <div
                            :id="'collapse' + x.Id"
                            class="accordion-collapse collapse"
                            data-bs-parent="#accordion-example"
                            style=""
                          >
                            <div class="accordion-body pt-0">
                              <div class="detalles-collapse">
                                <div class="ingreso">
                                  <div class="mb-2">
                                    <IconFileInfo class="icon text-info" />
                                    Ingreso: <strong>{{ x.Doc_i }}-{{ x.Numero_doc_i }}</strong>
                                  </div>
                                  <div class="mb-2">
                                    <IconFileInfo class="icon text-info" />
                                    Descripcion: <strong>{{ x.Descrip_ingre }}</strong>
                                  </div>
                                  <div class="mb-2">
                                    <IconFileInfo class="icon text-info" />
                                    Regimen: <strong>{{ x.Regimen }}</strong>
                                  </div>
                                  <div class="mb-2">
                                    <IconFileInfo class="icon text-info" />
                                    Sueldo: <strong>S/.{{ x.Sueldo }}</strong>
                                  </div>
                                </div>
                                <div class="salida">
                                  <div class="mb-2" v-if="x.Doc_s">
                                    <IconClipboardOff class="icon text-danger" />
                                    Salida: <strong>{{ x.Doc_s }}-{{ x.Numero_doc_s }}</strong>
                                  </div>
                                  <div class="mb-2" v-if="x.Doc_s">
                                    <IconClipboardOff class="icon text-danger" />
                                    Descripcion: <strong>{{ x.Descrip_s }}</strong>
                                  </div>
                                  <div class="mb-2" v-if="x.Doc_s">
                                    <IconClipboardOff class="icon text-danger" />
                                    Fecha Renuncia:
                                    <strong>
                                      {{
                                        x.Fecha_salida
                                          ? format(parseISO(x.Fecha_salida), 'yyyy/MM/dd')
                                          : 'Fecha no disponible'
                                      }}
                                    </strong>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </li>
            </ul>
          </div>
          <div class="col-lg-4">
            <div class="row row-cards">
              <div class="col-12">
                <div class="card">
                  <div class="card-body text-start">
                    <div class="card-title">Informacion Basica</div>
                    <div class="mb-2">
                      <IconAddressBook class="icon me-2 text-secondary" />
                      Vive en: <strong v-if="perfil.Direccion">{{ perfil.Direccion }}</strong>
                      <div class="status-dot status-dot-animated bg-youtube mx-2" />
                    </div>
                    <div class="mb-2">
                      <IconBriefcase class="icon me-2 text-secondary" />
                      Correo es: <strong v-if="perfil.Email">{{ perfil.Email }}</strong>
                      <div class="status-dot status-dot-animated bg-youtube mx-2" />
                    </div>
                    <div class="mb-2">
                      <IconCashBanknote class="icon me-2 text-secondary" />
                      Ruc: <strong v-if="perfil.Ruc">{{ perfil.Ruc }}</strong>
                      <div class="status-dot status-dot-animated bg-youtube mx-2" />
                    </div>
                    <div class="mb-2">
                      <IconPhoneCall class="icon me-2 text-secondary" />
                      Telefono:
                      <strong v-if="perfil.Telf1"
                        ><span class="flag flag-country-si"></span> {{ perfil.Telf1 }}</strong
                      >
                      <div class="status-dot status-dot-animated bg-youtube mx-2" />
                    </div>
                    <div class="mb-2" v-if="perfil.Telf2">
                      <IconPhoneCall class="icon me-2 text-secondary" />
                      Telefono 2:
                      <strong><span class="flag flag-country-si"></span> {{ perfil.Telf1 }}</strong>
                    </div>
                    <div class="mb-2" v-if="parseISO(perfil.Nacimiento)">
                      <CakeIcon class="icon me-2 text-secondary" />
                      Cumpleaños:
                      <strong>
                        {{
                          perfil.Nacimiento
                            ? format(parseISO(perfil.Nacimiento), 'yyyy/MM/dd')
                            : 'Fecha no disponible'
                        }}
                      </strong>
                    </div>
                    <div class="mb-2" v-if="parseISO(perfil.Nacimiento)">
                      <IconMoodHappy class="icon me-2 text-secondary" />
                      Edad:
                      <strong>
                        {{ getYear(new Date()) - getYear(parseISO(perfil.Nacimiento)) }} años
                      </strong>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Buscar_persona_by_dni, Search_by_dni_vinculos } from '@wails/services/PersonalService'
import { models } from '@wails/models'
import { router } from '@router/router'
import {
  IconAddressBook,
  IconBriefcase,
  IconCardboards,
  IconCashBanknote,
  IconClipboardOff,
  IconEdit,
  IconEyePlus,
  IconEyeX,
  IconFileInfo,
  IconMailAi,
  IconMoodHappy,
  IconOutlet,
  IconPhoneCall,
  IconTrashX
} from '@tabler/icons-vue'
import { format, formatDate, getYear, parse, parseISO } from 'date-fns'
import { CakeIcon } from 'lucide-vue-next'

const click_collapse = (x: number) => {
  document.getElementById(`collapse#${x}`)?.classList.toggle('toggle-collapse')
}

const perfil = ref<models.Perfil>({
  Dni: '',
  Nombre: '',
  Nacimiento: ''
})

const expanded = ref(true)
const vinculos = ref<Array<any>>()

onMounted(async () => {
  try {
    perfil.value = await Buscar_persona_by_dni(router.currentRoute.value.params.dni.toString())
    vinculos.value = await Search_by_dni_vinculos(router.currentRoute.value.params.dni.toString())
  } catch (error) {
    console.log(error)
  }
})
</script>

<style lang="scss" scoped>
.vinculos {
  height: 100%;
}

.toggle-collapse {
  svg {
    display: none;
  }
  #abierto {
    display: block !important;
  }
}
.detalles-collapse {
  display: flex;
  flex-wrap: wrap;
}
</style>

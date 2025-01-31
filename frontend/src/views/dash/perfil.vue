<template>
  <div class="main">
    <div class="page-header mt-0">
      <div class="container">
        <div class="text-start">
          <div class="page-pretitle fw-medium">Perfil e Informacion Personal</div>
        </div>
      </div>
    </div>
    <div class="row align-items-center text-start pt-2">
      <div class="col-auto">
        <span class="avatar avatar-lg rounded">
          <img :src="perfil.Foto" v-if="perfil.Foto" />
          <img src="../../assets/images/man.svg" v-else-if="perfil.Sexo == 'M'" />
          <img src="../../assets/images/mujer.svg" v-else />
        </span>
      </div>
      <div class="col-6">
        <h2 class="fw-bold m-0">{{ perfil.Nombre }} {{ perfil.Apaterno }} {{ perfil.Amaterno }}</h2>
        <div class="p-0 m-0 fs-4">
          <IconCardboards class="icon" /> DNI:
          {{ perfil.Dni }}
        </div>
      </div>
      <div class="col-2 ms-auto">
        <div class="btn-list">
          <button
            class="btn btn-outline-facebook p-1"
            data-bs-toggle="modal"
            data-bs-target="#editmodal"
          >
            <IconEdit class="icon p-0 m-0" />
          </button>

          <button
            class="btn btn-outline-warning p-1"
            data-bs-toggle="modal"
            @click="() => (open_asistencia = true)"
            data-bs-target="#searchcalendar"
          >
            <CalendarSearchIcon class="icon p-0 m-0" />
          </button>

          <modal_info :user="perfil" />
          <modal_asistencia
            :openmodal="open_asistencia"
            @close="(x) => (open_asistencia = x)"
            :dni="perfil.Dni"
            :nombre="perfil.Nombre"
          />
        </div>
      </div>
    </div>
    <div class="page-body">
      <div class="container-xl">
        <div class="row gap-3">
          <div class="col vinculos" v-if="!isloading">
            <ul class="timeline">
              <li class="timeline-event text-start" v-for="x in vinculos">
                <div
                  class="timeline-event-icon"
                  :class="x.Estado !== 'activo' ? 'bg-youtube ' : 'bg-facebook'"
                >
                  <IconBriefcase
                    class="icon"
                    :class="x.Estado !== 'activo' ? 'text-white' : 'text-white'"
                  />
                </div>
                <div class="card timeline-event-card">
                  <div class="card-body p-0 m-0 mx-2 my-2">
                    <div class="text-secondary float-end fs-5">
                      {{
                        x.Fecha_ingreso
                          ? format(addDays(parseISO(x.Fecha_ingreso), 1), 'yyyy/MM/dd')
                          : 'Fecha no disponible'
                      }}
                    </div>
                    <div class="page-subtitle">
                      <h4 class="p-0 m-0 fs-5">{{ x.Cargo }}</h4>
                      <p class="text-secondary p-0 m-0 fs-6">
                        {{ x.Area }}
                      </p>
                    </div>
                    <div class="info p-0 m-0">
                      <div class="accordion" id="accordion-example">
                        <div class="accordion-item border-0 p-0 m-0">
                          <div
                            class="accordion-header text-center w-100 d-flex justify-content-between align-content-center align-items-center"
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
                            <div class="d-flex justify-content-end gap-1 align-items-center">
                              <button class="btn-action btn btn-icon btn-md">
                                <IconTrashX class="icon text-youtube" />
                              </button>
                              <button class="btn-action btn btn-icon">
                                <IconEdit class="icon text-success" />
                              </button>
                              <button
                                v-if="!x.Doc_s"
                                class="btn-action btn btn-icon"
                                data-bs-toggle="modal"
                                :data-bs-target="`#${x.Id}`"
                              >
                                <IconX class="icon text-warning" />
                              </button>
                              <renunciaModal :id="x.Id" v-if="!x.Doc_s" />
                            </div>
                          </div>
                          <div
                            :id="'collapse' + x.Id"
                            class="accordion-collapse collapse"
                            data-bs-parent="#accordion-example"
                            style=""
                          >
                            <div class="accordion-body pt-0">
                              <div class="detalles-collapse">
                                <div class="ingreso">
                                  <div class="mb-2 fs-5">
                                    <IconFileInfo class="icon text-info" />
                                    Ingreso: <strong>{{ x.Doc_i }} {{ x.Numero_doc_i }}</strong>
                                  </div>
                                  <div class="mb-2 fs-5">
                                    <IconFileInfo class="icon text-info" />
                                    Descripcion: <strong>{{ x.Descrip_ingre }}</strong>
                                  </div>
                                  <div class="mb-2 fs-5">
                                    <IconFileInfo class="icon text-info" />
                                    Regimen: <strong>{{ x.Regimen }}</strong>
                                  </div>
                                  <div class="mb-2 fs-5">
                                    <IconFileInfo class="icon text-info" />
                                    Sueldo: <strong>S/.{{ x.Sueldo }}</strong>
                                  </div>
                                </div>
                                <div class="salida fs-5">
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
                                          ? format(
                                              addDays(parseISO(x.Fecha_salida), 1),
                                              'yyyy/MM/dd'
                                            )
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
          <div v-else class="col d-flex flex-column h-100 gap-4">
            <div class="card placeholder-glow" v-for="_ in 4">
              <div class="placeholder col-9 mb-3 pt-2"></div>
            </div>
          </div>

          <div class="col-lg-5">
            <div class="row row-cards">
              <div class="col-12">
                <div class="card" v-if="!isloading">
                  <div class="card-body px-2 py-0 text-start">
                    <div class="card-title fs-4 p-0 m-0 py-2">Informacion Basica</div>
                    <div class="mb-1 fs-5">
                      <IconAddressBook class="icon me-2 text-secondary fs-6" />
                      Vive en: <strong v-if="perfil.Direccion">{{ perfil.Direccion }}</strong>
                      <div
                        v-if="!perfil.Direccion"
                        class="status-dot status-dot-animated bg-youtube mx-2"
                      />
                    </div>
                    <div class="mb-1 fs-5">
                      <IconBriefcase class="icon me-2 text-secondary" />
                      Correo es: <strong v-if="perfil.Email">{{ perfil.Email }}</strong>
                      <div
                        v-if="!perfil.Email"
                        class="status-dot status-dot-animated bg-youtube mx-2"
                      />
                    </div>
                    <div class="mb-1 fs-5">
                      <IconCashBanknote class="icon me-2 text-secondary" />
                      Ruc: <strong v-if="perfil.Ruc">{{ perfil.Ruc }}</strong>
                      <div
                        v-if="!perfil.Ruc"
                        class="status-dot status-dot-animated bg-youtube mx-2"
                      />
                    </div>
                    <div class="mb-1 fs-5">
                      <IconPhoneCall class="icon me-2 text-secondary" />
                      Telefono:
                      <strong v-if="perfil.Telf1"
                        ><span class="flag flag-country-si"></span> {{ perfil.Telf1 }}</strong
                      >
                      <div
                        v-if="!perfil.Telf1"
                        class="status-dot status-dot-animated bg-youtube mx-2"
                      />
                    </div>
                    <div class="mb-1 fs-5" v-if="perfil.Telf2">
                      <IconPhoneCall class="icon me-2 text-secondary" />
                      Telefono 2:
                      <strong><span class="flag flag-country-si"></span> {{ perfil.Telf1 }}</strong>
                    </div>
                    <div class="mb-1 fs-5" v-if="parseISO(perfil.Nacimiento)">
                      <CakeIcon class="icon me-2 text-secondary" />
                      Cumpleaños:
                      <strong>
                        {{
                          perfil.Nacimiento
                            ? format(addDays(parseISO(perfil.Nacimiento), 1), 'yyyy/MM/dd')
                            : 'Fecha no disponible'
                        }}
                      </strong>
                    </div>
                    <div class="mb-1 fs-5" v-if="parseISO(perfil.Nacimiento)">
                      <IconMoodHappy class="icon me-2 text-secondary" />
                      Edad:
                      <strong>
                        {{ getYear(new Date()) - getYear(parseISO(perfil.Nacimiento)) }} años
                      </strong>
                    </div>
                  </div>
                </div>
                <div class="card placeholder-glow" v-else>
                  <div class="card-body pt-5">
                    <div class="placeholder col-9 mb-3"></div>
                    <div class="placeholder placeholder-lg col-10"></div>
                    <div class="placeholder placeholder-xs col-11"></div>
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
import modal_info from '@comp/perfil/avatar/modal_info.vue'
import renunciaModal from '@comp/perfil/vinculos/renuncia-modal.vue'
import modal_asistencia from '@comp/perfil/asistencia/moda.vue'
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
  IconMoodHappy,
  IconPhoneCall,
  IconTrashX,
  IconX
} from '@tabler/icons-vue'
import { format, getYear, parseISO, addDays } from 'date-fns'
import { CakeIcon, CalendarSearchIcon } from 'lucide-vue-next'

const click_collapse = (x: number) => {
  document.getElementById(`collapse#${x}`)?.classList.toggle('toggle-collapse')
}

const perfil = ref<models.Perfil>({
  Dni: '',
  Nacimiento: '',
  Nombre: '',
  Apaterno: '',
  Amaterno: ''
})

const vinculos = ref<Array<any>>()
const isloading = ref(false)

const open_asistencia = ref(false)

onMounted(async () => {
  try {
    isloading.value = true
    perfil.value = await Buscar_persona_by_dni(router.currentRoute.value.params.dni.toString())
    vinculos.value = await Search_by_dni_vinculos(router.currentRoute.value.params.dni.toString())
  } catch (error) {
    console.log(error)
  } finally {
    isloading.value = false
  }
})
</script>

<style lang="scss" scoped>
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

.main {
  max-height: 100vh;
  display: grid;
  grid-template-rows: 2vh 18vh 80vh;
  .page-body {
    max-height: 100%;
    height: 100%;
    padding: 0;
    margin: 0;
    .container-xl,
    .row {
      height: 100%;

      .vinculos {
        max-height: 100%;
        overflow-y: auto;
      }
    }
    @media (max-width: 960px) {
      .container-xl,
      .row {
        height: 100%;
        .vinculos {
          max-height: 50%;
        }
      }
    }
  }
  strong {
    font-size: 0.8rem;
  }
}
</style>

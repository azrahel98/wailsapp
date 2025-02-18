<template>
  <div class="main">
    <div class="px-4 pt-2">
      <div class="max-w-7xl mx-auto">
        <div class="space-y-1">
          <span class="text-sm font-medium"> Informacion Personal </span>
          <h2 class="font-bold text-gray tracking-tight">Perfil</h2>
        </div>
      </div>
    </div>

    <div
      class="row align-items-center align-content-end text-start p-4 pt-0"
      style="max-width: 80vw"
    >
      <div class="col-auto">
        <span class="avatar avatar-lg rounded">
          <img :src="perfil.Foto" v-if="perfil.Foto" />
          <img src="../../assets/images/man.svg" v-else-if="perfil.Sexo == 'M'" />
          <img src="../../assets/images/mujer.svg" v-else />
        </span>
      </div>
      <div class="col-6">
        <h2 class="fw-bold m-0">{{ perfil.Nombre }} {{ perfil.Apaterno }} {{ perfil.Amaterno }}</h2>
        <div class="p-0 m-0">
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
            class="btn btn-outline-bitbucket p-1"
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
    <div class="page-body pt-4">
      <div class="container-xl">
        <div class="row gap-3">
          <div class="col vinculos" v-if="!isloading">
            <ul class="timeline">
              <Timeline :x="x" v-for="x in vinculos" :click_collapse="click_collapse" />
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
                    <div class="card-title p-0 m-0 py-2">Informacion Basica</div>
                    <div class="mb-1">
                      <IconAddressBook class="icon me-2 text-secondary" />
                      Vive en: <strong v-if="perfil.Direccion">{{ perfil.Direccion }}</strong>
                      <div
                        v-if="!perfil.Direccion"
                        class="status-dot status-dot-animated bg-youtube mx-2"
                      />
                    </div>
                    <div class="mb-1">
                      <IconBriefcase class="icon me-2 text-secondary" />
                      Correo es: <strong v-if="perfil.Email">{{ perfil.Email }}</strong>
                      <div
                        v-if="!perfil.Email"
                        class="status-dot status-dot-animated bg-youtube mx-2"
                      />
                    </div>
                    <div class="mb-1">
                      <IconCashBanknote class="icon me-2 text-secondary" />
                      Ruc: <strong v-if="perfil.Ruc">{{ perfil.Ruc }}</strong>
                      <div
                        v-if="!perfil.Ruc"
                        class="status-dot status-dot-animated bg-youtube mx-2"
                      />
                    </div>
                    <div class="mb-1">
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
                    <div class="mb-1" v-if="perfil.Telf2">
                      <IconPhoneCall class="icon me-2 text-secondary" />
                      Telefono 2:
                      <strong><span class="flag flag-country-si"></span> {{ perfil.Telf1 }}</strong>
                    </div>
                    <div class="mb-1" v-if="parseISO(perfil.Nacimiento)">
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
                    <div class="mb-1" v-if="parseISO(perfil.Nacimiento)">
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
import modal_asistencia from '@comp/perfil/asistencia/moda.vue'
import { Buscar_persona_by_dni, Search_by_dni_vinculos } from '@wails/services/PersonalService'
import { models } from '@wails/models'
import { router } from '@router/router'
import {
  IconAddressBook,
  IconBriefcase,
  IconCardboards,
  IconCashBanknote,
  IconEdit,
  IconMoodHappy,
  IconPhoneCall
} from '@tabler/icons-vue'
import { format, getYear, parseISO, addDays } from 'date-fns'
import { CakeIcon, CalendarSearchIcon } from 'lucide-vue-next'
import Timeline from '@comp/perfil/vinculos/timeline.vue'

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
  grid-template-rows: min-content min-content 80vh;

  .page-body {
    max-height: 100%;
    height: 100%;
    padding: 0;
    margin: 0;
    .container-xl,
    .row {
      height: 100%;
      .vinculos {
        max-height: 60vh;
        overflow-y: auto;
        .timeline-event-icon {
          height: 3.5vh;
          margin: 0;
          padding: 0;
        }
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
    font-size: 0.95rem;
  }
}
</style>

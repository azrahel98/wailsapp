<template>
  <div class="main">
    <div class="row">
      <div class="col-sm-1 col-md-2 col-g-4">
        <div class="avatar rounded-3 w-100">
          <img
            v-if="perfil.Foto.Valid === false"
            src=""
            class="rounded-circle img-fluid"
            style="object-fit: fill"
          />
          <img
            v-else
            :src="perfil.Foto.String"
            class="rounded-circle img-fluid"
            style="width: 100px; height: 100px; object-fit: cover"
          />
        </div>
      </div>
      <div class="col-sm-1 col-md-1 col-lg-2 text-center">
        <button class="btn" data-bs-toggle="modal" data-bs-target="#editProfileModal">
          <IconEdit size="18" />
        </button>
      </div>
      <div class="col-sm-12 col-md-9 h-100 col-lg-6 text-start">
        <div class="card">
          <div class="card-body pb-1 pt-3 d-flex flex-column">
            <div class="" v-if="perfil.Ruc.Valid == true">
              <IconClipboardList class="icon me-2 text-secondary" />
              Ruc : <strong>{{ perfil.Ruc.String }}</strong>
            </div>
            <div class="" v-if="perfil.Email.Valid == true">
              <IconMail class="icon me-2 text-secondary" />
              Correo : <strong>{{ perfil.Email.String }}</strong>
            </div>
            <div class="" v-if="perfil.Direccion.Valid == true">
              <IconHome class="icon me-2 text-secondary" />
              Vive en: <strong>{{ perfil.Direccion.String }}</strong>
            </div>
            <div class="" v-if="perfil.Telf1.Valid == true">
              <IconPhoneCall class="icon me-2 text-secondary" />
              Telefono: <strong>{{ perfil.Telf1.String }}</strong>
            </div>
            <div class="" v-if="perfil.Nacimiento">
              <IconCalendar class="icon text-secondary me-2" />
              Cumpleaños:
              <strong>{{ formatDate(addDays(perfil.Nacimiento, 1), 'dd MMM yyyy') }}</strong>
            </div>
            <div v-if="perfil.Dni">
              <IconCreditCard class="icon text-secondary me-2" />
              Dni: <strong>{{ perfil.Dni }}</strong>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { addDays, formatDate } from 'date-fns'
import {
  IconHome,
  IconEdit,
  IconCalendar,
  IconClipboardList,
  IconMail,
  IconPhoneCall,
  IconCreditCard
} from '@tabler/icons-vue'
import { onMounted, ref } from 'vue'
import { Buscar_persona_by_dni } from '@wails/services/PersonalService'

const perfil = ref<any>({
  Dni: '',
  Nombre: '',
  Email: {},
  Foto: {},
  Nacimiento: '',
  Direccion: '',
  Ruc: {},
  Sexo: {},
  Telf1: {},
  Telf2: {}
})

onMounted(async () => {
  try {
    const res = await Buscar_persona_by_dni('41662616')
    console.log(res)
    perfil.value = res
  } catch (error) {
    console.log(error)
  }
})
</script>

<style lang="scss" scoped>
.main {
  display: flex;
  justify-content: center;
  flex-direction: column;
  justify-items: center;
  max-width: 1000px;
  .row {
    justify-content: space-around;

    .avatar {
      height: 100%;
      img {
        height: 20vh;
        width: 100%;
      }
    }
    .card {
      overflow-y: auto;
      height: max-content;
      padding-bottom: 1.4vh;
      .card-body {
        justify-content: start;
        align-content: start;
      }
    }
  }
}
</style>

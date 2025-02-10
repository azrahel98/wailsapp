<template>
  <li class="timeline-event text-start">
    <div
      class="timeline-event-icon"
      :class="x.Estado !== 'activo' ? 'bg-secondary ' : 'bg-primary'"
    >
      <IconBriefcase class="icon" :class="x.Estado !== 'activo' ? 'text-white' : 'text-white'" />
    </div>
    <div class="card timeline-event-card">
      <div class="card-body p-0 m-0 mx-2 my-1">
        <div class="float-end fs-5 text-secondary">
          {{
            x.Fecha_ingreso
              ? format(addDays(parseISO(x.Fecha_ingreso), 1), 'yyyy/MM/dd')
              : 'Fecha no disponible'
          }}
        </div>
        <div class="page-subtitle">
          <h5 class="p-0 m-0">{{ x.Cargo }}</h5>
          <p class="text-secondary p-0 m-0 fs-5">
            {{ x.Area }}
          </p>
        </div>

        <div class="info p-0 m-0">
          <div class="accordion" id="accordion-example">
            <div class="accordion-item border-0 p-0 m-0">
              <div
                class="accordion-header w-100 d-flex justify-content-between p-0 m-0"
                id="heading-1"
              >
                <button
                  :id="'collapse#' + x.Id"
                  class="btn btn-sm btn-action"
                  style="height: min-content"
                  data-bs-toggle="collapse"
                  :data-bs-target="'#collapse' + x.Id"
                  aria-expanded
                  @click="click_collapse(x.Id)"
                >
                  <IconEyePlus class="icon text-facebook fw-bold me-1" />
                  <IconEyeX id="abierto" class="icon me-1 text-warning fw-bold d-none" />
                </button>
                <div class="d-flex justify-content-end text-start">
                  <button
                    class="btn-action btn btn-icon btn-sm"
                    style="height: min-content"
                    v-if="store.role == '2' || store.role == '1'"
                  >
                    <IconTrashX class="icon text-youtube" />
                  </button>
                  <button
                    class="btn-action btn-sm btn btn-icon"
                    style="height: min-content"
                    v-if="store.role == '1'"
                  >
                    <IconEdit class="icon text-success" />
                  </button>
                  <button
                    v-if="store.role == '2' || (store.role == '1' && !x.Doc_s)"
                    class="btn-action btn btn-icon btn-sm"
                    data-bs-toggle="modal"
                    :data-bs-target="`#${x.Id}`"
                    style="height: min-content"
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
                              ? format(addDays(parseISO(x.Fecha_salida), 1), 'yyyy/MM/dd')
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
</template>

<script lang="ts" setup>
import { format, addDays, parseISO } from 'date-fns'
import renunciaModal from '@comp/perfil/vinculos/renuncia-modal.vue'
import {
  IconBriefcase,
  IconClipboardOff,
  IconEdit,
  IconEyePlus,
  IconEyeX,
  IconFileInfo,
  IconTrashX,
  IconX
} from '@tabler/icons-vue'
import { userStore } from '../../../store/user'

const store = userStore()

defineProps({
  x: { type: Object, required: true },
  click_collapse: { type: Function, required: true }
})
</script>

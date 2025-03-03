<template>
  <li class="timeline-event text-start">
    <div
      class="timeline-event-icon"
      :class="x.Estado !== 'activo' ? 'bg-secondary ' : 'bg-primary'"
    >
      <IconBriefcase class="icon" :class="x.Estado !== 'activo' ? 'text-white' : 'text-white'" />
    </div>
    <div class="card timeline-event-card">
      <div class="card-body p-0 m-0 mx-2 my-1 py-1">
        <div class="float-end text-secondary fw-medium d-flex flex-column gap-2">
          <span class="badge bg-primary-lt">
            {{
              x.Fecha_ingreso
                ? format(addDays(parseISO(x.Fecha_ingreso), 1), 'yyyy/MM/dd')
                : 'Fecha no disponible'
            }}</span
          >
          <span v-if="x.Fecha_salida" class="badge bg-secondary-lt">
            {{
              x.Fecha_salida
                ? format(addDays(parseISO(x.Fecha_salida), 1), 'yyyy/MM/dd')
                : 'Fecha no disponible'
            }}
          </span>
        </div>

        <div>
          <h4 class="p-0 m-0 fw-semibold">{{ x.Cargo }}</h4>
          <div class="is-sindicato">
            <p class="text-secondary p-0 py-1 m-0" style="width: auto">
              {{ x.Area }}
            </p>
            <span
              v-if="x.Sindicato"
              class="badge bg-warning-lt text-white align-middle"
              style="
                width: auto;
                word-wrap: break-word;
                white-space: normal;
                vertical-align: middle;
              "
              >{{ x.Sindicato }}</span
            >
          </div>
        </div>

        <div class="info p-0 m-0 pt-2">
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
                    class="btn-action btn btn-icon btn-sm"
                    style="height: min-content"
                    data-bs-toggle="modal"
                    :data-bs-target="`#sindicato-${x.Id}`"
                    v-if="store.role == '1' && x.Estado == 'activo' && !x.Sindicato"
                  >
                    <IconCopyPlus class="icon text-primary" />
                  </button>
                  <sindicato :id="x.Id" />
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
                class="accordion-collapse collapse pt-3"
                data-bs-parent="#accordion-example"
                style=""
              >
                <div class="accordion-body pt-0">
                  <div class="detalles-collapse">
                    <div class="ingreso">
                      <div class="mb-2">
                        <IconFileInfo class="icon text-info" />
                        Ingreso:
                        <strong>{{ x.Doc_i }} {{ x.Numero_doc_i }}</strong>
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
                        <div class="badge bg-danger-lt text-danger">
                          <IconClipboardOff class="icon" />
                          Salida:
                        </div>
                        <strong>{{ x.Doc_s }} - {{ x.Numero_doc_s }}</strong>
                      </div>
                      <div class="mb-2" v-if="x.Doc_s">
                        <IconClipboardOff class="icon text-danger" />
                        Descripcion: <strong>{{ x.Descrip_s }}</strong>
                      </div>
                      <div class="mb-2" v-if="x.Doc_s">
                        <div class="badge badge-pill bg-danger-lt">
                          <IconClipboardOff class="icon" />
                          Fecha Renuncia:
                        </div>
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
import sindicato from './sindicato.vue'
import {
  IconBriefcase,
  IconClipboardOff,
  IconCopyPlus,
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

<style lang="scss" scoped>
.detalles-collapse {
  width: 80%;
  justify-content: space-between;
  display: flex;
  flex-wrap: wrap;
  .ingreso,
  .salida {
    .mb-2 {
      text-decoration: dashed;

      strong {
        padding-left: 2vh;
        font-weight: 400;
      }
    }
  }
}
.is-sindicato {
  display: grid;
  grid-template-columns: auto min-content;
  width: 100%;
  .badge {
    justify-self: center;
    align-self: center;
    vertical-align: middle;
    text-align: center;
  }
}
</style>

<template>
  <div class="col-12">
    <div class="card">
      <!-- <div class="card-header">
        <h3 class="card-title">Historial Laboral</h3>
        <div class="card-actions">
          <a href="#" class="btn btn-primary">
            <IconPlus class="icon-inline me-1" />
            Agregar Trabajo
          </a>
        </div>
      </div> -->
      <div class="card-body p-0 m-0">
        <div class="table-responsive">
          <EasyDataTable
            :headers="headers"
            :items="vinculos"
            :buttons-pagination="false"
            :hide-footer="true"
          >
            <template #loading>
              <img
                src="https://i.pinimg.com/originals/94/fd/2b/94fd2bf50097ade743220761f41693d5.gif"
                style="width: 100px; height: 80px"
              />
            </template>
            <!-- <template #loading>
              <img
                src="https://thumbs.gfycat.com/AngelicYellowIberianmole.webp"
                style="width: 60px; height: 100px"
              />
            </template>
            <template #item-player="{ player, avator, page }">
              <div class="player-wrapper">
                <img class="avator" :src="avator" alt="" />
                <a target="_blank" :href="page">{{ player }}</a>
              </div>
            </template>
            <template #item-team="{ teamName, teamUrl }">
              <a target="_blank" :href="teamUrl">{{ teamName }}</a>
            </template>
            <template #expand="item">
              <div style="padding: 15px">{{ item.player }} won championships in</div>
            </template> -->
          </EasyDataTable>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Search_by_dni_vinculos } from '@wails/services/PersonalService'
import { router } from '@router/router'
import { IconPlus } from '@tabler/icons-vue'
const vinculos = ref<Array<any>>([])

onMounted(async () => {
  try {
    vinculos.value = await Search_by_dni_vinculos(router.currentRoute.value.params.dni.toString())
  } catch (error) {
    console.log(error)
  }
})
import type { Header, Item } from 'vue3-easy-data-table'

const headers: Header[] = [
  { text: 'Ingreso', value: 'Doc_i' },
  { text: 'Numero', value: 'Numero_doc_i' },
  { text: 'Descr', value: 'Descrip_ingre' },
  { text: 'Fecha', value: 'Fecha_ingreso' },
  { text: 'Area', value: 'Area' },
  { text: 'Cargo', value: 'Cargo', sortable: true },
  { text: 'Regimen', value: 'Regimen' },
  { text: 'Sueldo', value: 'Sueldo' },
  { text: 'Estado', value: 'Estado' }
]
</script>

<style lang="scss" scoped>
.operation-wrapper .operation-icon {
  width: 20px;
  cursor: pointer;
}
.player-wrapper {
  padding: 5px;
  display: flex;
  align-items: center;
  justify-items: center;
}
.avator {
  margin-right: 10px;
  display: inline-block;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
  box-shadow: inset 0 2px 4px 0 rgb(0 0 0 / 10%);
}
</style>

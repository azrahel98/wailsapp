<template>
  <div class="flex flex-column">
    <div class="bg-azure">
      <input type="file" v-on:change="check" />
      <button class="btn btn-action" @click="files">Cargar</button>
    </div>
    <div>
      <button @click="buscar">Generar</button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { ReadXmls_folder } from '@wails/services/BoletaService'

const file = ref()

const check = (e: any) => {
  console.log(e.target.files)
}

const buscar = async () => {
  console.log(await ReadXmls_folder())
}

const files = async () => {
  try {
    const res = () => {}
    console.log(res)
  } catch (error) {
    console.log(error)
  }
}
</script>

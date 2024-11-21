<template>
  <div class="flex flex-column">
    <div class="bg-azure">
      <input type="file" v-on:change="check" />
      <button class="btn btn-action" @click="files">Cargar</button>
    </div>
    <div></div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { Upload_file } from '@wails/services/BoletaService'

const file = ref()

const check = (e: any) => {
  console.log(e.target.files)
}

const files = async () => {
  try {
    const res = await Upload_file()
    console.log(res)
  } catch (error) {
    console.log(error)
  }
}
</script>

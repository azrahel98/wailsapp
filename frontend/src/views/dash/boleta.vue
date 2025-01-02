<template>
  <div class="main">
    <div class="col text-start mt-4">
      <div class="page-pretitle fw-medium">Generar Boletas de Pago</div>
      <h2 class="page-title">Boleta</h2>
    </div>

    <div class="archivos-xml card">
      <div class="folder card" v-for="x in archivos">
        <IconFileTypeXml class="icon text-facebook icon-md text-center align-middle" />
        <span>{{ x.Full }}</span>
      </div>
    </div>
    <div class="botones">
      <button @click="archivos_xml" class="btn btn-md btn-facebook">archivos_xml</button>
      <button @click="Generar_pdf" class="btn btn-md btn-facebook">pdf</button>
    </div>
    <div class="pdf card">
      <div class="folder card" v-for="x in pdfresultado">
        <IconFileTypePdf class="icon text-youtube icon-md text-center align-middle" />
        <span>{{ x.Nombre }}</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { ReadXmls_folder, Listar_Xmls } from '@wails/services/BoletaService'
import { IconFileTypePdf, IconFileTypeXml } from '@tabler/icons-vue'

const file = ref()

const archivos = ref<Array<any>>([])
const pdfresultado = ref<Array<any>>([])

const archivos_xml = async () => {
  archivos.value = await Listar_Xmls()
}

const Generar_pdf = async () => {
  try {
    pdfresultado.value = await ReadXmls_folder()
    console.log(pdfresultado.value)
  } catch (error) {
    console.log(error)
  }
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

<style lang="scss" scoped>
.main {
  height: 100vh;
  display: grid;
  grid-template-rows: min-content 1fr min-content 1fr;
  row-gap: 2vh;
  .archivos-xml,
  .pdf {
    display: flex;
    flex-direction: row;
    justify-content: center;
    width: 100%;
    padding: 2vh;
    flex-wrap: wrap;
    max-height: 100%;
    overflow-y: scroll;
    gap: 2vh;

    .folder {
      display: flex;
      flex-direction: column;
      align-items: center;
      width: 15vw;
      height: 15vh;
      span {
        width: 100%;
        font-size: 0.6rem;
        white-space: wrap;
      }
    }
  }
  .botones {
    display: flex;
    justify-content: center;
    gap: 2vh;
  }
  .pdf {
  }
}
</style>

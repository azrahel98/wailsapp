<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Regimenes Laborales</h3>
    </div>
    <table class="table card-table table-vcenter">
      <thead>
        <tr>
          <th>Regimenes</th>
          <th colspan="2">Servidores</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="x in (reg as Array<any>)">
          <td>{{ x.nombre }}</td>
          <td>{{ x.cantidad }}</td>
          <td class="w-50">
            <div class="progress progress-xs">
              <div class="progress-bar bg-primary" :style="`width:${x.porcentaje}%`"></div>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts" setup>
defineProps({
  reg: { type: Array, required: true }
})
</script>

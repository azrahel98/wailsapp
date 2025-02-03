<template>
  <div class="chart-container card card-sm">
    <div class="card-body p-0 m-0">
      <div class="text-secondary subheader">Personal por regimen</div>
      <BarChart v-if="chartData" :chartData="chartData" :options="chartOptions" :height="200" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { BarChart } from 'vue-chart-3'
import { Chart, registerables } from 'chart.js'
import { computed } from 'vue'

Chart.register(...registerables)

const chartData = computed(() => {
  if (!prop.regimes?.length) return null
  const sortedData = [...prop.regimes].sort((a, b) => b.Cantidad - a.Cantidad)

  return {
    labels: sortedData.map((r) => r.Nombre),
    datasets: [
      {
        label: 'Masculino',
        data: sortedData.map((r) => (r.Nombre == 'F' ? 1 : r.Cantidad)),
        backgroundColor: '#007BFF',
        borderColor: '#FFFFFF',
        borderWidth: 1,
        borderRadius: 4
      },
      {
        label: 'Femenino',
        data: sortedData.map((r) => (r.Nombre == 'M' ? 1 : r.Cantidad)),
        backgroundColor: '#FF69B4',
        borderColor: '#FFFFFF',
        borderWidth: 1,
        borderRadius: 4
      }
    ]
  }
})

const chartOptions = {
  responsive: true,
  plugins: { legend: { display: false } },
  scales: {
    y: { beginAtZero: true, ticks: { font: { size: 10 } } },
    x: { ticks: { font: { size: 10 } } }
  }
}

interface Regimen {
  Cantidad: number
  Nombre: string
}

const prop = defineProps({
  regimes: { type: Array<Regimen>, required: true },
  title: { type: String, default: 'Personal por regimen' }
})
</script>

<style scoped>
.chart-container {
  position: relative;
  width: 100%;
  max-width: 900px;
  height: 100%;
  margin: 0 auto;
  padding: 1rem;

  border-radius: 0.5rem;
}
</style>

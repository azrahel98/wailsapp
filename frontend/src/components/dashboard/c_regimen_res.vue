<template>
  <div class="chart-container card card-sm">
    <div class="card-body p-0 m-0">
      <DoughnutChart
        v-if="chartData"
        :chartData="chartData"
        :options="chartOptions"
        :height="200"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { DoughnutChart } from 'vue-chart-3'
import { Chart, registerables } from 'chart.js'
import { computed, ref } from 'vue'

Chart.register(...registerables)

interface Regimen {
  Cantidad: number
  Nombre: string
}

const props = defineProps({
  regim: { type: Array<Regimen>, required: true },
  responsive: { type: Boolean, default: true }
})

const chartData = computed(() => {
  if (!props.regim.length) return null

  const sortedData = [...props.regim].sort((a, b) => b.Cantidad - a.Cantidad)

  return {
    labels: sortedData.map((r) => r.Nombre),
    datasets: [
      {
        label: 'Cantidad de Personal',
        data: sortedData.map((r) => r.Cantidad),
        backgroundColor: ['#222940', '#A7C8F2', '#2B88D9', '#04D976', '#F2F2F2'],
        borderColor: '#FFFFFF',
        borderWidth: 1,
        borderRadius: 4
      }
    ]
  }
})

const chartOptions = computed(() => ({
  indexAxis: 'y' as const,
  responsive: props.responsive,
  maintainAspectRatio: true,
  plugins: {
    legend: {
      display: true,
      position: 'right',
      labels: {
        boxWidth: 10,
        padding: 5,
        font: {
          size: 10
        }
      }
    },
    tooltip: {
      callbacks: {
        label: (context: any) => {
          const value = context.raw
          const total = context.dataset.data.reduce((a: number, b: number) => a + b, 0)
          const percentage = ((value / total) * 100).toFixed(1)
          return `${value} personas (${percentage}%)`
        }
      }
    }
  }
}))
</script>

<style scoped>
.chart-container {
  position: relative;
  width: 100%;
  max-width: 900px;
  margin: 0 auto;
  padding: 1rem;
  border-radius: 0.5rem;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}
</style>

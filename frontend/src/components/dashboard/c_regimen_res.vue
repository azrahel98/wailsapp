<template>
  <div class="chart-container card card-sm">
    <div class="card-body p-0 m-0">
      <div class="text-secondary subheader">Personal por regimen</div>
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
import { computed } from 'vue'

Chart.register(...registerables)

interface Regimen {
  Cantidad: number
  Nombre: string
}

const prop = defineProps({
  regim: { type: Array<Regimen>, required: true }
})

const chartData = computed(() => {
  if (!prop.regim.length) return null
  const sortedData = [...prop.regim].sort((a, b) => b.Cantidad - a.Cantidad)

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

const chartOptions = {
  indexAxis: 'y' as const,
  responsive: true,
  maintainAspectRatio: true,
  plugins: {
    animations: {
      tension: {
        duration: 1000,
        easing: 'linear',
        from: 1,
        to: 0,
        loop: true
      }
    },
    title: {
      display: true,
      padding: {
        top: 1
      }
    },
    legend: {
      display: true,
      position: 'bottom',
      labels: {
        pointStyle: 'circle',
        usePointStyle: true
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
}
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

<template>
  <div class="chart-container card">
    <div class="card-body p-0 m-0">
      <div class="text-secondary subheader">Personal por regimen</div>
      <BarChart v-if="chartData" :chartData="chartData" :options="chartOptions" :height="200" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { BarChart } from 'vue-chart-3'
import { Chart, registerables } from 'chart.js'
import { onMounted, ref, computed } from 'vue'
import { Sexo_cantidad } from '@wails/services/DashboardService'

Chart.register(...registerables)

interface Regimen {
  Cantidad: number
  Nombre: string
}

const regimenes = ref<Array<Regimen>>([])

const chartData = computed(() => {
  if (!regimenes.value?.length) return null
  const sortedData = [...regimenes.value].sort((a, b) => b.Cantidad - a.Cantidad)

  return {
    labels: sortedData.map((r) => r.Nombre),
    datasets: [
      {
        label: 'Masculino',
        data: sortedData.map((r) => (r.Nombre == 'F' ? 1 : r.Cantidad)), // 1 para masculino, 0 para no masculino
        backgroundColor: '#007BFF', // Azul para masculino
        borderColor: '#FFFFFF',
        borderWidth: 1,
        borderRadius: 4
      },
      {
        label: 'Femenino',
        data: sortedData.map((r) => (r.Nombre == 'M' ? 1 : r.Cantidad)), // 1 para femenino, 0 para no femenino
        backgroundColor: '#FF69B4', // Rosa para femenino
        borderColor: '#FFFFFF',
        borderWidth: 1,
        borderRadius: 4
      }
    ]
  }
})

const chartOptions = {
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

onMounted(async () => {
  try {
    regimenes.value = await Sexo_cantidad()
  } catch (error) {
    console.error('Error al cargar los datos:', error)
  }
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

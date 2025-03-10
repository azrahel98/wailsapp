<template>
  <div class="organigrama">
    <!-- Nivel 1: Gerente General -->
    <div class="nivel nivel-1">
      <div class="card node">
        <div class="card-body text-center">
          <h5 class="card-title">{{ generalManager.Area }}</h5>
        </div>
      </div>
    </div>

    <!-- Nivel 2: Gerencias -->
    <div class="nivel nivel-2">
      <div class="linea-vertical"></div>
      <div class="grid-container">
        <div v-for="gerencia in gerencias" :key="gerencia.Id" class="card-container">
          <div class="linea-vertical-item"></div>
          <div class="card node">
            <div class="card-body text-center">
              <h6 class="card-title">{{ gerencia.Area }}</h6>
              <p class="card-text" v-if="gerencia.Jefe">Jefe: {{ gerencia.Jefe }}</p>
            </div>
          </div>

          <!-- Nivel 3: Subgerencias -->
          <div class="nivel nivel-3" v-if="gerencia.Subgerencias && gerencia.Subgerencias.length">
            <div class="grid-container">
              <div
                v-for="subgerencia in gerencia.Subgerencias"
                :key="subgerencia.Id"
                class="card-container"
              >
                <div class="linea-vertical-item"></div>
                <div class="card node">
                  <div class="card-body text-center">
                    <h6 class="card-title">{{ subgerencia.Area }}</h6>
                    <p class="card-text" v-if="subgerencia.Jefe">Jefe: {{ subgerencia.Jefe }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { Reporte_Organigrama } from '@wails/services/DashboardService'

const generalManager = ref({ Id: 5, Area: 'GERENCIA MUNICIPAL' })
const gerencias = ref([])

onMounted(async () => {
  try {
    const data = await Reporte_Organigrama()
    gerencias.value = data[0]?.Subgerencias || []
  } catch (error) {
    console.log(error)
  }
})
</script>

<style scoped>
.organigrama {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 3rem;
  padding: 2rem 1rem;
  overflow-x: hidden;
}

.nivel {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
}

.grid-container {
  width: 100%;
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  justify-content: center;
  position: relative;
}

.card-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
  min-width: 150px;
}

.linea-vertical-item {
  width: 2px;
  height: 1.5rem;
  background-color: #dee2e6;
}

.node {
  width: 100%;
  max-width: 250px;
  border: none;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  background-color: white;
  padding: 10px;
  text-align: center;
}
</style>

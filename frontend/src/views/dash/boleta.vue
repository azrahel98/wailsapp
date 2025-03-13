<template>
  <div class="organigrama">
    <!-- Nivel 1: Gerente General -->
    <div class="nivel nivel-1">
      <div class="card">
        <div class="card-body text-center p-4">
          <div class="card-title">{{ generalManager.Area }}</div>
        </div>
      </div>
    </div>

    <!-- Nivel 2: Gerencias -->
    <div class="nivel nivel-2">
      <div class="linea-vertical"></div>
      <div class="grid-container">
        <div v-for="gerencia in gerencias" :key="gerencia.Id" class="card-container">
          <div class="linea-vertical-item"></div>
          <div class="card">
            <div class="card-body text-center p-3">
              <div class="card-title">{{ gerencia.Area }}</div>
              <div class="text-muted" v-if="gerencia.Jefe">Jefe: {{ gerencia.Jefe }}</div>
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
                <div class="card">
                  <div class="card-body text-center p-2">
                    <div class="card-title fs-6">{{ subgerencia.Area }}</div>
                    <div class="text-muted small" v-if="subgerencia.Jefe">
                      Jefe: {{ subgerencia.Jefe }}
                    </div>
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
  overflow-x: auto;
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
  gap: 1.5rem;
  justify-content: center;
  position: relative;
}

.card-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
  min-width: 180px;
}

.linea-vertical {
  width: 2px;
  height: 2rem;
  background-color: var(--tblr-border-color);
}

.linea-vertical-item {
  width: 2px;
  height: 1.5rem;
  background-color: var(--tblr-border-color);
}

.card {
  width: 100%;
  max-width: 250px;
  border-radius: 4px;
  box-shadow: rgba(30, 41, 59, 0.04) 0 2px 4px 0;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.card:hover {
  transform: translateY(-2px);
  box-shadow: rgba(30, 41, 59, 0.1) 0 4px 8px 0;
}

.nivel-1 .card {
  background-color: var(--tblr-primary);
  color: #fff;
}

.nivel-1 .card-title {
  font-weight: 600;
  font-size: 1.1rem;
}

.nivel-2 .card {
  background-color: var(--tblr-light);
  border-left: 3px solid var(--tblr-primary);
}

.nivel-3 .card {
  background-color: #fff;
  border: 1px solid var(--tblr-border-color);
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .organigrama {
    padding: 1rem 0.5rem;
  }

  .grid-container {
    gap: 1rem;
  }

  .card-container {
    min-width: 150px;
  }
}
</style>

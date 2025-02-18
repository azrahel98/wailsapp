<template>
  <div class="organigrama">
    <!-- Nivel 1: Gerente General -->
    <div class="nivel nivel-1">
      <div class="card node">
        <div class="card-body text-center">
          <img :src="generalManager.image" class="rounded-circle mb-2" alt="Gerente General" />
          <h5 class="card-title">{{ generalManager.name }}</h5>
          <p class="card-text">{{ generalManager.position }}</p>
        </div>
      </div>
    </div>

    <!-- Nivel 2: Gerencias -->
    <div class="nivel nivel-2">
      <div class="linea-vertical"></div>
      <div class="grid-container">
        <div v-for="gerente in gerentes" :key="gerente.id" class="card-container">
          <div class="linea-vertical-item"></div>
          <div class="card node">
            <div class="card-body text-center">
              <img :src="gerente.image" class="rounded-circle mb-2" :alt="gerente.name" />
              <h6 class="card-title">{{ gerente.name }}</h6>
              <p class="card-text">{{ gerente.position }}</p>
            </div>
          </div>

          <!-- Nivel 3: Subgerencias -->
          <div class="nivel nivel-3" v-if="gerente.subgerencias.length">
            <div class="grid-container">
              <div
                v-for="subgerente in gerente.subgerencias"
                :key="subgerente.id"
                class="card-container"
              >
                <div class="linea-vertical-item"></div>
                <div class="card node">
                  <div class="card-body text-center">
                    <img
                      :src="subgerente.image"
                      class="rounded-circle mb-2"
                      :alt="subgerente.name"
                    />
                    <h6 class="card-title">{{ subgerente.name }}</h6>
                    <p class="card-text">{{ subgerente.position }}</p>
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
import { ref } from 'vue'

const generalManager = ref({
  id: 1,
  name: 'Juan Pérez',
  position: 'Gerente General',
  image: '/path/to/general-manager.jpg'
})

const gerentes = ref([
  {
    id: 2,
    name: 'Carlos Rodríguez',
    position: 'Gerente de Rentas',
    image: '/path/to/gerente1.jpg',
    subgerencias: [
      {
        id: 5,
        name: 'Luis Gómez',
        position: 'Subgerente de Recaudación',
        image: '/path/to/subgerente1.jpg'
      },
      {
        id: 6,
        name: 'Ana Torres',
        position: 'Subgerente de Fiscalización',
        image: '/path/to/subgerente2.jpg'
      }
    ]
  },
  {
    id: 3,
    name: 'María Sánchez',
    position: 'Gerente de Administración',
    image: '/path/to/gerente2.jpg',
    subgerencias: [
      {
        id: 7,
        name: 'Pedro Ramírez',
        position: 'Subgerente de Recursos Humanos',
        image: '/path/to/subgerente3.jpg'
      },
      {
        id: 8,
        name: 'Carmen López',
        position: 'Subgerente de Tesorería',
        image: '/path/to/subgerente4.jpg'
      }
    ]
  },
  {
    id: 4,
    name: 'Fernando Morales',
    position: 'Gerente de Seguridad Ciudadana',
    image: '/path/to/gerente3.jpg',
    subgerencias: [
      {
        id: 9,
        name: 'Andrea Castro',
        position: 'Subgerente de Serenazgo',
        image: '/path/to/subgerente5.jpg'
      },
      {
        id: 10,
        name: 'Miguel Ruiz',
        position: 'Subgerente de Transporte',
        image: '/path/to/subgerente6.jpg'
      },
      {
        id: 11,
        name: 'Andrea Castro',
        position: 'Subgerente de Serenazgo',
        image: '/path/to/subgerente5.jpg'
      },
      {
        id: 12,
        name: 'Miguel Ruiz',
        position: 'Subgerente de Transporte',
        image: '/path/to/subgerente6.jpg'
      }
    ]
  }
])
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
  max-width: 180px;
  border: none;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  background-color: white;
}

.linea-horizontal {
  width: 100%;
  height: 2px;
  background-color: #dee2e6;
  position: absolute;
  top: 50%;
  left: 0;
}
</style>

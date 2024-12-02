<template>
  <div class="perspective-1000 w-64 h-96 relative cursor-pointer" @click="toggleCard">
    <div
      class="card-inner relative w-full h-full transition-transform duration-700 ease-in-out transform-style-3d"
      :class="{ 'rotate-y-180': isFlipped }"
    >
      <!-- Lado Frontal -->
      <div
        class="card-front absolute w-full h-full backface-hidden bg-white shadow-lg rounded-xl overflow-hidden"
      >
        <img
          src="/api/placeholder/300/400"
          alt="Frente de la tarjeta"
          class="w-full h-48 object-cover"
        />
        <div class="p-4">
          <h2 class="text-xl font-bold mb-2">Título de la Tarjeta</h2>
          <p class="text-gray-600">Descripción breve del contenido de la tarjeta.</p>
        </div>
      </div>

      <!-- Lado Posterior -->
      <div
        class="card-back absolute w-full h-full backface-hidden rotate-y-180 bg-blue-500 text-white shadow-lg rounded-xl overflow-hidden"
      >
        <div class="p-4 flex flex-col h-full">
          <h3 class="text-xl font-bold mb-4">Detalles Adicionales</h3>
          <p class="flex-grow">
            Aquí puedes mostrar información adicional, detalles más específicos o cualquier
            contenido que quieras revelar al voltear la tarjeta.
          </p>
          <div class="mt-4 text-sm opacity-75">Toca para volver</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const isFlipped = ref(false)

const toggleCard = () => {
  isFlipped.value = !isFlipped.value
}
</script>

<style scoped>
.perspective-1000 {
  perspective: 1000px;
}

.transform-style-3d {
  transform-style: preserve-3d;
}

.backface-hidden {
  backface-visibility: hidden;
}

.rotate-y-180 {
  transform: rotateY(180deg);
}
</style>

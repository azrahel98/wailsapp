<template>
  <div class="card card-sm">
    <div class="card-body">
      <div class="row align-items-center">
        <div class="col-auto">
          <slot />
        </div>
        <div class="col">
          <div class="font-weight-medium">{{ title }}</div>
          <div class="text-secondary">{{ cantidad }} {{ descripcion }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps({
  title: { type: String, required: true },
  cantidad: { type: Number, required: false, default: 50 },
  descripcion: { type: String, required: false, default: '' }
})
</script>

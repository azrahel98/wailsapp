<template>
  <Transition name="fade">
    <div class="toast-container position-fixed bottom-0 end-0 p-3">
      <div
        class="toast fade"
        :class="isVisible ? 'show' : 'hide'"
        id="toast-simple"
        role="alert"
        aria-live="assertive"
        aria-atomic="true"
        data-bs-autohide="true"
      >
        <div class="toast-header bg-primary-lt text-warning">
          <span class="avatar avatar-xs me-2"></span>
          <strong class="me-auto">{{ store.username }}</strong>

          <button
            type="button"
            class="ms-2 btn-close"
            data-bs-dismiss="toast"
            aria-label="Close"
          ></button>
        </div>
        <div class="toast-body">{{ message }}</div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { ref, defineExpose } from 'vue'
import { userStore } from '../store/user'

const isVisible = ref(false)
const message = ref('Proceso completado con éxito')
const store = userStore()

const showToast = (text: string = 'Proceso completado') => {
  message.value = text
  console.log('aaui')
  isVisible.value = true
  setTimeout(() => {
    isVisible.value = false
  }, 5000)
}

defineExpose({ showToast })
</script>

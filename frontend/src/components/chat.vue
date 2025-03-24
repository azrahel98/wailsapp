<template>
  <div class="chat-container">
    <div v-if="isOpen" class="chat-overlay" @click="toggleChat"></div>

    <button class="btn btn-primary chat-toggle" @click="toggleChat">
      <IconMessage />
    </button>

    <div v-if="isOpen" class="chat-window card">
      <div class="card-header d-flex justify-content-between align-items-center">
        <strong>Chat</strong>
        <button class="btn btn-icon btn-sm p-2 btn-danger" @click="toggleChat">
          <IconX class="icon" />
        </button>
      </div>
      <div ref="chatBody" class="card-body chat-body">
        <div
          v-for="(msg, index) in messages.filter((msg) => msg.role !== 'developer')"
          :key="index"
          class="chat-message chat-bubble chat-bubble"
          :class="[msg.role === 'user' ? 'bg-primary-lt' : 'bg-secondary-lt']"
        >
          <span
            :class="[
              'badge text-white fw-normal',
              msg.role === 'user' ? 'bg-primary' : 'bg-secondary'
            ]"
          >
            {{ msg.role == 'user' ? store.username : msg.role == 'assistant' ? 'Bot' : 'xd' }}
          </span>
          <p class="mb-1 text-black" v-html="msg.content" />
        </div>
      </div>
      <div class="card-footer">
        <input
          v-model="newMessage"
          @keyup.enter="sendMessage"
          :disabled="loading"
          class="form-control"
          placeholder="Escribe un mensaje..."
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { IconMessage, IconX } from '@tabler/icons-vue'
import { ref, nextTick } from 'vue'
import { Consulta_IA } from '@wails/services/IaService'
import { userStore } from '../store/user'

const store = userStore()

const isOpen = ref(false)
const messages = ref([])
const newMessage = ref('')
const loading = ref(false)
const chatBody = ref(null)

const toggleChat = () => {
  isOpen.value = !isOpen.value
}

const sendMessage = async () => {
  if (newMessage.value.trim() === '') return
  loading.value = true
  messages.value.push(
    { role: 'user', content: newMessage.value },
    {
      role: 'developer',
      content: `el usuario ${store.isAdmin == 1 ? 'es admin' : 'no es admin'}`
    }
  )
  await nextTick()
  scrollToBottom()

  await gtpquestion()
  newMessage.value = ''
  console.log(messages.value)
  loading.value = false
}

const gtpquestion = async () => {
  try {
    var res = await Consulta_IA(newMessage.value, messages.value, store.isAdmin == 1 ? true : false)
    messages.value.push({ role: 'assistant', content: res })
  } catch (error) {
    messages.value.push({ role: 'assistant', content: 'Error: ' + error })
  }

  await nextTick()
  scrollToBottom()
}

const scrollToBottom = () => {
  if (chatBody.value) {
    chatBody.value.scrollTop = chatBody.value.scrollHeight
  }
}
</script>

<style scoped>
.chat-container {
  position: fixed;
  bottom: 20px;
  right: 20px;
  z-index: 1050;
}

.chat-overlay {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 1040;
  animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.chat-toggle {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1051;
}

.chat-window {
  width: 350px;
  max-height: 500px;
  display: flex;
  flex-direction: column;
  z-index: 1051;
  position: relative;
  animation: slideIn 0.3s ease;
}

@keyframes slideIn {
  from {
    transform: translateY(20px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

.chat-body {
  overflow-y: auto;
  max-height: 4000px;
  padding-right: 5px;
  scroll-behavior: smooth;
}

.chat-message {
  margin-bottom: 10px;
}
</style>

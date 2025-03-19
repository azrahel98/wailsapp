<template>
  <div class="chat-container">
    <button class="btn btn-primary chat-toggle" @click="toggleChat">
      <IconMessage />
    </button>

    <div v-if="isOpen" class="chat-window card">
      <div class="card-header d-flex justify-content-between align-items-center">
        <strong>Chat</strong>
        <button class="btn btn-sm btn-danger" @click="toggleChat">
          <IconX class="icon" />
        </button>
      </div>
      <div ref="chatBody" class="card-body chat-body">
        <div v-for="(msg, index) in messages" :key="index" class="chat-message">
          <span
            :class="['badge', msg.user === 'Me' ? 'bg-primary' : 'bg-secondary']"
            class="text-white fw-normal"
          >
            {{ msg.user }}
          </span>
          <p class="mb-1" v-html="msg.text"></p>
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

const isOpen = ref(false)
const messages = ref([{ user: 'Bot', text: '¡Hola! ¿Cómo puedo ayudarte?' }])
const newMessage = ref('')
const loading = ref(false)
const chatBody = ref(null)

const toggleChat = () => {
  isOpen.value = !isOpen.value
}

const sendMessage = async () => {
  if (newMessage.value.trim() === '') return
  loading.value = true
  messages.value.push({ user: 'Me', text: newMessage.value })

  await nextTick()
  scrollToBottom()

  await gtpquestion()
  newMessage.value = ''
  loading.value = false
}

const gtpquestion = async () => {
  try {
    var res = await Consulta_IA(newMessage.value)

    console.log(res)
    messages.value.push({ user: 'Bot', text: res })
  } catch (error) {
    messages.value.push({ user: 'Bot', text: 'Error: ' + error })
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

.chat-toggle {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.chat-window {
  width: 35de a0px;
  max-height: 400px;
  display: flex;
  flex-direction: column;
}

.chat-body {
  overflow-y: auto;
  max-height: 250px;
  padding-right: 5px;
  scroll-behavior: smooth;
}

.chat-message {
  margin-bottom: 10px;
}
</style>

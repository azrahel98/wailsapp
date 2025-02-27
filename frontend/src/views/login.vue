<template>
  <div class="login-container">
    <div class="login-card">
      <h1 class="login-title">Iniciar sesión</h1>
      <form @submit.prevent="handleSubmit" class="login-form" autocomplete="off">
        <div class="form-group">
          <label for="email">Correo electrónico</label>
          <input v-model="email" type="text" id="nick" placeholder="rscl" required />
        </div>
        <div class="form-group">
          <label for="password">Contraseña</label>
          <div class="password-input">
            <input
              v-model="password"
              :type="showPassword ? 'text' : 'password'"
              id="password"
              placeholder="Tu contraseña"
              required
              autocomplete="off"
            />
            <button type="button" @click="togglePassword" class="password-toggle">
              <span v-if="showPassword">🙈</span>
              <span v-else>👁️</span>
            </button>
          </div>
        </div>
        <div class="form-group remember-me">
          <a href="#" class="forgot-password">¿Olvidaste tu contraseña?</a>
        </div>
        <button type="submit" class="login-button">Iniciar sesión</button>
      </form>
      <p class="signup-link">¿No tienes una cuenta? <a href="#">Regístrate aquí</a></p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Login } from '@wails/services/LoginService'
import { router } from '@router/router'

const email = ref('')
const password = ref('')
const showPassword = ref(false)

const togglePassword = () => {
  showPassword.value = !showPassword.value
}

const handleSubmit = async () => {
  try {
    const r = await Login(email.value, password.value)

    localStorage.clear()
    localStorage.setItem('jwt', r)

    router.replace({ name: 'dashboard' })
  } catch (error) {
    console.log(error)
  }
}
</script>

<style scoped lang="scss">
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 20px;

  .login-card {
    background: white;
    padding: 40px;
    border-radius: 20px;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 400px;

    .login-title {
      color: #333;
      font-size: 28px;
      margin-bottom: 30px;
      text-align: center;
    }

    .login-form {
      display: flex;
      flex-direction: column;
      gap: 20px;

      .form-group {
        display: flex;
        flex-direction: column;

        label {
          margin-bottom: 5px;
          font-weight: 500;
          color: #555;
        }

        input {
          padding: 12px;
          border: 1px solid #ddd;
          border-radius: 8px;
          font-size: 16px;
          transition: border-color 0.3s ease;

          &:focus {
            outline: none;
            border-color: #667eea;
          }
        }
      }

      .password-input {
        position: relative;

        input {
          width: 100%;
        }

        .password-toggle {
          position: absolute;
          right: 10px;
          top: 50%;
          transform: translateY(-50%);
          background: none;
          border: none;
          cursor: pointer;
          font-size: 20px;
        }
      }

      .remember-me {
        display: flex;
        justify-content: space-between;
        align-items: center;
        font-size: 14px;

        label {
          display: flex;
          align-items: center;
          gap: 5px;
        }

        .forgot-password {
          color: #667eea;
          text-decoration: none;
        }
      }

      .login-button {
        background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
        color: white;
        padding: 12px;
        border: none;
        border-radius: 8px;
        font-size: 16px;
        cursor: pointer;
        transition: opacity 0.3s ease;

        &:hover {
          opacity: 0.9;
        }
      }
    }

    .signup-link {
      text-align: center;
      margin-top: 20px;
      font-size: 14px;
      color: black;

      a {
        color: #667eea;
        text-decoration: none;
        font-weight: 500;
      }
    }
  }
}
</style>

import { router } from '@router/router'
import { jwtDecode } from 'jwt-decode'
import { defineStore } from 'pinia'

export const userStore = defineStore('userStore', {
  state: () => ({ exp: 0, iat: 0, iss: '', nbf: 0, role: '', sub: '', use_id: '', username: '' }),
  getters: {
    isAdmin: (state) => state.role
  },
  actions: {
    create(token: string) {
      try {
        const decoded: any = jwtDecode(token)
        const currentTime = Math.floor(Date.now() / 1000)

        if (decoded.exp && decoded.exp < currentTime) {
          this.logout()
          return
        }

        this.exp = decoded.exp!
        this.iat = decoded.iat!
        this.nbf = decoded.nbf!
        this.iss = decoded.iss!
        this.sub = decoded.sub!
        this.role = decoded.role!
        this.use_id = decoded.use_id!
        this.username = decoded.username!
      } catch (error) {
        this.logout()
      }
    },
    async logout() {
      localStorage.clear()
      router.go(0)
    }
  }
})
